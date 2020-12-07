// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// ApplicationInstanceResponse is an object.
type ApplicationInstanceResponse struct {
	// Data:
	Data ApplicationResponse `json:"data"`
	// Fieldspec:
	Fieldspec []FieldSpec `json:"fieldspec"`
}

// GetData returns the Data property
func (m ApplicationInstanceResponse) GetData() ApplicationResponse {
	return m.Data
}

// SetData sets the Data property
func (m ApplicationInstanceResponse) SetData(val ApplicationResponse) {
	m.Data = val
}

// GetFieldspec returns the Fieldspec property
func (m ApplicationInstanceResponse) GetFieldspec() []FieldSpec {
	return m.Fieldspec
}

// SetFieldspec sets the Fieldspec property
func (m ApplicationInstanceResponse) SetFieldspec(val []FieldSpec) {
	m.Fieldspec = val
}
