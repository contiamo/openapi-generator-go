// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// UseCaseInstanceResponse is an object.
type UseCaseInstanceResponse struct {
	// Data:
	Data UseCaseResponse `json:"data"`
	// Fieldspec:
	Fieldspec []FieldSpec `json:"fieldspec"`
}

// GetData returns the Data property
func (m UseCaseInstanceResponse) GetData() UseCaseResponse {
	return m.Data
}

// SetData sets the Data property
func (m UseCaseInstanceResponse) SetData(val UseCaseResponse) {
	m.Data = val
}

// GetFieldspec returns the Fieldspec property
func (m UseCaseInstanceResponse) GetFieldspec() []FieldSpec {
	return m.Fieldspec
}

// SetFieldspec sets the Fieldspec property
func (m UseCaseInstanceResponse) SetFieldspec(val []FieldSpec) {
	m.Fieldspec = val
}
