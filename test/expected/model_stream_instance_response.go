// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// StreamInstanceResponse is an object.
type StreamInstanceResponse struct {
	// Data:
	Data StreamResponse `json:"data"`
	// Fieldspec:
	Fieldspec []FieldSpec `json:"fieldspec"`
}

// GetData returns the Data property
func (m StreamInstanceResponse) GetData() StreamResponse {
	return m.Data
}

// SetData sets the Data property
func (m StreamInstanceResponse) SetData(val StreamResponse) {
	m.Data = val
}

// GetFieldspec returns the Fieldspec property
func (m StreamInstanceResponse) GetFieldspec() []FieldSpec {
	return m.Fieldspec
}

// SetFieldspec sets the Fieldspec property
func (m StreamInstanceResponse) SetFieldspec(val []FieldSpec) {
	m.Fieldspec = val
}
