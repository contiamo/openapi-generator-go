package models

import (
	"path/filepath"
	"sort"
	"strings"

	"github.com/rs/zerolog/log"

	"github.com/getkin/kin-openapi/openapi3"

	tpl "github.com/contiamo/openapi-generator-go/pkg/generators/templates"
)

func goTypeFromSpec(schemaRef *openapi3.SchemaRef) string {
	if schemaRef == nil {
		log.Fatal().Msg("got nil schema ref")
	}
	// add missing object types
	if len(schemaRef.Value.Properties) > 0 {
		schemaRef.Value.Type = "object"
	}
	schema := schemaRef.Value
	propertyType := schemaRef.Value.Type
	switch propertyType {
	case "object":
		propertyType = goTypeForObject(schemaRef)
	case "string":
		if schema.Format == "date-time" || schema.Format == "time" {
			propertyType = "time.Time"
		}
		if len(schema.Enum) > 0 && schemaRef.Ref != "" {
			propertyType = filepath.Base(schemaRef.Ref)
		}
	case "array":
		subType := "interface{}"
		if schema.Items != nil {
			subType = goTypeFromSpec(schema.Items)
		}
		propertyType = "[]" + subType
	case "boolean":
		propertyType = "bool"
	case "integer":
		switch schemaRef.Value.Format {
		case "int32":
			propertyType = "int32"
		case "int64":
			propertyType = "int64"
		default:
			propertyType = "int32" // the default is int32 because this is always json-serialized to a proper number in contrast to int64 which will become a string
		}
	case "number":
		switch schemaRef.Value.Format {
		case "float":
			propertyType = "float32"
		case "double":
			propertyType = "float64"
		default:
			propertyType = "float32" // the default is float32 because this is always json-serialized to a proper number in contrast to float64 which will become a string
		}
	case "":
		if schemaRef.Ref != "" {
			propertyType = filepath.Base(schemaRef.Ref)
		} else {
			propertyType = "interface{}"
		}
	}
	if schema.Nullable && !strings.HasPrefix(propertyType, "[]") && !strings.HasPrefix(propertyType, "map[") {
		propertyType = "*" + propertyType
	}
	return propertyType
}

func goTypeForObject(schemaRef *openapi3.SchemaRef) (propType string) {
	switch {
	case schemaRef.Ref != "":
		propType = filepath.Base(schemaRef.Ref)
	case schemaRef.Value.AdditionalProperties != nil:
		subType := goTypeFromSpec(schemaRef.Value.AdditionalProperties)
		propType = "map[string]" + subType
	case schemaRef.Value.AdditionalPropertiesAllowed != nil && *schemaRef.Value.AdditionalPropertiesAllowed:
		propType = "map[string]interface{}"
	case len(schemaRef.Value.Properties) > 0:
		structBuilder := &strings.Builder{}
		structBuilder.WriteString("struct {\n")
		for _, name := range sortedKeys(schemaRef.Value.Properties) {
			ref := schemaRef.Value.Properties[name]
			propName := tpl.ToPascalCase(name)
			omitEmpty := true
			for _, required := range ref.Value.Required {
				if required == propName {
					omitEmpty = false
					break
				}
			}
			jsonTags := "`json:\"" + name
			if omitEmpty {
				jsonTags += ",omitempty"
			}
			jsonTags += "\"`"
			structBuilder.WriteString(tpl.ToPascalCase(name))
			structBuilder.WriteString(" ")
			structBuilder.WriteString(goTypeFromSpec(ref))
			structBuilder.WriteString(jsonTags)
			structBuilder.WriteString("\n")
		}
		structBuilder.WriteString("}")
		propType = structBuilder.String()
	case len(schemaRef.Value.OneOf) > 0:
		return "interface{}"
	default:
		return "map[string]interface{}"
	}
	return propType
}

func sortedKeys(obj map[string]*openapi3.SchemaRef) (res []string) {
	for k := range obj {
		res = append(res, k)
	}
	sort.Strings(res)
	return res
}
