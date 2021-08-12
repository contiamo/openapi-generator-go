package models

import (
	"context"
	"fmt"
	"io"
	"math"
	"regexp"
	"sort"
	"text/template"

	tpl "github.com/contiamo/openapi-generator-go/pkg/generators/templates"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/pkg/errors"
)

var (
	// validatedTypesRegExp helps to match a type to a type that we can generate validation for.
	// Exported types may have a Validate() function
	// Exported types with a leading pointer may have a Validate() function
	// Slices may contain data that is validatable (they are also never pointers)
	// Maps may contain data that is validateable (they are also never pointers )
	validatedTypesRegExp = regexp.MustCompile(`^(\*?[A-Z])|(\[\])|(map\[)`)
)

// something is just something that does not have to be allocated
type something struct{}

// passedSchemas is a map of pointers to a schema.
// Used for marking already visited types when recursively resolve `allOf` definitions and merging types.
type passedSchemas map[*openapi3.SchemaRef]something

// ModelKind is a kind of a Go model to render
type ModelKind string

const (
	// Struct is a regular Go struct
	Struct ModelKind = "struct"
	// Enum is a Go enum definition
	Enum ModelKind = "enum"
	// Value is a value type
	Value ModelKind = "value"
	// OneOf is a oneof value
	OneOf ModelKind = "oneof"
)

// Model is a template model for rendering Go code for a given API schema
type Model struct {
	// Name is a name of the generated type that follows the `type` keyword in the definition
	// For example, `type Name GoType`.
	Name string
	// Description is a description that will become a comment on the generated type
	Description string
	// Imports is a list of imports that will be present in the Go file
	Imports []string
	// Kind is a kind of generated model (e.g. `struct` or `enum`)
	Kind ModelKind
	// Properties is a list of type's property descriptors
	Properties []PropSpec
	// ConvertSpecs contains a list of convert functions for this model
	ConvertSpecs []ConvertSpec
	// GoType is a string that represents the Go type that follows the model type name.
	// For example, `type Name GoType`.
	GoType string
	// SpecTitle is the spec title used in the auto-generated header
	SpecTitle string
	// SpecVersion is the spec version used in the auto-generated header
	SpecVersion string
	// PackageName is the name of the package used in the Go code
	PackageName string
}

// ConvertSpec holds all info to build one As{Type}() function
type ConvertSpec struct {
	// TargetGoType is the target type of the conversion
	TargetGoType string
}

// PropSpec is a Go property descriptor
type PropSpec struct {
	// Name is a property name in structs, variable name in enums, etc
	Name string
	// Description used in the comment of the property
	Description string
	// GoType used for this property (e.g. `string`, `int`, etc)
	GoType string
	// JSONTags is a string of JSON tags used for marshaling (e.g. `json:omitempty`)
	JSONTags string
	// Value is a value used for a enum item
	Value string
	// IsRequired is true when the property is a required value and should not be omitted on marshaling
	IsRequired bool
	// IsEnum is true when the property is a enum item
	IsEnum bool

	// Validation stuff

	IsNullable                 bool
	NeedsValidation            bool
	IsRequiredInValidation     bool
	HasMin, HasMax             bool
	Min, Max                   float64
	HasMinLength, HasMaxLength bool
	MinLength, MaxLength       uint64
	HasFormat                  bool
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

// NewModelFromRef creates a model out of a schema
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

	case ref.Value.Type == "object" || len(ref.Value.Properties) > 0:
		model.Kind = Struct
		model.Properties, model.Imports, err = structPropsFromRef(ref)
		if len(model.Properties) == 0 {
			model.GoType = "map[string]interface{}"
			if ref.Value.AdditionalProperties != nil {
				model.GoType = "map[string]" + goTypeFromSpec(ref.Value.AdditionalProperties)
			}
		}
	case len(ref.Value.OneOf) > 0:
		model.Kind = OneOf
		model.fillConvertSpecs(ref)
	default:
		model.Kind = Value
		model.GoType = goTypeFromSpec(ref)
	}
	if err != nil {
		return nil, err
	}

	return model, nil
}

// NewModelFromParameters returns a model built from operation parameters
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

func (m *Model) fillConvertSpecs(ref *openapi3.SchemaRef) {
	for _, oneOf := range ref.Value.OneOf {
		m.ConvertSpecs = append(m.ConvertSpecs, ConvertSpec{
			TargetGoType: goTypeFromSpec(oneOf),
		})
	}

}

