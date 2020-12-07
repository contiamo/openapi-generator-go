// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// UpdateCustomFieldRequest is an object.
type UpdateCustomFieldRequest struct {
	// Description: A description of the purpose of this field
	Description *string `json:"description,omitempty"`
	// DisplayName: A nice human-readable name for that field
	DisplayName *string `json:"displayName,omitempty"`
	// Filterable: Should this field show up as a filter in the catalog
	Filterable *bool `json:"filterable,omitempty"`
	// HideIfEmpty: Should this field be hidden when its value is empty
	HideIfEmpty *bool `json:"hideIfEmpty,omitempty"`
	// MaxValue:
	MaxValue *float32 `json:"maxValue,omitempty"`
	// MinValue:
	MinValue *float32 `json:"minValue,omitempty"`
	// Searchable: Should this field be used in the catalog fuzzy search
	Searchable *bool `json:"searchable,omitempty"`
	// ValidValues:
	ValidValues []string `json:"validValues,omitempty"`
}

// GetDescription returns the Description property
func (m UpdateCustomFieldRequest) GetDescription() *string {
	return m.Description
}

// SetDescription sets the Description property
func (m UpdateCustomFieldRequest) SetDescription(val *string) {
	m.Description = val
}

// GetDisplayName returns the DisplayName property
func (m UpdateCustomFieldRequest) GetDisplayName() *string {
	return m.DisplayName
}

// SetDisplayName sets the DisplayName property
func (m UpdateCustomFieldRequest) SetDisplayName(val *string) {
	m.DisplayName = val
}

// GetFilterable returns the Filterable property
func (m UpdateCustomFieldRequest) GetFilterable() *bool {
	return m.Filterable
}

// SetFilterable sets the Filterable property
func (m UpdateCustomFieldRequest) SetFilterable(val *bool) {
	m.Filterable = val
}

// GetHideIfEmpty returns the HideIfEmpty property
func (m UpdateCustomFieldRequest) GetHideIfEmpty() *bool {
	return m.HideIfEmpty
}

// SetHideIfEmpty sets the HideIfEmpty property
func (m UpdateCustomFieldRequest) SetHideIfEmpty(val *bool) {
	m.HideIfEmpty = val
}

// GetMaxValue returns the MaxValue property
func (m UpdateCustomFieldRequest) GetMaxValue() *float32 {
	return m.MaxValue
}

// SetMaxValue sets the MaxValue property
func (m UpdateCustomFieldRequest) SetMaxValue(val *float32) {
	m.MaxValue = val
}

// GetMinValue returns the MinValue property
func (m UpdateCustomFieldRequest) GetMinValue() *float32 {
	return m.MinValue
}

// SetMinValue sets the MinValue property
func (m UpdateCustomFieldRequest) SetMinValue(val *float32) {
	m.MinValue = val
}

// GetSearchable returns the Searchable property
func (m UpdateCustomFieldRequest) GetSearchable() *bool {
	return m.Searchable
}

// SetSearchable sets the Searchable property
func (m UpdateCustomFieldRequest) SetSearchable(val *bool) {
	m.Searchable = val
}

// GetValidValues returns the ValidValues property
func (m UpdateCustomFieldRequest) GetValidValues() []string {
	return m.ValidValues
}

// SetValidValues sets the ValidValues property
func (m UpdateCustomFieldRequest) SetValidValues(val []string) {
	m.ValidValues = val
}
