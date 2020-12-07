// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// FieldSpec is an object.
type FieldSpec struct {
	// Description: A description of the purpose of this field
	Description string `json:"description"`
	// DisplayName: A nice human-readable name for that field
	DisplayName string `json:"displayName,omitempty"`
	// Filterable: Should this field show up as a filter in the catalog
	Filterable bool `json:"filterable"`
	// HideIfEmpty: Should this field be hidden when the value is empty
	HideIfEmpty bool `json:"hideIfEmpty"`
	// Key: lookup key for the field on the specific response types in dot notation.
	// Examples:
	//   `name` -> name is found toplevel of the instance object, i.e. `obj.name`
	//   `properties.category` -> category is found in the properties subobject, i.e. `obj.properties.category`
	Key string `json:"key"`
	// Kind: The source type of the field.
	//
	// - System fields are non-editable default fields.  These can not be modified for deleted.
	//
	// - Configurable fields are editable default fields. These can be modified but not deleted.
	//
	// - Custom are user defined fields. These are fully defined by the user.
	Kind FieldKind `json:"kind"`
	// MaxValue:
	MaxValue *float32 `json:"maxValue,omitempty"`
	// MinValue:
	MinValue *float32 `json:"minValue,omitempty"`
	// Name: A nice computer-readable name for that field usable in postgres
	Name string `json:"name"`
	// Searchable: Should this field be used in the catalog fuzzy search
	Searchable bool `json:"searchable"`
	// Type: The data type of the field:
	//
	// - `text` is an arbitrary string
	//
	// - `multitext` is an arbitrary array of strings
	//
	// - `select` a value which can be selected from a list
	//
	// - `multiselect` is a `select` with multiple selected values, represented as an array
	//
	// - `boolean` is `true` or `false` value
	//
	// - `date` is a date string in `yyyy-mm-dd` format, e.g. `2019-08-21`
	//
	// - `datetime` is an RFC3339 datetime string, e.g. `2019-08-21T11:54:30.369917Z`
	//
	// - `number` is a numeric value, described by RFC8259, section 6.
	//
	// - `link` is a link object with a displayed text and URL, e.g. `{"displayText": "text", "url": "http://example.com"}`. Only absolute URLs with HTTP/HTTPS protocol are allowed.
	Type FieldType `json:"type"`
	// ValidValues:
	ValidValues []string `json:"validValues,omitempty"`
}

// GetDescription returns the Description property
func (m FieldSpec) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m FieldSpec) SetDescription(val string) {
	m.Description = val
}

// GetDisplayName returns the DisplayName property
func (m FieldSpec) GetDisplayName() string {
	return m.DisplayName
}

// SetDisplayName sets the DisplayName property
func (m FieldSpec) SetDisplayName(val string) {
	m.DisplayName = val
}

// GetFilterable returns the Filterable property
func (m FieldSpec) GetFilterable() bool {
	return m.Filterable
}

// SetFilterable sets the Filterable property
func (m FieldSpec) SetFilterable(val bool) {
	m.Filterable = val
}

// GetHideIfEmpty returns the HideIfEmpty property
func (m FieldSpec) GetHideIfEmpty() bool {
	return m.HideIfEmpty
}

// SetHideIfEmpty sets the HideIfEmpty property
func (m FieldSpec) SetHideIfEmpty(val bool) {
	m.HideIfEmpty = val
}

// GetKey returns the Key property
func (m FieldSpec) GetKey() string {
	return m.Key
}

// SetKey sets the Key property
func (m FieldSpec) SetKey(val string) {
	m.Key = val
}

// GetKind returns the Kind property
func (m FieldSpec) GetKind() FieldKind {
	return m.Kind
}

// SetKind sets the Kind property
func (m FieldSpec) SetKind(val FieldKind) {
	m.Kind = val
}

// GetMaxValue returns the MaxValue property
func (m FieldSpec) GetMaxValue() *float32 {
	return m.MaxValue
}

// SetMaxValue sets the MaxValue property
func (m FieldSpec) SetMaxValue(val *float32) {
	m.MaxValue = val
}

// GetMinValue returns the MinValue property
func (m FieldSpec) GetMinValue() *float32 {
	return m.MinValue
}

// SetMinValue sets the MinValue property
func (m FieldSpec) SetMinValue(val *float32) {
	m.MinValue = val
}

// GetName returns the Name property
func (m FieldSpec) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m FieldSpec) SetName(val string) {
	m.Name = val
}

// GetSearchable returns the Searchable property
func (m FieldSpec) GetSearchable() bool {
	return m.Searchable
}

// SetSearchable sets the Searchable property
func (m FieldSpec) SetSearchable(val bool) {
	m.Searchable = val
}

// GetType returns the Type property
func (m FieldSpec) GetType() FieldType {
	return m.Type
}

// SetType sets the Type property
func (m FieldSpec) SetType(val FieldType) {
	m.Type = val
}

// GetValidValues returns the ValidValues property
func (m FieldSpec) GetValidValues() []string {
	return m.ValidValues
}

// SetValidValues sets the ValidValues property
func (m FieldSpec) SetValidValues(val []string) {
	m.ValidValues = val
}