// Render renders the model to a Go file
func (m *Model) Render(ctx context.Context, writer io.Writer) error {
	var tpl *template.Template

	switch m.Kind {
	case Struct:
		tpl = modelTemplate
	case Enum:
		tpl = enumTemplate
	case Value:
		tpl = valueTemplate
	case OneOf:
		tpl = oneOfTemplate
	}

	err := tpl.Execute(writer, m)
	if err != nil {
		return errors.Wrap(err, "failed to render the model")
	}

	return nil
}

// resolveAllOf resolves the list of `allOf` definitions in the schema to a complete type merging
// all the mentioned types.
// `passed` can be nil, it's used recursively in order to avoid infinite loops
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

	// very special case where the allOf references
	// another named schema but is not actually merging
	// several schemas, this is usually just overriding
	// a value like nullable or the description
	if len(ref.Value.AllOf) == 1 {
		pointer := ref.Value.AllOf[0]
		passed[pointer] = something{}
		out = deepMerge(out, pointer, passed)
		return out
	}

	// this is used for the special edge cases like this
	//   allOf:
	//      - description: "this will only set the description value"
	// 	- $ref: '#/components/schemas/ColumnTypeMetadata'
	//
	// without this method, the allOf will generate an inline type, but the resulting code
	// is much better and easier to use if we drop the first allOf item and just treat it like
	// the special case of a single `ref`, which will used the named `ColumnTypeMetadata` type.
	ref.Value.AllOf = removeSemanticallyEmptyRefs(ref.Value.AllOf)

	isSelfRef := false
	for _, subSchema := range ref.Value.AllOf {
		if _, exists := passed[subSchema]; exists {
			isSelfRef = true
			continue
		}
		passed[subSchema] = something{}
		out = deepMerge(out, subSchema, passed)
	}

	// remove the reference field on AllOf values, these should be
	// flattened into a single (potentially inlined) struct.
	// We make exceptions for
	// 		* types are contain a self-reference, these can not be flattened
	// 		* types are a single allOf, these usually indicate merging in nullable properties, etc
	// 		  it is not a _real_ allOf, rather we are just overriding fields on a single type.
	// 		  this case should already be handled above, so this is really checking if AllOf > 0,
	// 		  but checking >1 is more semantically correct
	//
	if !isSelfRef && len(ref.Value.AllOf) > 1 {
		out.Ref = ""
	}

	return out
}

// removeSemanticallyEmptyRefs removes SchemaRefs from an allOf list  that will not produce any
// meaningful changes to the generated type. Specifically this will remove things like
//
//   allOf:
//      - description: "this will only set the description value"
func removeSemanticallyEmptyRefs(allOf []*openapi3.SchemaRef) []*openapi3.SchemaRef {
	if allOf == nil {
		return allOf
	}

	refCount := 0
	for _, subRef := range allOf {
		if subRef.Ref != "" {
			refCount++
		}
	}

	// nothing we can do here, there are multiple references, so any
	// difficult mixins can be handled by the deepMerge
	if refCount != 1 {
		return allOf
	}

	reduced := []*openapi3.SchemaRef{}
	for _, subRef := range allOf {
		if isNonEmptySchemaRef(subRef) {
			reduced = append(reduced, subRef)
		}
	}

	return reduced
}

// isNonEmptySchemaRef tests if the given SchemaRef has a non-trivial
// definition. This means that at least one of its properties has an impact
// on the resulting type or validation of the generated model. Changes
// to the description, title, etc, are considered "trivial" changes that
// we can choose to ignore.
func isNonEmptySchemaRef(ref *openapi3.SchemaRef) bool {
	if ref == nil {
		return false
	}

	// it refers to another schema
	if ref.Ref != "" {
		return true
	}

	// has empty ref _and_ empty value
	if ref.Value == nil {
		return false
	}

	val := ref.Value
	switch {
	case
		// schema properties
		val.Nullable, val.ReadOnly, val.WriteOnly,
		val.Format != "",
		val.Default != nil,
		val.Type != "",
		// string related
		val.MinLength != 0,
		val.MaxLength != nil,
		val.Pattern != "",
		// number related
		val.Min != nil,
		val.Max != nil,
		val.MultipleOf != nil,
		// array related
		val.Items != nil,
		val.MinItems != 0,
		val.MaxItems != nil,
		// object related
		val.MinProps != 0,
		val.MaxProps != nil,
		// array values properties
		len(val.Required) != 0,
		len(val.Properties) != 0, len(val.Enum) != 0,
		len(val.AllOf) != 0, len(val.AnyOf) != 0:
		return true
	default:
		return false
	}
}

