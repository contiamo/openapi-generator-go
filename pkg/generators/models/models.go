package models

import (
	"context"
	"fmt"
	"io"
	"sort"
	"text/template"

	tpl "github.com/contiamo/openapi-generator-go/pkg/generators/templates"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/pkg/errors"
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
	AsPointer   bool
	Value       string
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
	default:
		model.Kind = Struct
		model.Properties, model.Imports, err = structPropsFromRef(ref)
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
	props := ref.Value.Properties
	for _, r := range ref.Value.AllOf {
		for k, v := range r.Value.Properties {
			props[k] = v
		}
	}
	timeImportRequired := false
	for _, name := range sortedKeys(ref.Value.Properties) {
		prop := ref.Value.Properties[name]
		prop = resolveAllOf(prop)
		goType := goTypeFromSpec(prop)
		if goType == "time.Time" {
			timeImportRequired = true
		}
		omitEmpty := true
		for _, required := range ref.Value.Required {
			if required == name {
				omitEmpty = false
				break
			}
		}
		jsonTags := "`json:\"" + name
		if omitEmpty {
			jsonTags += ",omitempty"
		}
		jsonTags += "\"`"

		specs = append(specs, PropSpec{
			Name:        tpl.ToPascalCase(name),
			GoType:      goType,
			JSONTags:    jsonTags,
			Description: prop.Value.Description,
			AsPointer:   !checkIfRequired(name, ref.Value.Required),
		})
	}
	if timeImportRequired {
		imports = []string{"time"}
	}
	return specs, imports, err
}

func enumPropsFromRef(ref *openapi3.SchemaRef, model *Model) (specs []PropSpec, err error) {
	for _, val := range ref.Value.Enum {
		valueVarName := tpl.FirstUpper(fmt.Sprintf("%v", val))
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

var modelTemplateSource = `// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: {{.SpecTitle}}
//     Version: {{.SpecVersion}}
package {{ .PackageName }}

{{ if .Imports }}import ({{end}}
{{- range .Imports}}
	"{{.}}"
{{ end}}
{{- if .Imports }}){{end}}

{{ (printf "%s is an object. %s" .Name .Description) | commentBlock }}
{{- if not .Properties }}
type {{.Name}} map[string]interface{}
{{- else }}
type {{.Name}} struct {
{{- range .Properties}}
	{{ (printf "%s: %s" .Name .Description) | commentBlock }}
	{{.Name}} {{.GoType}} {{.JSONTags}}
{{- end}}
}
{{- end}}

{{- $modelName := .Name }}
{{ range .Properties}}
// Get{{.Name}} returns the {{.Name}} property
func (m {{$modelName}}) Get{{.Name}}() {{.GoType}} {
	return m.{{.Name}}
}

// Set{{.Name}} sets the {{.Name}} property
func (m {{$modelName}}) Set{{.Name}}(val {{.GoType}}) {
	m.{{.Name}} = val
}
{{ end}}`

var enumTemplateSource = `
// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: {{.SpecTitle}}
//     Version: {{.SpecVersion}}
package {{ .PackageName }}

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

{{ (printf "%s is an enum. %s" .Name .Description) | commentBlock }}
type {{.Name}} {{.GoType}}

var (
	{{- $enum :=. }}
	{{- range $v := .Properties }}
	{{$enum.Name}}{{$v.Name}} {{$enum.Name}} = {{$v.Value}}
	{{- end}}

	// Known{{.Name}} is the list of valid {{.Name}}
	Known{{.Name }} = []{{.Name}}{
		{{- range $v := .Properties }}
		{{$enum.Name}}{{$v.Name}},
		{{- end}}
	}

	{{- $enum :=. }}
	// Known{{$enum.Name}}{{$enum.GoType | firstUpper}} is the list of valid {{$enum.Name}} as {{$enum.GoType}}
	Known{{$enum.Name}}{{$enum.GoType | firstUpper}} = []{{$enum.GoType}}{
		{{- range $v := .Properties }}
		{{$enum.GoType}}({{$enum.Name}}{{$v.Name}}),
		{{- end}}
	}

	// InKnown{{.Name}} is an ozzo-validator for {{.Name}}
	InKnown{{.Name}} = validation.In(
		{{- range $v := .Properties }}
		{{$enum.Name}}{{$v.Name}},
		{{- end}}
	)
)
`

var constTemplateSource = `
// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: {{.SpecTitle}}
//     Version: {{.SpecVersion}}
package {{ .PackageName }}

{{ (printf "%s is an enum. %s" .Name .Description) | commentBlock }}
const {{.Name}} = {{ (index .Properties 0).Value}}
`

var fmap = template.FuncMap{
	"firstLower":   tpl.FirstLower,
	"firstUpper":   tpl.FirstUpper,
	"commentBlock": tpl.CommentBlock,
	"toPascalCase": tpl.ToPascalCase,
	"toSnakeCase":  tpl.ToSnakeCase,
}

var modelTemplate = template.Must(
	template.New("model").
		Funcs(fmap).
		Parse(modelTemplateSource),
)
var enumTemplate = template.Must(
	template.New("enum").
		Funcs(fmap).
		Parse(enumTemplateSource),
)
var constTemplate = template.Must(
	template.New("const").
		Funcs(fmap).
		Parse(constTemplateSource),
)
