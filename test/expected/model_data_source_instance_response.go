// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// DataSourceInstanceResponse is an object.
type DataSourceInstanceResponse struct {
	// Data:
	Data DataSourceResponse `json:"data"`
	// Fieldspec:
	Fieldspec []FieldSpec `json:"fieldspec"`
}

// GetData returns the Data property
func (m DataSourceInstanceResponse) GetData() DataSourceResponse {
	return m.Data
}

// SetData sets the Data property
func (m DataSourceInstanceResponse) SetData(val DataSourceResponse) {
	m.Data = val
}

// GetFieldspec returns the Fieldspec property
func (m DataSourceInstanceResponse) GetFieldspec() []FieldSpec {
	return m.Fieldspec
}

// SetFieldspec sets the Fieldspec property
func (m DataSourceInstanceResponse) SetFieldspec(val []FieldSpec) {
	m.Fieldspec = val
}
