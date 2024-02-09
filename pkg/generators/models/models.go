package models

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"reflect"
	"regexp"
	"sort"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/pkg/errors"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"

	tpl "github.com/contiamo/openapi-generator-go/v2/pkg/generators/templates"
)

// ExtensionPatternError is an extension property that, if set, allows the API author
// to control the error message for pattern validation failures.
const ExtensionPatternError = "x-pattern-error"

var (
	// validatedTypesRegExp helps to match a type to a type that we can generate validation for.
	// Exported types may have a Validate() function
	// Exported types with a leading pointer may have a Validate() function
	// Slices may contain data that is validatable (they are also never pointers)
	// Maps may contain data that is validatable (they are also never pointers )
	validatedTypesRegExp = regexp.MustCompile(`^(\*?[A-Z])|(\[])|(map\[)`)
)

// passedSchemas is a map of pointers to a schema.
// Used for marking already visited types when recursively resolve `allOf` definitions and merging types.
type passedSchemas map[*openapi3.SchemaRef]bool

// TemplateKind is a kind of template to render
type TemplateKind string

const (
	// Struct is a regular Go struct
	Struct TemplateKind = "struct"
	// Enum is a Go enum definition
	Enum TemplateKind = "enum"
	// Value is a value type
	Value TemplateKind = "value"
	// OneOf is a oneof value
	OneOf TemplateKind = "oneof"
)

// Model is a template model for rendering Go code for a given API schema
type Model struct {
	// Validation stuff
	ValidationSpec

	// Name is a name of the generated type that follows the `type` keyword in the definition
	// For example, `type Name GoType`.
	Name string
	// Description is a description that will become a comment on the generated type
	Description string
	// Imports is a list of imports that will be present in the Go file. Key is the package and value is the alias
	Imports map[string]string
	// TemplateKind is a kind of generated model (e.g. `struct` or `enum`)
	TemplateKind TemplateKind
	// Properties is a list of type's property descriptors
	Properties []PropSpec
	// ConvertSpecs contains a list of convert functions for this model
	ConvertSpecs []ConvertSpec
	// Discriminator contains the optional oneOf discriminator
	Discriminator Discriminator
	// GoType is a string that represents the Go type that follows the model type name.
	// For example, `type Name GoType`.
	GoType string
	// HasDefault is true when the model has a default value
	HasDefault bool
	// DefaultValue is the default value of the model
	DefaultValue string
	// SpecTitle is the spec title used in the auto-generated header
	SpecTitle string
	// SpecVersion is the spec version used in the auto-generated header
	SpecVersion string
	// PackageName is the name of the package used in the Go code
	PackageName string
	// AdditionalPropertiesGoType is the optional type of additional properties
	// that exist _in addition_ to `Properties`
	AdditionalProperties *PropSpec
	// IsMap is true when the model is a map
	IsMap bool
}

// ConvertSpec holds all info to build one As{Type}() function
type ConvertSpec struct {
	// TargetGoType is the target type of the conversion
	TargetGoType string
	// IsPtr is true when the target type is a pointer
	IsPtr bool
	// IsMap is true when the target type is a map
	IsMap        bool
	HasDefault   bool
	DefaultValue string
	IsStruct     bool
	IsRef        bool
}

type ValidationSpec struct {
	// Validation stuff
	HasPattern                 bool
	Pattern                    string
	PatternErrorMsg            string
	NeedsValidation            bool
	HasMin, HasMax             bool
	Min, Max                   float64
	HasMinLength, HasMaxLength bool
	MinLength, MaxLength       uint64
	HasMinItems, HasMaxItems   bool
	MinItems, MaxItems         uint64
	HasMinProps, HasMaxProps   bool
	MinProps, MaxProps         uint64
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
	IsEnumWithNil              bool
	IsEnumWithZero             bool

	// Derived validations that should be added to the property
	DerivedValidations []string
}

