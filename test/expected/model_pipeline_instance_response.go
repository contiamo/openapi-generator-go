// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// PipelineInstanceResponse is an object.
type PipelineInstanceResponse struct {
	// Data:
	Data PipelineResponse `json:"data"`
	// Fieldspec:
	Fieldspec []FieldSpec `json:"fieldspec"`
}

// GetData returns the Data property
func (m PipelineInstanceResponse) GetData() PipelineResponse {
	return m.Data
}

// SetData sets the Data property
func (m PipelineInstanceResponse) SetData(val PipelineResponse) {
	m.Data = val
}

// GetFieldspec returns the Fieldspec property
func (m PipelineInstanceResponse) GetFieldspec() []FieldSpec {
	return m.Fieldspec
}

// SetFieldspec sets the Fieldspec property
func (m PipelineInstanceResponse) SetFieldspec(val []FieldSpec) {
	m.Fieldspec = val
}
