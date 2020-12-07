// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// BIReportInstanceResponse is an object.
type BIReportInstanceResponse struct {
	// Data:
	Data BIReportResponse `json:"data"`
	// Fieldspec:
	Fieldspec []FieldSpec `json:"fieldspec"`
}

// GetData returns the Data property
func (m BIReportInstanceResponse) GetData() BIReportResponse {
	return m.Data
}

// SetData sets the Data property
func (m BIReportInstanceResponse) SetData(val BIReportResponse) {
	m.Data = val
}

// GetFieldspec returns the Fieldspec property
func (m BIReportInstanceResponse) GetFieldspec() []FieldSpec {
	return m.Fieldspec
}

// SetFieldspec sets the Fieldspec property
func (m BIReportInstanceResponse) SetFieldspec(val []FieldSpec) {
	m.Fieldspec = val
}
