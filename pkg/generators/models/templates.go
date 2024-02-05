package models

import (
	"embed"
	"text/template"

	tpl "github.com/contiamo/openapi-generator-go/v2/pkg/generators/templates"
)

var (
	//go:embed templates
	templateData embed.FS

	modelTpls     = template.Must(template.New("base").Funcs(fmap).ParseFS(templateData, "templates/*.gotpl"))
	modelTemplate = modelTpls.Lookup("model.gotpl")
	oneOfTemplate = modelTpls.Lookup("oneof.gotpl")
	enumTemplate  = modelTpls.Lookup("enum.gotpl")
	valueTemplate = modelTpls.Lookup("value.gotpl")

	fmap = template.FuncMap{
		"firstLower":      tpl.FirstLower,
		"firstUpper":      tpl.FirstUpper,
		"commentBlock":    tpl.CommentBlock,
		"toPascalCase":    tpl.ToPascalCase,
		"toSnakeCase":     tpl.ToSnakeCase,
		"removeSpecial":   tpl.RemoveSpecial,
		"typeDisplayName": tpl.TypeDisplayName,
		"ternary":         tpl.Ternary,
	}
)
