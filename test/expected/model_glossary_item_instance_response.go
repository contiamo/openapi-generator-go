// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// GlossaryItemInstanceResponse is an object.
type GlossaryItemInstanceResponse struct {
	// Data:
	Data GlossaryItemResponse `json:"data"`
	// Fieldspec:
	Fieldspec []FieldSpec `json:"fieldspec"`
}

// GetData returns the Data property
func (m GlossaryItemInstanceResponse) GetData() GlossaryItemResponse {
	return m.Data
}

// SetData sets the Data property
func (m GlossaryItemInstanceResponse) SetData(val GlossaryItemResponse) {
	m.Data = val
}

// GetFieldspec returns the Fieldspec property
func (m GlossaryItemInstanceResponse) GetFieldspec() []FieldSpec {
	return m.Fieldspec
}

// SetFieldspec sets the Fieldspec property
func (m GlossaryItemInstanceResponse) SetFieldspec(val []FieldSpec) {
	m.Fieldspec = val
}
