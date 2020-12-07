// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// ModelInstanceResponse is an object.
type ModelInstanceResponse struct {
	// Data:
	Data ModelResponse `json:"data"`
	// Fieldspec:
	Fieldspec []FieldSpec `json:"fieldspec"`
}

// GetData returns the Data property
func (m ModelInstanceResponse) GetData() ModelResponse {
	return m.Data
}

// SetData sets the Data property
func (m ModelInstanceResponse) SetData(val ModelResponse) {
	m.Data = val
}

// GetFieldspec returns the Fieldspec property
func (m ModelInstanceResponse) GetFieldspec() []FieldSpec {
	return m.Fieldspec
}

// SetFieldspec sets the Fieldspec property
func (m ModelInstanceResponse) SetFieldspec(val []FieldSpec) {
	m.Fieldspec = val
}
