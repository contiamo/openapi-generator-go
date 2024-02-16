package models

import (
	"fmt"
	"path/filepath"
	"sort"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"

	tpl "github.com/contiamo/openapi-generator-go/v2/pkg/generators/templates"
)

// goTypeFromSpec gets a Go type that can represent the schema
// For example, `type Name GoType` or `type Example string`.
func goTypeFromSpec(schemaRef *openapi3.SchemaRef) string {
	if schemaRef == nil {
		schemaRef = &openapi3.SchemaRef{}
	}
	if schemaRef.Value == nil {
		schemaRef.Value = &openapi3.Schema{}
	}

	// add missing object types
	if len(schemaRef.Value.Properties) > 0 {
		schemaRef.Value.Type = "object"
	}

	schema := schemaRef.Value
	propertyType := schemaRef.Value.Type
	// special case for references, we prefer named types,
	// except for arrays, we don't want to create an reference slice aliases
	if schemaRef.Ref != "" && propertyType != "array" {
		propertyType = "ref"
	}

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
			// the default is int32 because this is always json-serialized
			// to a proper number in contrast to int64 which will become a string
			propertyType = "int32"
		}

	case "number":
		switch schemaRef.Value.Format {
		case "float":
			propertyType = "float32"
		case "double":
			propertyType = "float64"
		default:
			// the default is float32 because this is always json-serialized
			// to a proper number in contrast to float64 which will become a string
			propertyType = "float32"
		}

	case "ref":
		propertyType = typeNameFromRef(schemaRef.Ref)
	case "":
		propertyType = "interface{}"
	}

	return propertyType
}

// goTypeForObject returns a string that can represent the entire object type of the given schema
func goTypeForObject(schemaRef *openapi3.SchemaRef) (propType string) {
	switch {
	case schemaRef.Ref != "":
		propType = typeNameFromRef(schemaRef.Ref)

	case schemaRef.Value.AdditionalProperties != nil:
		subType := goTypeFromSpec(schemaRef.Value.AdditionalProperties)
		propType = "map[string]" + subType

	case schemaRef.Value.AdditionalPropertiesAllowed != nil &&
		*schemaRef.Value.AdditionalPropertiesAllowed:
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

			omit := ""
			if omitEmpty {
				omit += ",omitempty"
			}

			jsonTags := fmt.Sprintf(" `json:\"%s%s\"", name, omit)
			jsonTags += fmt.Sprintf(" mapstructure:\"%s%s\"", name, omit)
			jsonTags += "`"

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

var typeNameFromRef = func(ref string) string {
	return tpl.ToPascalCase(filepath.Base(ref))
}

func sortedKeys(obj map[string]*openapi3.SchemaRef) (res []string) {
	for k := range obj {
		res = append(res, k)
	}
	sort.Strings(res)
	return res
}
