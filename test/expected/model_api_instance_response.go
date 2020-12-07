// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// APIInstanceResponse is an object.
type APIInstanceResponse struct {
	// Data:
	Data APIResponse `json:"data"`
	// Fieldspec:
	Fieldspec []FieldSpec `json:"fieldspec"`
}

// GetData returns the Data property
func (m APIInstanceResponse) GetData() APIResponse {
	return m.Data
}

// SetData sets the Data property
func (m APIInstanceResponse) SetData(val APIResponse) {
	m.Data = val
}

// GetFieldspec returns the Fieldspec property
func (m APIInstanceResponse) GetFieldspec() []FieldSpec {
	return m.Fieldspec
}

// SetFieldspec sets the Fieldspec property
func (m APIInstanceResponse) SetFieldspec(val []FieldSpec) {
	m.Fieldspec = val
}
