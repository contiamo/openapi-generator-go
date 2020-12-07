// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// ColumnFieldSpec is an object.
type ColumnFieldSpec struct {
	// Columns:
	Columns []FieldSpec `json:"columns"`
}

// GetColumns returns the Columns property
func (m ColumnFieldSpec) GetColumns() []FieldSpec {
	return m.Columns
}

// SetColumns sets the Columns property
func (m ColumnFieldSpec) SetColumns(val []FieldSpec) {
	m.Columns = val
}
