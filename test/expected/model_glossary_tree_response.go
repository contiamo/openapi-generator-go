// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// GlossaryTreeResponse is an object.
type GlossaryTreeResponse struct {
	// Data:
	Data []GlossaryTree `json:"data"`
	// Fieldspec:
	Fieldspec []FieldSpec `json:"fieldspec"`
}

// GetData returns the Data property
func (m GlossaryTreeResponse) GetData() []GlossaryTree {
	return m.Data
}

// SetData sets the Data property
func (m GlossaryTreeResponse) SetData(val []GlossaryTree) {
	m.Data = val
}

// GetFieldspec returns the Fieldspec property
func (m GlossaryTreeResponse) GetFieldspec() []FieldSpec {
	return m.Fieldspec
}

// SetFieldspec sets the Fieldspec property
func (m GlossaryTreeResponse) SetFieldspec(val []FieldSpec) {
	m.Fieldspec = val
}
