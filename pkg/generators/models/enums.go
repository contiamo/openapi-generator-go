package models

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"text/template"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/sirupsen/logrus"

	tpl "github.com/contiamo/go-base/v2/pkg/generators/templates"
)

// DefaultPackageName used in the models source code
const DefaultPackageName = "openapi"

// Options represent all the possible options of the generator
type Options struct {
	// PackageName of the generated models source code (`DefaultPackageName` by default)
	PackageName string
}

// GenerateEnums outputs the Go enum models with validators
func GenerateEnums(specFile io.Reader, dst string, opts Options) error {
	if opts.PackageName == "" {
		opts.PackageName = DefaultPackageName
	}

	data, err := ioutil.ReadAll(specFile)
	if err != nil {
		return fmt.Errorf("can not read spec file: %w", err)
	}
	swagger, err := openapi3.NewSwaggerLoader().LoadSwaggerFromData(data)
	if err != nil {
		return fmt.Errorf("can not parse the OpenAPI spec: %w", err)
	}

	enums := enumCtxs{}

	// to do sort and iterate over the sorted schema
	for name, s := range swagger.Components.Schemas {
		if len(s.Value.Enum) == 0 {
			continue
		}

		tctx := templateCtx{
			SpecTitle:   swagger.Info.Title,
			SpecVersion: swagger.Info.Version,
			PackageName: opts.PackageName,
			Name:        name,
			VarName:     tpl.ToPascalCase(name),
			Description: s.Value.Description,
			Type:        s.Value.Type,
			Values:      enumValues{},
		}

		for idx, v := range s.Value.Enum {
			valstr := tpl.FirstUpper(fmt.Sprintf("%v", v))
			valueVarName := tpl.ToPascalCase(fmt.Sprintf("%s %s", tctx.VarName, valstr))

			if valueVarName == tctx.VarName {
				valueVarName = tpl.ToPascalCase(fmt.Sprintf("%s Value %d", tctx.VarName, idx))
			}

			value := enumValue{
				Name:    valueVarName,
				VarName: valueVarName,
				Value:   fmt.Sprintf("%q", v),
			}
			tctx.Values = append(tctx.Values, value)
		}

		sort.Sort(tctx.Values)
		enums = append(enums, tctx)
	}

	sort.Sort(enums)
	for _, tctx := range enums {
		modelName := strings.ToLower(strings.ReplaceAll(fmt.Sprintf("model_%s.go", tpl.ToSnakeCase(tctx.Name)), " ", "_"))
		filename := filepath.Join(dst, modelName)
		logrus.Debugf("writing %s\n", filename)
		f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
		if err != nil {
			return err
		}

		if len(tctx.Values) == 1 {
			err = constTemplate.Execute(f, tctx)
		} else {
			err = enumTemplate.Execute(f, tctx)
		}
		if err != nil {
			return fmt.Errorf("failed to generate enum code: %w", err)
		}

		err = f.Close()
		if err != nil {
			return fmt.Errorf("failed to close output file: %w", err)
		}
	}
	return nil
}

type templateCtx struct {
	Filename    string
	SpecTitle   string
	SpecVersion string
	PackageName string
	Name        string
	VarName     string
	Type        string
	Description string
	Values      enumValues
}

type enumCtxs []templateCtx

func (e enumCtxs) Len() int           { return len(e) }
func (e enumCtxs) Swap(i, j int)      { e[i], e[j] = e[j], e[i] }
func (e enumCtxs) Less(i, j int) bool { return e[i].Name < e[j].Name }

type enumValue struct {
	Name    string
	VarName string
	Value   string
}

type enumValues []enumValue

func (e enumValues) Len() int           { return len(e) }
func (e enumValues) Swap(i, j int)      { e[i], e[j] = e[j], e[i] }
func (e enumValues) Less(i, j int) bool { return e[i].Name < e[j].Name }

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

var fmap = template.FuncMap{
	"firstLower":   tpl.FirstLower,
	"firstUpper":   tpl.FirstUpper,
	"commentBlock": tpl.CommentBlock,
	"toPascalCase": tpl.ToPascalCase,
	"toSnakeCase":  tpl.ToSnakeCase,
}

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
type {{.VarName}} {{.Type}}

var (
	{{- $enum :=. }}
	{{- range $v := .Values }}
	{{$v.VarName }} {{$enum.VarName}} = {{$v.Value}}
	{{- end}}

	// Known{{.VarName}} is the list of valid {{.VarName}}
	Known{{.VarName }} = []{{.VarName}}{
		{{- range $v := .Values }}
		{{$v.VarName}},
		{{- end}}
	}

	{{- $enum :=. }}
	// Known{{$enum.VarName}}{{$enum.Type | firstUpper}} is the list of valid {{$enum.VarName}} as {{$enum.Type}}
	Known{{$enum.VarName}}{{$enum.Type | firstUpper}} = []{{$enum.Type}}{
		{{- range $v := .Values }}
		{{$enum.Type}}({{$v.VarName}}),
		{{- end}}
	}

	// InKnown{{.VarName}} is an ozzo-validator for {{.VarName}}
	InKnown{{.VarName}} = validation.In(
		{{- range $v := .Values }}
		{{$v.VarName}},
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
const {{.VarName}} = {{ (index .Values 0).Value}}
`