// PropSpec is a Go property descriptor
type PropSpec struct {
	// Validation stuff
	ValidationSpec

	// Name is a property name in structs, variable name in enums, etc
	Name string
	// PropertyName is the original name of the property
	PropertyName string
	// Description used in the comment of the property
	Description string
	// GoType used for this property (e.g. `string`, `int`, etc)
	GoType string
	// HasDefault is true when the property has a default value
	HasDefault bool
	// DefaultValue is the default value of the property
	DefaultValue string
	// JSONTags is a string of JSON tags used for marshaling (e.g. `json:omitempty`)
	JSONTags string
	// Value is a value used for a enum item
	Value string
	// IsRequired is true when the property is a required value and should not be omitted on marshaling
	IsRequired bool
	// IsPtr is true when the property is a pointer. This is derived from the nullable property in the schema
	IsPtr bool
	// IsEnum is true when the property is a enum item
	IsEnum bool
	// IsOneOf is true when the property is a `oneof` schema
	IsOneOf bool
	// IsArray is true when the property is an array
	IsArray bool
	// IsMap is true when the property is a map
	IsMap bool
	// IsStruct is true when the property is a struct
	IsStruct bool
	// IsString is true when the property is a string
	IsString bool
	// IsNumber is true when the property is a number
	IsNumber bool
	// IsInteger is true when the property is an integer
	IsInteger bool
	// IsRef is true when the property is a Ref to a schema
	IsRef bool
}

type Discriminator struct {
	// Field points to the property that specifies the data type name
	Field string
	// Map contains the mapping from Field values to target schemas
	Map map[string]string
}

// NewModelFromRef creates a model out of a schema
func NewModelFromRef(ref *openapi3.SchemaRef) (model *Model, err error) {
	model = &Model{
		Description: ref.Value.Description,
	}

	switch {
	case len(ref.Value.Enum) > 0:
		model.TemplateKind = Enum
		model.GoType = goTypeFromSpec(ref)
		model.Properties = enumPropsFromRef(ref, model)
		if ref.Value.Default != nil {
			err = isDefValueValid(ref.Value)
			if err != nil {
				return nil, err
			}
			model.HasDefault = true
			if defVal, present := getDefaultValue(ref.Value); present {
				model.DefaultValue = defVal
			}
		}

	case ref.Value.Type == "object" || len(ref.Value.Properties) > 0:
		model.TemplateKind = Struct
		model.Properties, model.Imports, err = structPropsFromRef(ref)
		if ref.Value.Default != nil {
			err = isDefValueValid(ref.Value)
			if err != nil {
				return nil, err
			}
			model.HasDefault = true
			if defVal, present := getDefaultValue(ref.Value); present {
				model.DefaultValue = defVal
			}
		}

		if ref.Value.AdditionalProperties != nil {
			model.AdditionalProperties, err = newPropertySpecFromRef("AdditionalProperties", ref.Value.AdditionalProperties, true, model.Imports)
			if err != nil {
				return nil, err
			}
		} else if ref.Value.AdditionalPropertiesAllowed != nil && *ref.Value.AdditionalPropertiesAllowed {
			model.AdditionalProperties = &PropSpec{
				Name:   "AdditionalProperties",
				GoType: "interface{}",
			}
		}

		// This is following what was resolved in: https://github.com/contiamo/openapi-generator-go/issues/139#issuecomment-1958083054
		if len(ref.Value.Properties) == 0 {
			model.IsMap = true
			model.GoType = "map[string]interface{}"
			if ref.Value.AdditionalProperties != nil {
				model.GoType = "map[string]"
				if model.AdditionalProperties.IsPtr {
					model.GoType += "*"
				}
				model.GoType += model.AdditionalProperties.GoType
			} else if ref.Value.AdditionalPropertiesAllowed != nil && !*ref.Value.AdditionalPropertiesAllowed {
				model.GoType = "struct{}"
				model.IsMap = false
			}
		}

	case len(ref.Value.OneOf) > 0:
		model.TemplateKind = OneOf
		err = model.configureOneOf(ref)
		if err != nil {
			return nil, err
		}
	default:
		model.TemplateKind = Value
		model.GoType = goTypeFromSpec(ref)
		if ref.Value.Default != nil {
			err = isDefValueValid(ref.Value)
			if err != nil {
				return nil, err
			}
			model.HasDefault = true
			if defVal, present := getDefaultValue(ref.Value); present {
				model.DefaultValue = defVal
			}
		}
		model.Imports = make(map[string]string)
		err = fillValidationSpec(ref, &model.ValidationSpec, model.GoType, model.Imports)
	}
	if err != nil {
		return nil, err
	}

	return model, nil
}

