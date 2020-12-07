// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// ColumnInstanceResponse is an object.
type ColumnInstanceResponse struct {
	// Data:
	Data ColumnResponse `json:"data"`
	// Fieldspec:
	Fieldspec ColumnFieldSpec `json:"fieldspec"`
}

// GetData returns the Data property
func (m ColumnInstanceResponse) GetData() ColumnResponse {
	return m.Data
}

// SetData sets the Data property
func (m ColumnInstanceResponse) SetData(val ColumnResponse) {
	m.Data = val
}

// GetFieldspec returns the Fieldspec property
func (m ColumnInstanceResponse) GetFieldspec() ColumnFieldSpec {
	return m.Fieldspec
}

// SetFieldspec sets the Fieldspec property
func (m ColumnInstanceResponse) SetFieldspec(val ColumnFieldSpec) {
	m.Fieldspec = val
}
