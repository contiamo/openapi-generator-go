// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// FieldsResponse is an object. The `fields` will contain all of the custom metadata fields as well as the system fields,
// in order.
type FieldsResponse struct {
	// Fields:
	Fields []FieldSpec `json:"fields"`
}

// GetFields returns the Fields property
func (m FieldsResponse) GetFields() []FieldSpec {
	return m.Fields
}

// SetFields sets the Fields property
func (m FieldsResponse) SetFields(val []FieldSpec) {
	m.Fields = val
}