func newPropertySpecFromRef(name string, prop *openapi3.SchemaRef, isRequired bool, imports map[string]string) (spec *PropSpec, err error) {
	spec = &PropSpec{
		Name:         tpl.ToPascalCase(name),
		PropertyName: name,
		GoType:       goTypeFromSpec(prop),
		Description:  prop.Value.Description,
		IsRequired:   isRequired, // additional properties are always required for them to be validated properly
		IsEnum:       len(prop.Value.Enum) > 0,
		IsArray:      prop.Value.Type == "array",
		IsMap:        (prop.Value.Type == "" || prop.Value.Type == "object") && len(prop.Value.Properties) == 0,
		IsStruct:     (prop.Value.Type == "" || prop.Value.Type == "object") && len(prop.Value.Properties) > 0,
		IsString:     prop.Value.Type == "string",
		IsNumber:     prop.Value.Type == "number",
		IsInteger:    prop.Value.Type == "integer",
		IsRef:        prop.Ref != "",
		ValidationSpec: ValidationSpec{
			IsEnumWithNil: len(prop.Value.Enum) > 0 && slices.ContainsFunc(prop.Value.Enum, func(i any) bool {
				// Reflect considers the enum value nil to be invalid because it doesn't have a type for it
				return !reflect.ValueOf(i).IsValid()
			}),
			IsEnumWithZero: len(prop.Value.Enum) > 0 && slices.ContainsFunc(prop.Value.Enum, func(i any) bool {
				val := reflect.ValueOf(i)
				return val.IsValid() && val.IsZero()
			}),
		},
		IsOneOf: prop.Value.OneOf != nil && len(prop.Value.OneOf) > 0,
	}
	if prop.Value.Default != nil {
		spec.HasDefault = true
		if prop.Ref == "" {
			err = isDefValueValid(prop.Value)
			if err != nil {
				return nil, fmt.Errorf("invalid default value for additionalProperty: %w", err)
			}
			if defVal, present := getDefaultValue(prop.Value); present {
				spec.DefaultValue = defVal
			}
		}
	}

	spec.IsPtr = prop.Value.Nullable && !(spec.IsMap || spec.IsArray) && (!spec.IsEnum || spec.IsEnumWithNil)

	spec.IsString = spec.IsString && !spec.IsEnum

	if spec.GoType == "time.Time" || spec.GoType == "*time.Time" {
		imports["time"] = ""
	}

	omit := ""
	if !spec.IsRequired {
		omit += ",omitempty"
	}

	jsonTags := fmt.Sprintf(" `json:\"%s%s\"", name, omit)
	jsonTags += fmt.Sprintf(" mapstructure:\"%s%s\"", name, omit)
	jsonTags += "`"

	spec.JSONTags = jsonTags

	err = fillValidationSpec(prop, &spec.ValidationSpec, spec.GoType, imports)
	if err != nil {
		return nil, err
	}
	resolveDerivedValidation(spec, imports)

	return spec, nil
}

