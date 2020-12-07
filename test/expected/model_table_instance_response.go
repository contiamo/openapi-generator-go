// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// TableInstanceResponse is an object.
type TableInstanceResponse struct {
	// Data:
	Data TableResponse `json:"data"`
	// Fieldspec:
	Fieldspec TableFieldSpec `json:"fieldspec"`
}

// GetData returns the Data property
func (m TableInstanceResponse) GetData() TableResponse {
	return m.Data
}

// SetData sets the Data property
func (m TableInstanceResponse) SetData(val TableResponse) {
	m.Data = val
}

// GetFieldspec returns the Fieldspec property
func (m TableInstanceResponse) GetFieldspec() TableFieldSpec {
	return m.Fieldspec
}

// SetFieldspec sets the Fieldspec property
func (m TableInstanceResponse) SetFieldspec(val TableFieldSpec) {
	m.Fieldspec = val
}
