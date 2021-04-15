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

type something struct{}
type passedSchemas map[*openapi3.SchemaRef]something

type ModelKind string

const (
	Struct ModelKind = "struct"
	Enum   ModelKind = "enum"
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
	Name        string // Property name of structs, variable name of enumns
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
	IsURI                      bool
	IsRequestURI               bool
	IsHostname                 bool
	IsIP                       bool
	IsIPv4                     bool
	IsIPv6                     bool
}

func NewModelFromRef(ref *openapi3.SchemaRef) (model *Model, err error) {
	model = &Model{
		Description: ref.Value.Description,
	}

	ref = resolveAllOf(ref, nil)

	switch {
	case len(ref.Value.Enum) > 0:
		model.Kind = Enum
		model.Properties, err = enumPropsFromRef(ref, model)
		model.GoType = goTypeFromSpec(ref)
	case ref.Value.Type == "object" ||
		len(ref.Value.Properties) > 0 ||
		len(ref.Value.AllOf) > 0 ||
		len(ref.Value.OneOf) > 0:
		model.Kind = Struct
		model.Properties, model.Imports, err = structPropsFromRef(ref)
		if len(model.Properties) == 0 {
			model.GoType = "map[string]interface{}"
			if ref.Value.AdditionalProperties != nil {
				model.GoType = "map[string]" + goTypeFromSpec(ref.Value.AdditionalProperties)
			}
		}
	default:
		return nil, errors.New("type not handled")
	}
	if err != nil {
		return nil, err
	}

	return model, nil
}

func NewModelFromParameters(params openapi3.Parameters) (model *Model, err error) {
	model = &Model{
		Kind: Struct,
	}
	for _, param := range params {
		spec := PropSpec{
			Name:        tpl.ToPascalCase(param.Value.Name),
			Description: param.Value.Description,
			GoType:      goTypeFromSpec(param.Value.Schema),
			IsRequired:  param.Value.Required,
		}

		if spec.GoType == "time.Time" || spec.GoType == "*time.Time" {
			model.Imports = append(model.Imports, "time")
		}

		model.Imports = append(model.Imports, fillValidationRelatedProperties(param.Value.Schema, &spec)...)

		spec.JSONTags = "`json:\"" + param.Value.Name
		if !spec.IsRequired {
			spec.JSONTags += ",omitempty"
		}
		spec.JSONTags += "\"`"

		model.Properties = append(model.Properties, spec)
	}

	model.Imports = uniqueStrings(model.Imports)
	return model, nil
}

func (m *Model) Render(ctx context.Context, writer io.Writer) error {
	var tpl *template.Template
	switch m.Kind {
	case Struct:
		tpl = modelTemplate
	case Enum:
		tpl = enumTemplate
	}
	err := tpl.Execute(writer, m)
	if err != nil {
		return errors.Wrap(err, "failed to render model")
	}
	return nil
}

func resolveAllOf(ref *openapi3.SchemaRef, passed passedSchemas) (out *openapi3.SchemaRef) {
	if ref == nil {
		ref = &openapi3.SchemaRef{}
	}
	if ref.Value == nil {
		ref.Value = &openapi3.Schema{}
	}

	out = ref

	if passed == nil {
		passed = make(passedSchemas)
	}

	for _, subSchema := range ref.Value.AllOf {
		if _, exists := passed[subSchema]; exists {
			continue
		}
		passed[subSchema] = something{}
		out = deepMerge(out, subSchema, passed)
	}

	return out
}

func deepMerge(left *openapi3.SchemaRef, right *openapi3.SchemaRef, passed passedSchemas) (out *openapi3.SchemaRef) {
	out = left

	if out == nil {
		out = &openapi3.SchemaRef{}
	}
	if out.Value == nil {
		out.Value = &openapi3.Schema{}
	}

	if right == nil {
		right = &openapi3.SchemaRef{}
	}
	if right.Value == nil {
		right.Value = &openapi3.Schema{}
	}

	if right.Ref != "" {
		out.Ref = right.Ref
	}
	if right.Value.Type != "" {
		out.Value.Type = right.Value.Type
	}

	// merge docs
	if right.Value.Title != "" {
		out.Value.Title = right.Value.Title
	}
	if right.Value.Description != "" {
		out.Value.Description = right.Value.Description
	}
	if right.Value.ExternalDocs != nil {
		out.Value.ExternalDocs = right.Value.ExternalDocs
	}
	if right.Value.Example != nil {
		out.Value.Example = right.Value.Example
	}

	// override main type properties
	if right.Value.Format != "" {
		out.Value.Format = right.Value.Format
	}
	if right.Value.Nullable != out.Value.Nullable {
		out.Value.Nullable = true
	}
	if right.Value.AdditionalPropertiesAllowed != nil {
		out.Value.AdditionalPropertiesAllowed = right.Value.AdditionalPropertiesAllowed
	}

	out.Value.OneOf = append(out.Value.OneOf, right.Value.OneOf...)
	out.Value.AllOf = append(out.Value.AllOf, right.Value.AllOf...)
	out.Value.AnyOf = append(out.Value.AnyOf, right.Value.AnyOf...)
	out.Value.Enum = append(out.Value.Enum, right.Value.Enum...)
	out.Value.Required = append(out.Value.Required, right.Value.Required...)

	if right.Value.AdditionalProperties != nil {
		out.Value.AdditionalProperties = deepMerge(
			out.Value.AdditionalProperties,
			right.Value.AdditionalProperties,
			nil, // here we merge all over, without `passed`
		)
	}

	if len(right.Value.Properties) > 0 {
		if out.Value.Properties == nil {
			out.Value.Properties = make(map[string]*openapi3.SchemaRef)
		}
		for k, v := range right.Value.Properties {
			out.Value.Properties[k] = resolveAllOf(v, passed)
		}
	}

	if right.Value.Type == "array" && right.Value.Items != nil {
		out.Value.Items = resolveAllOf(right.Value.Items, passed)
	}

	return out
}

func structPropsFromRef(ref *openapi3.SchemaRef) (specs []PropSpec, imports []string, err error) {
	for _, name := range sortedKeys(ref.Value.Properties) {
		prop := ref.Value.Properties[name]
		prop = resolveAllOf(prop, nil)

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
	case "uri":
		spec.IsRequiredInValidation = !spec.IsNullable && spec.IsRequired
		spec.NeedsValidation = true
		spec.IsURI = true
		imports = append(imports, "github.com/go-ozzo/ozzo-validation/v4/is")
	case "request-uri":
		spec.IsRequiredInValidation = !spec.IsNullable && spec.IsRequired
		spec.NeedsValidation = true
		spec.IsRequestURI = true
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
	for idx, val := range ref.Value.Enum {
		valueVarName := tpl.ToPascalCase(fmt.Sprintf("%v", val))
		if valueVarName == "" {
			valueVarName = fmt.Sprintf("Item%d", idx)
		}
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