// NewModelFromParameters returns a model built from operation parameters
func NewModelFromParameters(params openapi3.Parameters) (model *Model, err error) {
	model = &Model{
		TemplateKind: Struct,
		Imports:      make(map[string]string),
	}

	// convert to SchemaRef to reuse struct props logic
	ref := &openapi3.SchemaRef{
		Value: &openapi3.Schema{
			Type:       "object",
			Required:   make([]string, 0),
			Properties: make(map[string]*openapi3.SchemaRef),
		},
	}
	for _, param := range params {
		if param.Value.Schema != nil {
			schema := param.Value.Schema.Value
			if param.Value.Description != "" {
				// use a copy to modify the description without affecting the source schema
				schemaCopy := *param.Value.Schema.Value
				schemaCopy.Description = param.Value.Description
				schema = &schemaCopy
			}
			ref.Value.Properties[param.Value.Name] = &openapi3.SchemaRef{
				Ref:   param.Value.Schema.Ref,
				Value: schema,
			}
			if param.Value.Required {
				ref.Value.Required = append(ref.Value.Required, param.Value.Name)
			}
		}
	}

	model.Properties, model.Imports, err = structPropsFromRef(ref)
	return model, err
}

func (m *Model) configureOneOf(ref *openapi3.SchemaRef) error {
	for _, oneOf := range ref.Value.OneOf {
		spec := ConvertSpec{
			TargetGoType: goTypeFromSpec(oneOf),
			IsPtr:        oneOf.Value.Nullable && oneOf.Value.Type != "" && oneOf.Value.Type != "object" && oneOf.Value.Type != "array",
			IsMap:        (oneOf.Value.Type == "" || oneOf.Value.Type == "object") && len(oneOf.Value.Properties) == 0,
			IsStruct:     (oneOf.Value.Type == "" || oneOf.Value.Type == "object") && len(oneOf.Value.Properties) > 0,
			IsRef:        oneOf.Ref != "",
			HasDefault:   oneOf.Value.Default != nil,
		}
		m.ConvertSpecs = append(m.ConvertSpecs, spec)
		if oneOf.Value.Default != nil {
			err := isDefValueValid(oneOf.Value)
			if err != nil {
				return err
			}
			spec.HasDefault = true
			if defVal, present := getDefaultValue(oneOf.Value); present {
				spec.DefaultValue = defVal
			}
		}
	}

	if ref.Value.Discriminator != nil {
		m.Discriminator.Field = ref.Value.Discriminator.PropertyName
		m.Discriminator.Map = map[string]string{}

		for value, ref := range ref.Value.Discriminator.Mapping {
			m.Discriminator.Map[value] = typeNameFromRef(ref)
		}

		// the default mapping is to use the name of the model as the field value.
		if len(ref.Value.Discriminator.Mapping) == 0 {
			for _, schema := range ref.Value.OneOf {
				name := typeNameFromRef(schema.Ref)
				if name == "" {
					// skip mappings for inlined types, ie non-references
					continue
				}
				m.Discriminator.Map[name] = name
			}
		}
	}
	return nil
}

// Render renders the model to a Go file
func (m *Model) Render(ctx context.Context, writer io.Writer) error {
	var name string

	switch m.TemplateKind {
	case Struct:
		name = "model.gotpl"
	case Enum:
		name = "enum.gotpl"
	case Value:
		name = "value.gotpl"
	case OneOf:
		name = "oneof.gotpl"
	}

	err := modelTpls.ExecuteTemplate(writer, name, m)
	if err != nil {
		return errors.Wrap(err, "failed to render the model")
	}

	return nil
}