// deepMerge merges `right` into `left` schema recursively resolving types (e.g. `allOf`).
// `passed` map will be populated with all the visited types during the resolution process, so we can
// avoid the infinite loop.
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
	if right.Value.Title != "" && out.Value.Title == "" {
		out.Value.Title = right.Value.Title
	}
	if right.Value.Description != "" && out.Value.Description == "" {
		out.Value.Description = right.Value.Description
	}
	if right.Value.ExternalDocs != nil && out.Value.ExternalDocs == nil {
		out.Value.ExternalDocs = right.Value.ExternalDocs
	}
	if right.Value.Example != nil && out.Value.Example == nil {
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

// structPropsFromRef creates property descriptors for a Go struct from a schema
func structPropsFromRef(ref *openapi3.SchemaRef) (specs []PropSpec, imports []string, err error) {
	if ref == nil || ref.Value == nil {
		return nil, nil, nil
	}

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

// fillValidationRelatedProperties sets validation rules for the property descriptor out of format and
// bounds in the spec.
// Returns a list of unique imports in order to implement the validation.
func fillValidationRelatedProperties(ref *openapi3.SchemaRef, spec *PropSpec) (imports []string) {
	importsMap := make(map[string]something)

	if validatedTypesRegExp.MatchString(spec.GoType) {
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

	if ref.Value.Format != "" {
		spec.HasFormat = true
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
		importsMap["time"] = something{}

	case "byte":
		spec.IsRequiredInValidation = !spec.IsNullable && spec.IsRequired
		spec.NeedsValidation = true
		spec.IsBase64 = true
		importsMap["github.com/go-ozzo/ozzo-validation/v4/is"] = something{}

	case "email":
		spec.IsRequiredInValidation = !spec.IsNullable && spec.IsRequired
		spec.NeedsValidation = true
		spec.IsEmail = true
		importsMap["github.com/go-ozzo/ozzo-validation/v4/is"] = something{}

	case "uuid":
		spec.IsRequiredInValidation = !spec.IsNullable && spec.IsRequired
		spec.NeedsValidation = true
		spec.IsUUID = true
		importsMap["github.com/go-ozzo/ozzo-validation/v4/is"] = something{}

	case "url":
		spec.IsRequiredInValidation = !spec.IsNullable && spec.IsRequired
		spec.NeedsValidation = true
		spec.IsURL = true
		importsMap["github.com/go-ozzo/ozzo-validation/v4/is"] = something{}

	case "uri":
		spec.IsRequiredInValidation = !spec.IsNullable && spec.IsRequired
		spec.NeedsValidation = true
		spec.IsURI = true
		importsMap["github.com/go-ozzo/ozzo-validation/v4/is"] = something{}

	case "request-uri":
		spec.IsRequiredInValidation = !spec.IsNullable && spec.IsRequired
		spec.NeedsValidation = true
		spec.IsRequestURI = true
		importsMap["github.com/go-ozzo/ozzo-validation/v4/is"] = something{}

	case "hostname":
		spec.IsRequiredInValidation = !spec.IsNullable && spec.IsRequired
		spec.NeedsValidation = true
		spec.IsHostname = true
		importsMap["github.com/go-ozzo/ozzo-validation/v4/is"] = something{}

	case "ipv4":
		spec.IsRequiredInValidation = !spec.IsNullable && spec.IsRequired
		spec.NeedsValidation = true
		spec.IsIPv4 = true
		importsMap["github.com/go-ozzo/ozzo-validation/v4/is"] = something{}

	case "ipv6":
		spec.IsRequiredInValidation = !spec.IsNullable && spec.IsRequired
		spec.NeedsValidation = true
		spec.IsIPv6 = true
		importsMap["github.com/go-ozzo/ozzo-validation/v4/is"] = something{}

	case "ip":
		spec.IsRequiredInValidation = !spec.IsNullable && spec.IsRequired
		spec.NeedsValidation = true
		spec.IsIP = true
		importsMap["github.com/go-ozzo/ozzo-validation/v4/is"] = something{}

	case "", "int32", "int64", "float", "double":
		break // do nothing

	}

	imports = make([]string, 0, len(importsMap))
	for importStr := range importsMap {
		imports = append(imports, importStr)
	}

	return imports
}

// enumPropsFromRef generates a list of enum property/item descriptors.
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

// checkIfRequired returns true if `name` is in the `list`
func checkIfRequired(name string, list []string) bool {
	for _, n := range list {
		if n == name {
			return true
		}
	}
	return false
}

// uniqueStrings filters out duplicates from `input` and forms a list of unique values `output`
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
