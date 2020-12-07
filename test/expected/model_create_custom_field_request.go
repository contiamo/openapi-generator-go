// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// CreateCustomFieldRequest is an object.
type CreateCustomFieldRequest struct {
	// Description: A description of the purpose of this field
	Description string `json:"description,omitempty"`
	// DisplayName: A nice human-readable name for that field
	DisplayName string `json:"displayName,omitempty"`
	// Filterable: Should this field show up as a filter in the catalog
	Filterable bool `json:"filterable"`
	// HideIfEmpty: Should this field be hidden when its value is empty
	HideIfEmpty bool `json:"hideIfEmpty"`
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
func (m CreateCustomFieldRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m CreateCustomFieldRequest) SetDescription(val string) {
	m.Description = val
}

// GetDisplayName returns the DisplayName property
func (m CreateCustomFieldRequest) GetDisplayName() string {
	return m.DisplayName
}

// SetDisplayName sets the DisplayName property
func (m CreateCustomFieldRequest) SetDisplayName(val string) {
	m.DisplayName = val
}

// GetFilterable returns the Filterable property
func (m CreateCustomFieldRequest) GetFilterable() bool {
	return m.Filterable
}

// SetFilterable sets the Filterable property
func (m CreateCustomFieldRequest) SetFilterable(val bool) {
	m.Filterable = val
}

// GetHideIfEmpty returns the HideIfEmpty property
func (m CreateCustomFieldRequest) GetHideIfEmpty() bool {
	return m.HideIfEmpty
}

// SetHideIfEmpty sets the HideIfEmpty property
func (m CreateCustomFieldRequest) SetHideIfEmpty(val bool) {
	m.HideIfEmpty = val
}

// GetMaxValue returns the MaxValue property
func (m CreateCustomFieldRequest) GetMaxValue() *float32 {
	return m.MaxValue
}

// SetMaxValue sets the MaxValue property
func (m CreateCustomFieldRequest) SetMaxValue(val *float32) {
	m.MaxValue = val
}

// GetMinValue returns the MinValue property
func (m CreateCustomFieldRequest) GetMinValue() *float32 {
	return m.MinValue
}

// SetMinValue sets the MinValue property
func (m CreateCustomFieldRequest) SetMinValue(val *float32) {
	m.MinValue = val
}

// GetName returns the Name property
func (m CreateCustomFieldRequest) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m CreateCustomFieldRequest) SetName(val string) {
	m.Name = val
}

// GetSearchable returns the Searchable property
func (m CreateCustomFieldRequest) GetSearchable() bool {
	return m.Searchable
}

// SetSearchable sets the Searchable property
func (m CreateCustomFieldRequest) SetSearchable(val bool) {
	m.Searchable = val
}

// GetType returns the Type property
func (m CreateCustomFieldRequest) GetType() FieldType {
	return m.Type
}

// SetType sets the Type property
func (m CreateCustomFieldRequest) SetType(val FieldType) {
	m.Type = val
}

// GetValidValues returns the ValidValues property
func (m CreateCustomFieldRequest) GetValidValues() []string {
	return m.ValidValues
}

// SetValidValues sets the ValidValues property
func (m CreateCustomFieldRequest) SetValidValues(val []string) {
	m.ValidValues = val
}
