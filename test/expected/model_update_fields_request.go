// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// UpdateFieldsRequest is an object. The `fields` array should contain all of the fields in the desired order.
//
// Additionally, we require all field specifications _must_ contain the following fields:
//
//   *   `category`
type UpdateFieldsRequest struct {
	// Fields:
	Fields []FieldSpecUpdateRequest `json:"fields"`
}

// GetFields returns the Fields property
func (m UpdateFieldsRequest) GetFields() []FieldSpecUpdateRequest {
	return m.Fields
}

// SetFields sets the Fields property
func (m UpdateFieldsRequest) SetFields(val []FieldSpecUpdateRequest) {
	m.Fields = val
}
