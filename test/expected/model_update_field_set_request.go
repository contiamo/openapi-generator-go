// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// UpdateFieldSetRequest is an object.
type UpdateFieldSetRequest struct {
	// Fields:
	Fields []string `json:"fields"`
}

// GetFields returns the Fields property
func (m UpdateFieldSetRequest) GetFields() []string {
	return m.Fields
}

// SetFields sets the Fields property
func (m UpdateFieldSetRequest) SetFields(val []string) {
	m.Fields = val
}