// resolveAllOf resolves the list of `allOf` definitions in the schema to a complete type merging
// all the mentioned types.
// `passed` can be nil, it's used recursively in order to avoid infinite loops
func resolveAllOf(ref *openapi3.SchemaRef, passed passedSchemas) (out *openapi3.SchemaRef) {
	defer func() {
		// clean up the allOf field after resolving it
		ref.Value.AllOf = nil
	}()

	if ref == nil {
		ref = &openapi3.SchemaRef{}
	}
	if ref.Value == nil {
		ref.Value = &openapi3.Schema{}
	}

	out = ref
	if len(ref.Value.AllOf) == 0 {
		return out
	}

	if passed == nil {
		passed = make(passedSchemas)
	}

	// very special case where the allOf references
	// another named schema but is not actually merging
	// several schemas, this is usually just overriding
	// a value like nullable or the description
	if len(ref.Value.AllOf) == 1 {
		pointer := ref.Value.AllOf[0]
		passed[pointer] = true
		out = deepMerge(out, pointer, passed)
		return out
	}

	// this is used for the special edge cases like this
	//   allOf:
	//      - description: "this will only set the description value"
	//      - $ref: '#/components/schemas/ColumnTypeMetadata'
	//
	// without this method, the allOf will generate an inline type, but the resulting code
	// is much better and easier to use if we drop the first allOf item and just treat it like
	// the special case of a single `ref`, which will used the named `ColumnTypeMetadata` type.
	ref.Value.AllOf = removeSemanticallyEmptyRefs(ref.Value.AllOf)

	isSelfRef := false
	for _, subSchema := range ref.Value.AllOf {
		if passed[subSchema] {
			isSelfRef = true
			continue
		}
		passed[subSchema] = true
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
//	allOf:
//	   - description: "this will only set the description value"
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

	right = resolveAllOf(right, passed)
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

	// it is possible during oneOf or allOf statements to have multiple references to the
	// same schema. We need to deduplicate them to avoid infinite loops.
	out.Value.OneOf = dedup(append(out.Value.OneOf, right.Value.OneOf...))
	out.Value.AllOf = dedup(append(out.Value.AllOf, right.Value.AllOf...))
	out.Value.AnyOf = dedup(append(out.Value.AnyOf, right.Value.AnyOf...))
	out.Value.Enum = dedup(append(out.Value.Enum, right.Value.Enum...))
	out.Value.Required = dedup(append(out.Value.Required, right.Value.Required...))

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

func dedup[T any](input []T) []T {
	seen := make(map[string]T)

	for _, x := range input {
		key := fmt.Sprintf("%v", x)
		seen[key] = x
	}
	return maps.Values(seen)
}

// structPropsFromRef creates property descriptors for a Go struct from a schema
func structPropsFromRef(ref *openapi3.SchemaRef) (specs []PropSpec, imports map[string]string, err error) {
	if ref == nil || ref.Value == nil {
		return nil, nil, nil
	}

	imports = make(map[string]string)
	for _, name := range sortedKeys(ref.Value.Properties) {
		prop := ref.Value.Properties[name]

		// special case for singleton allOf
		// generally these are used to add
		// a new description or override the
		// nullable flag
		if len(prop.Value.AllOf) == 1 {
			pointer := prop.Value.AllOf[0]
			prop.Ref = pointer.Ref

			// a sub-special case when the referenced type
			// is an array. Because we don't create top-level
			// models for arrays, we actually need to follow
			// this reference.
			resolved := resolveAllOf(prop, nil)
			if resolved.Value.Type == "array" {
				prop = resolved
			}
		}

		if prop.Ref == "" {
			prop = resolveAllOf(prop, nil)
		}

		isRequired := checkIfRequired(name, ref.Value.Required)
		var spec *PropSpec
		spec, err = newPropertySpecFromRef(name, prop, isRequired, imports)
		if err != nil {
			return nil, nil, err
		}

		// If an object property is not required but also not nullable, it must be a pointer to allow it to either be
		// fully specified or not present at all.
		spec.IsPtr = spec.IsPtr || !isRequired && !prop.Value.Nullable && spec.IsStruct

		specs = append(specs, *spec)
	}

	return specs, imports, err
}

// resolveDerivedValidation resolves the extra validation rules that are derived from the required, nullable and other
// properties that affect the validation rules. Example: OZZO min length validation for strings consider the empty string
// valid, so we need to add the OZZO Required validation to catch that.
func resolveDerivedValidation(spec *PropSpec, imports map[string]string) {
	for validation, specs := range DerivedRulesByValidationType {
		for _, rule := range specs {
			if rule.IsStruct == spec.IsStruct &&
				rule.IsEnum == spec.IsEnum &&
				rule.IsString == spec.IsString &&
				rule.IsMap == spec.IsMap &&
				rule.IsArray == spec.IsArray &&
				rule.IsNumber == spec.IsNumber &&
				rule.IsInteger == spec.IsInteger &&
				rule.IsPtr == spec.IsPtr &&
				(rule.IsEnumWithNull == nil || *rule.IsEnumWithNull == spec.IsEnumWithNil) &&
				(rule.IsEnumWithZero == nil || *rule.IsEnumWithZero == spec.IsEnumWithZero) &&
				(rule.IsRequired == nil || *rule.IsRequired == spec.IsRequired) &&
				(rule.HasFormat == nil || *rule.HasFormat == spec.HasFormat) &&
				(rule.HasPattern == nil || *rule.HasPattern == spec.HasPattern) &&
				(rule.HasMin == nil || *rule.HasMin == spec.HasMin) &&
				(rule.IsMinGreaterThanZero == nil || *rule.IsMinGreaterThanZero == (spec.Min > 0)) &&
				(rule.HasMax == nil || *rule.HasMax == spec.HasMax) &&
				(rule.IsMaxLessThanZero == nil || *rule.IsMaxLessThanZero == (spec.Max < 0)) &&
				(rule.HasMinLength == nil || *rule.HasMinLength == spec.HasMinLength) &&
				(rule.HasMaxLength == nil || *rule.HasMaxLength == spec.HasMaxLength) &&
				(rule.HasMinItems == nil || *rule.HasMinItems == spec.HasMinItems) &&
				(rule.HasMaxItems == nil || *rule.HasMaxItems == spec.HasMaxItems) &&
				(rule.HasMinProps == nil || *rule.HasMinProps == spec.HasMinProps) &&
				(rule.HasMaxProps == nil || *rule.HasMaxProps == spec.HasMaxProps) {

				/*  // Enable for easier debugging of the rules
				rowJSON, _ := json.Marshal(rule)
				fmt.Printf("Adding derived validation %s to %s, matching rule: %s\n", validation, spec.Name, string(rowJSON))
				*/
				imports["github.com/go-ozzo/ozzo-validation/v4"] = "validation"
				spec.DerivedValidations = append(spec.DerivedValidations, "validation."+validation)
				// We only add one validation type per property, so we break here
				break
			}
		}
	}
}

func fillValidationSpec(ref *openapi3.SchemaRef, spec *ValidationSpec, goType string, imports map[string]string) (err error) {
	if validatedTypesRegExp.MatchString(goType) {
		// enable recursive validation
		spec.NeedsValidation = true
	}

	if ref.Value.Min != nil {
		spec.NeedsValidation = true
		spec.HasMin = true
		spec.Min = *ref.Value.Min
	}

	if ref.Value.Max != nil {
		spec.NeedsValidation = true
		spec.HasMax = true
		spec.Max = *ref.Value.Max
	}

	if ref.Value.MinLength > 0 {
		spec.NeedsValidation = true
		spec.HasMinLength = true
		spec.MinLength = ref.Value.MinLength
	}

	if ref.Value.MaxLength != nil {
		spec.NeedsValidation = true
		spec.HasMaxLength = true
		spec.MaxLength = *ref.Value.MaxLength
	}

	if ref.Value.MinItems > 0 {
		spec.NeedsValidation = true
		spec.HasMinItems = true
		spec.MinItems = ref.Value.MinItems
	}

	if ref.Value.MaxItems != nil {
		spec.NeedsValidation = true
		spec.HasMaxItems = true
		spec.MaxItems = *ref.Value.MaxItems
	}

	if ref.Value.MinProps > 0 {
		spec.NeedsValidation = true
		spec.HasMinProps = true
		spec.MinProps = ref.Value.MinProps
	}

	if ref.Value.MaxProps != nil {
		spec.NeedsValidation = true
		spec.HasMaxProps = true
		spec.MaxProps = *ref.Value.MaxProps
	}

	if ref.Value.ExclusiveMin {
		spec.Min += math.SmallestNonzeroFloat64
	}

	if ref.Value.ExclusiveMax {
		spec.Min -= math.SmallestNonzeroFloat64
	}

	if ref.Value.Pattern != "" {
		spec.NeedsValidation = true
		spec.HasPattern = true
		spec.Pattern = ref.Value.Pattern
		imports["regexp"] = ""
		msg, ok := ref.Value.ExtensionProps.Extensions[ExtensionPatternError]
		if ok {
			rawMsg, ok := msg.(json.RawMessage)
			if !ok {
				return fmt.Errorf("invalid %s value, expected json msg, got %T", ExtensionPatternError, msg)
			}
			err = json.Unmarshal(rawMsg, &spec.PatternErrorMsg)
			if err != nil {
				return fmt.Errorf("invalid %s value, must be a string", ExtensionPatternError)
			}
		}
	}

	if ref.Value.Format != "" {
		spec.HasFormat = true
	}
	switch ref.Value.Format {
	case "date":
		spec.NeedsValidation = true
		spec.IsDate = true

	case "date-time":
		spec.NeedsValidation = true
		spec.IsDateTime = true
		if ref.Ref == "" {
			imports["time"] = ""
		}

	case "byte":
		spec.NeedsValidation = true
		spec.IsBase64 = true
		if ref.Ref == "" {
			imports["github.com/go-ozzo/ozzo-validation/v4/is"] = ""
		}

	case "email":
		spec.NeedsValidation = true
		spec.IsEmail = true
		if ref.Ref == "" {
			imports["github.com/go-ozzo/ozzo-validation/v4/is"] = ""
		}

	case "uuid":
		spec.NeedsValidation = true
		spec.IsUUID = true
		if ref.Ref == "" {
			imports["github.com/go-ozzo/ozzo-validation/v4/is"] = ""
		}

	case "url":
		spec.NeedsValidation = true
		spec.IsURL = true
		if ref.Ref == "" {
			imports["github.com/go-ozzo/ozzo-validation/v4/is"] = ""
		}

	case "uri":
		spec.NeedsValidation = true
		spec.IsURI = true
		if ref.Ref == "" {
			imports["github.com/go-ozzo/ozzo-validation/v4/is"] = ""
		}

	case "request-uri":
		spec.NeedsValidation = true
		spec.IsRequestURI = true
		if ref.Ref == "" {
			imports["github.com/go-ozzo/ozzo-validation/v4/is"] = ""
		}

	case "hostname":
		spec.NeedsValidation = true
		spec.IsHostname = true
		if ref.Ref == "" {
			imports["github.com/go-ozzo/ozzo-validation/v4/is"] = ""
		}

	case "ipv4":
		spec.NeedsValidation = true
		spec.IsIPv4 = true
		if ref.Ref == "" {
			imports["github.com/go-ozzo/ozzo-validation/v4/is"] = ""
		}

	case "ipv6":
		spec.NeedsValidation = true
		spec.IsIPv6 = true
		if ref.Ref == "" {
			imports["github.com/go-ozzo/ozzo-validation/v4/is"] = ""
		}

	case "ip":
		spec.NeedsValidation = true
		spec.IsIP = true
		if ref.Ref == "" {
			imports["github.com/go-ozzo/ozzo-validation/v4/is"] = ""
		}

	case "", "int32", "int64", "float", "double":
		break // do nothing
	}

	if spec.NeedsValidation {
		imports["github.com/go-ozzo/ozzo-validation/v4"] = "validation"
	}

	return nil
}

// enumPropsFromRef generates a list of enum property/item descriptors.
func enumPropsFromRef(ref *openapi3.SchemaRef, model *Model) (specs []PropSpec) {
	for idx, val := range ref.Value.Enum {
		// Ignore nil values as it is handled by the referencing struct
		if val == nil {
			continue
		}
		valueVarName := tpl.ToPascalCase(fmt.Sprintf("%v", val))
		if valueVarName == "" {
			valueVarName = fmt.Sprintf("Item%d", idx)
		}

		var enumValue string
		switch ref.Value.Type {
		case "integer", "number":
			enumValue = fmt.Sprintf("%v", val)
		default:
			enumValue = fmt.Sprintf(`"%v"`, val)
		}
		specs = append(specs, PropSpec{
			Name:   valueVarName,
			Value:  enumValue,
			GoType: model.Name,
		})
	}

	sort.Slice(specs, func(i, j int) bool {
		return specs[i].Name < specs[j].Name
	})

	return specs
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

func isDefValueValid(schema *openapi3.Schema) error {
	if schema.Default == nil {
		return nil
	}
	rVal := reflect.ValueOf(schema.Default)
	if !rVal.IsValid() {
		return fmt.Errorf("default value '%v' is not valid for type %s", schema.Default, schema.Type)
	}
	switch schema.Type {
	case "number":
		if !rVal.CanFloat() {
			return fmt.Errorf("default value %#v can't be assigned as type %s", rVal, schema.Type)
		}
	case "integer":
		if !rVal.CanConvert(reflect.TypeOf(int64(0))) {
			return fmt.Errorf("default value %#v can't be assigned as type %s", rVal, schema.Type)
		}
	case "boolean":
		if rVal.Kind() != reflect.Bool {
			return fmt.Errorf("default value %#v can't be assigned as type %s", rVal, schema.Type)
		}
	case "string":
		if rVal.Kind() != reflect.String {
			return fmt.Errorf("default value %#v can't be assigned as type %s", rVal, schema.Type)
		}
	case "array":
		if rVal.Kind() != reflect.Slice {
			return fmt.Errorf("default value %#v can't be assigned as type %s", rVal, schema.Type)
		}
		for i := 0; i < rVal.Len(); i++ {
			err := isDefValueValid(&openapi3.Schema{Type: schema.Items.Value.Type, Default: rVal.Index(i).Interface()})
			if err != nil {
				return fmt.Errorf("%w in array", err)
			}
		}
	case "object", "":
		if rVal.Kind() != reflect.Map {
			return fmt.Errorf("default value %#v can't be assigned as type %s", rVal, schema.Type)
		}
	}

	if len(schema.Enum) > 0 {
		// Check if the value is present in the enum
		for _, v := range schema.Enum {
			if v == schema.Default {
				return nil
			}
		}
		return fmt.Errorf("default value '%v' is not present in the enum values", schema.Default)
	}

	return nil
}

// getDefaultValue returns the default value of a schema ready to be used in the template
func getDefaultValue(schema *openapi3.Schema) (defVal string, present bool) {
	if schema.Default == nil {
		return "", false
	}

	if schema.Type == "object" || schema.Type == "" {
		bytes, _ := json.Marshal(schema.Default)
		return string(bytes), true
	}

	if schema.Type == "array" {
		anyItems, _ := schema.Default.([]interface{})
		var typedItems any
		switch goTypeFromSpec(schema.Items) {
		case "string":
			items := make([]string, 0, len(anyItems))
			for _, anyItem := range anyItems {
				// Type was verified in isDefValueValid
				item, _ := anyItem.(string)
				items = append(items, item)
			}
			typedItems = items
		case "int32":
			items := make([]int32, 0, len(anyItems))
			for _, anyItem := range anyItems {
				// Type was verified in isDefValueValid
				item, _ := anyItem.(float64)
				items = append(items, int32(item))
			}
			typedItems = items
		case "int64":
			items := make([]int64, 0, len(anyItems))
			for _, anyItem := range anyItems {
				// Type was verified in isDefValueValid
				item, _ := anyItem.(float64)
				items = append(items, int64(item))
			}
			typedItems = items
		case "float32":
			items := make([]float32, 0, len(anyItems))
			for _, anyItem := range anyItems {
				// Type was verified in isDefValueValid
				item, _ := anyItem.(float64)
				items = append(items, float32(item))
			}
			typedItems = items
		case "float64":
			items := make([]float64, 0, len(anyItems))
			for _, anyItem := range anyItems {
				// Type was verified in isDefValueValid
				item, _ := anyItem.(float64)
				items = append(items, item)
			}
			typedItems = items
		}
		return fmt.Sprintf("%#v", typedItems), true
	}

	return fmt.Sprintf("%#v", schema.Default), true
}
