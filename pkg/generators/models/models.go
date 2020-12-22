package models

import (
	"context"
	"fmt"
	"io"
	"math"
	"sort"
	"text/template"

	tpl "github.com/contiamo/openapi-generator-go/pkg/generators/templates"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

type ModelKind string

const (
	Struct   ModelKind = "struct"
	Enum     ModelKind = "enum"
	Constant ModelKind = "constant"
)

type Model struct {
	Name        string
	Description string
	Imports     []string
	Kind        ModelKind
	Properties  []PropSpec
	GoType      string
	SpecTitle   string
	SpecVersion string
	PackageName string
}

type PropSpec struct {
	Name        string // Property name of structs, variable name of enumns and constants
	Description string
	GoType      string
	JSONTags    string
	Value       string
	IsRequired  bool
	IsEnum      bool

	// Validation stuff
	IsNullable                 bool
	NeedsValidation            bool
	IsRequiredInValidation     bool
	HasMin, HasMax             bool
	Min, Max                   float64
	HasMinLength, HasMaxLength bool
	MinLength, MaxLength       uint64
	IsDate                     bool
	IsDateTime                 bool
	IsBase64                   bool
	IsEmail                    bool
	IsUUID                     bool
	IsURL                      bool
	IsHostname                 bool
	IsIP                       bool
	IsIPv4                     bool
	IsIPv6                     bool
}

func NewModelFromRef(ref *openapi3.SchemaRef) (model *Model, err error) {
	model = &Model{
		Description: ref.Value.Description,
	}

	ref = resolveAllOf(ref)

	switch {
	case len(ref.Value.Enum) > 1:
		model.Kind = Enum
		model.Properties, err = enumPropsFromRef(ref, model)
		model.GoType = goTypeFromSpec(ref)
	case len(ref.Value.Enum) == 1:
		model.Kind = Constant
		model.Properties, err = enumPropsFromRef(ref, model)
		model.GoType = goTypeFromSpec(ref)
	case ref.Value.Type == "object" ||
		len(ref.Value.Properties) > 0 ||
		len(ref.Value.AllOf) > 0 ||
		len(ref.Value.OneOf) > 0:
		model.Kind = Struct
		model.Properties, model.Imports, err = structPropsFromRef(ref)
	default:
		return nil, errors.New("type not handled")
	}
	if err != nil {
		return nil, err
	}

	return model, nil
}

func (m *Model) Render(ctx context.Context, writer io.Writer) error {
	var tpl *template.Template
	switch m.Kind {
	case Struct:
		tpl = modelTemplate
	case Enum:
		tpl = enumTemplate
	case Constant:
		tpl = constTemplate
	}
	err := tpl.Execute(writer, m)
	if err != nil {
		return errors.Wrap(err, "failed to render model")
	}
	return nil
}

func resolveAllOf(ref *openapi3.SchemaRef) *openapi3.SchemaRef {
	if len(ref.Value.AllOf) > 0 {
		ref.Value.Type = "object"
		if len(ref.Value.AllOf) == 1 {
			if ref.Value.AllOf[0].Ref != "" {
				ref.Ref = ref.Value.AllOf[0].Ref
			}
		}
		ref.Value.Properties = make(map[string]*openapi3.SchemaRef)
		for _, subSpec := range ref.Value.AllOf {
			for propName, propSpec := range subSpec.Value.Properties {
				ref.Value.Properties[propName] = propSpec
			}
			// Here we bubble up the additionalProperties if we find it in any of the allof entrieref.
			// If we find it, we delete all collected property information and will return an map[string]interface{}
			if subSpec.Value.AdditionalPropertiesAllowed != nil && *subSpec.Value.AdditionalPropertiesAllowed {
				ref.Value.AdditionalPropertiesAllowed = subSpec.Value.AdditionalPropertiesAllowed
			}
		}
		if ref.Value.AdditionalPropertiesAllowed != nil && *ref.Value.AdditionalPropertiesAllowed {
			ref.Value.Properties = nil
		}
		ref.Value.AllOf = nil
	}
	return ref
}

func structPropsFromRef(ref *openapi3.SchemaRef) (specs []PropSpec, imports []string, err error) {
	for _, name := range sortedKeys(ref.Value.Properties) {
		prop := ref.Value.Properties[name]
		prop = resolveAllOf(prop)

		spec := PropSpec{
			Name:        tpl.ToPascalCase(name),
			Description: prop.Value.Description,
			GoType:      goTypeFromSpec(prop),
			IsRequired:  checkIfRequired(name, ref.Value.Required),
			IsEnum:      len(prop.Value.Enum) > 0,
			IsNullable:  prop.Value.Nullable,
		}

		if spec.GoType == "time.Time" || spec.GoType == "*time.Time" {
			imports = append(imports, "time")
		}

		spec.JSONTags = "`json:\"" + name
		if !spec.IsRequired {
			spec.JSONTags += ",omitempty"
		}
		spec.JSONTags += "\"`"

		imports = append(imports, fillValidationRelatedProperties(prop, &spec)...)

		specs = append(specs, spec)
	}
	return specs, uniqueStrings(imports), err
}

func fillValidationRelatedProperties(ref *openapi3.SchemaRef, spec *PropSpec) (imports []string) {
	if (len(spec.GoType) > 0 && spec.GoType[0] >= 'A' && spec.GoType[0] <= 'Z') ||
		(len(spec.GoType) > 1 && spec.GoType[:2] == "[]") ||
		(len(spec.GoType) > 2 && spec.GoType[:3] == "map") {
		// enable recursive validation
		spec.NeedsValidation = true
		spec.IsRequiredInValidation = !spec.IsNullable && spec.IsRequired
	}

	if ref.Value.Min != nil {
		spec.NeedsValidation = true
		spec.HasMin = true
		spec.Min = *ref.Value.Min
		if spec.Min > 0 {
			spec.IsRequiredInValidation = !spec.IsNullable && spec.IsRequired
		}
	}
	if ref.Value.Max != nil {
		spec.NeedsValidation = true
		spec.HasMax = true
		spec.Max = *ref.Value.Max
		if spec.Max > 0 {
			spec.IsRequiredInValidation = !spec.IsNullable && spec.IsRequired
		}
	}
	if ref.Value.MinLength > 0 {
		spec.NeedsValidation = true
		spec.HasMinLength = true
		spec.MinLength = ref.Value.MinLength
		if spec.MinLength > 0 {
			spec.IsRequiredInValidation = !spec.IsNullable && spec.IsRequired
		}
	}
	if ref.Value.MaxLength != nil {
		spec.NeedsValidation = true
		spec.HasMaxLength = true
		spec.MaxLength = *ref.Value.MaxLength
		if spec.MaxLength > 0 {
			spec.IsRequiredInValidation = !spec.IsNullable && spec.IsRequired
		}
	}
	if ref.Value.ExclusiveMin {
		spec.Min += math.SmallestNonzeroFloat64
	}
	if ref.Value.ExclusiveMax {
		spec.Min -= math.SmallestNonzeroFloat64
	}
	switch ref.Value.Format {
	case "date":
		spec.IsRequiredInValidation = !spec.IsNullable && spec.IsRequired
		spec.IsRequired = true
		spec.IsDate = true
	case "date-time":
		spec.IsRequiredInValidation = !spec.IsNullable && spec.IsRequired
		spec.IsRequired = true
		spec.IsDateTime = true
		imports = append(imports, "time")
	case "byte":
		spec.IsRequiredInValidation = !spec.IsNullable && spec.IsRequired
		spec.NeedsValidation = true
		spec.IsBase64 = true
		imports = append(imports, "github.com/go-ozzo/ozzo-validation/v4/is")
	case "email":
		spec.IsRequiredInValidation = !spec.IsNullable && spec.IsRequired
		spec.NeedsValidation = true
		spec.IsEmail = true
		imports = append(imports, "github.com/go-ozzo/ozzo-validation/v4/is")
	case "uuid":
		spec.IsRequiredInValidation = !spec.IsNullable && spec.IsRequired
		spec.NeedsValidation = true
		spec.IsUUID = true
		imports = append(imports, "github.com/go-ozzo/ozzo-validation/v4/is")
	case "url":
		spec.IsRequiredInValidation = !spec.IsNullable && spec.IsRequired
		spec.NeedsValidation = true
		spec.IsURL = true
		imports = append(imports, "github.com/go-ozzo/ozzo-validation/v4/is")
	case "hostname":
		spec.IsRequiredInValidation = !spec.IsNullable && spec.IsRequired
		spec.NeedsValidation = true
		spec.IsHostname = true
		imports = append(imports, "github.com/go-ozzo/ozzo-validation/v4/is")
	case "ipv4":
		spec.IsRequiredInValidation = !spec.IsNullable && spec.IsRequired
		spec.NeedsValidation = true
		spec.IsIPv4 = true
		imports = append(imports, "github.com/go-ozzo/ozzo-validation/v4/is")
	case "ipv6":
		spec.IsRequiredInValidation = !spec.IsNullable && spec.IsRequired
		spec.NeedsValidation = true
		spec.IsIPv6 = true
		imports = append(imports, "github.com/go-ozzo/ozzo-validation/v4/is")
	case "ip":
		spec.IsRequiredInValidation = !spec.IsNullable && spec.IsRequired
		spec.NeedsValidation = true
		spec.IsIP = true
		imports = append(imports, "github.com/go-ozzo/ozzo-validation/v4/is")
	case "", "int32", "int64", "float", "double":
		break // do nothing
	default:
		log.Warn().Str("name", spec.Name).Str("format", ref.Value.Format).Msg("unknown format")
	}
	return imports
}

func enumPropsFromRef(ref *openapi3.SchemaRef, model *Model) (specs []PropSpec, err error) {
	for _, val := range ref.Value.Enum {
		valueVarName := tpl.ToPascalCase(tpl.FirstUpper(fmt.Sprintf("%v", val)))
		specs = append(specs, PropSpec{
			Name:   valueVarName,
			Value:  fmt.Sprintf(`"%v"`, val),
			GoType: model.Name,
		})
	}
	sort.Slice(specs, func(i, j int) bool {
		return specs[i].Name < specs[j].Name
	})
	return specs, err
}

func checkIfRequired(name string, list []string) bool {
	for _, n := range list {
		if n == name {
			return true
		}
	}
	return false
}

func uniqueStrings(input []string) (output []string) {
	m := make(map[string]struct{})
	for _, item := range input {
		if _, ok := m[item]; ok {
			continue
		}
		output = append(output, item)
		m[item] = struct{}{}
	}
	return output
}
