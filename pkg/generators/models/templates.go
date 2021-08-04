package models

import (
	"text/template"

	tpl "github.com/contiamo/openapi-generator-go/pkg/generators/templates"
)

var (
	modelTemplate = template.Must(
		template.New("model").
			Funcs(fmap).
			Parse(modelTemplateSource),
	)

	enumTemplate = template.Must(
		template.New("enum").
			Funcs(fmap).
			Parse(enumTemplateSource),
	)

	valueTemplate = template.Must(
		template.New("value").
			Funcs(fmap).
			Parse(valueTemplateSource),
	)

	oneOfTemplate = template.Must(
		template.New("oneof").
			Funcs(fmap).
			Parse(oneOfTemplateSource),
	)

	fmap = template.FuncMap{
		"firstLower":   tpl.FirstLower,
		"firstUpper":   tpl.FirstUpper,
		"commentBlock": tpl.CommentBlock,
		"toPascalCase": tpl.ToPascalCase,
		"toSnakeCase":  tpl.ToSnakeCase,
	}

	modelTemplateSource = `// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: {{.SpecTitle}}
//     Version: {{.SpecVersion}}
package {{ .PackageName }}

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
{{- range .Imports}}
	"{{.}}"
{{ end}}
)

{{ (printf "%s is an object. %s" .Name .Description) | commentBlock }}
{{- if not .Properties }}
type {{.Name}} {{ .GoType }}
{{- else }}
type {{.Name}} struct {
{{- range .Properties}}
	{{ (printf "%s: %s" .Name .Description) | commentBlock }}
	{{.Name}} {{.GoType}} {{.JSONTags}}
{{- end}}
}
{{- end}}

{{- $modelName := .Name }}
// Validate implements basic validation for this model
func (m {{$modelName}}) Validate() error {
	return validation.Errors{
		{{- range .Properties}}
			{{- if .NeedsValidation }}
				"{{ firstLower .Name }}": validation.Validate(
					m.{{ .Name }},
					{{- if and .IsRequiredInValidation}}
						{{- if or .IsEnum .Min .MinLength .HasFormat}}validation.Required,{{ else }}validation.NotNil,{{- end }}
					{{- end }}
					{{- if .HasMin }}validation.Min({{ .GoType }}({{ .Min }})),{{ end }}
					{{- if .HasMax }}validation.Max({{ .GoType }}({{ .Max }})),{{ end }}
					{{- if or .HasMinLength .HasMaxLength }}validation.Length({{ .MinLength }},{{ .MaxLength }}),{{ end }}
					{{- if .IsDate }}validation.Date("2006-01-02"),{{ end }}
					{{- if .IsDateTime }}validation.Date(time.RFC3339),{{ end }}
					{{- if .IsBase64 }}is.Base64,{{ end }}
					{{- if .IsEmail }}is.EmailFormat,{{ end }}
					{{- if .IsUUID }}is.UUID,{{ end }}
					{{- if .IsURL }}is.URL.Error("must be a valid URL with HTTP or HTTPS scheme"),{{ end }}
					{{- if .IsURI }}is.RequestURI,{{ end }}
					{{- if .IsRequestURI }}is.RequestURL.Error("must be valid URI with scheme"),{{ end }}
					{{- if .IsHostname }}is.Host,{{ end }}
					{{- if .IsIPv4 }}is.IPv4,{{ end }}
					{{- if .IsIPv6 }}is.IPv6,{{ end }}
					{{- if .IsIP }}is.IP,{{ end }}
				),
			{{- end }}
		{{- end }}
	}.Filter()
}
{{ range .Properties}}
// Get{{.Name}} returns the {{.Name}} property
func (m {{$modelName}}) Get{{.Name}}() {{.GoType}} {
	return m.{{.Name}}
}

// Set{{.Name}} sets the {{.Name}} property
func (m *{{$modelName}}) Set{{.Name}}(val {{.GoType}}) {
	m.{{.Name}} = val
}
{{ end}}`

	enumTemplateSource = `
// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: {{.SpecTitle}}
//     Version: {{.SpecVersion}}
package {{ .PackageName }}

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

{{ (printf "%s is an enum. %s" .Name .Description) | commentBlock }}
type {{.Name}} {{.GoType}}

// Validate implements basic validation for this model
func (m {{.Name}}) Validate() error {
	return InKnown{{.Name}}.Validate(m)
}

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

	oneOfTemplateSource = `
// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: {{.SpecTitle}}
//     Version: {{.SpecVersion}}
package {{ .PackageName }}

import (
	"encoding/json"
	"github.com/mitchellh/mapstructure"
)

{{ (printf "%s is a oneOf type. %s" .Name .Description) | commentBlock }}
type {{.Name}} struct {
	data interface{}
}

{{- $modelName := .Name }}

// MarshalJSON implementes the json.Marshaller interface
func (m {{$modelName}}) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.data)
}

// UnmarshalJSON implementes the json.Unmarshaller interface
func (m *{{$modelName}}) UnmarshalJSON(bs []byte) error {
	return json.Unmarshal(bs, &m.data)
}


{{- range $convert := .ConvertSpecs }}
// From{{firstUpper $convert.TargetGoType}} sets the {{$modelName}} data.
func (m *{{$modelName}}) From{{firstUpper $convert.TargetGoType}}(data {{$convert.TargetGoType}}) {
	m.data = data
}

{{- end }}

// As converts {{$modelName}} to a user defined structure.
func (m {{$modelName}}) As(target interface{}) error {
	return mapstructure.Decode(m.data, target)
}

{{- range $convert := .ConvertSpecs }}
// As{{firstUpper $convert.TargetGoType}} converts {{$modelName}} to a {{$convert.TargetGoType}}
func (m {{$modelName}}) As{{firstUpper $convert.TargetGoType}}() (result {{$convert.TargetGoType}}, err error) {
	return result, mapstructure.Decode(m.data, &result)
}

{{- end}}
`

	valueTemplateSource = `
// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: {{.SpecTitle}}
//     Version: {{.SpecVersion}}
package {{ .PackageName }}

{{ (printf "%s is a value type. %s" .Name .Description) | commentBlock }}
type {{.Name}} {{.GoType}}
`
)
