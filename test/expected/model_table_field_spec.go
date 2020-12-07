// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// TableFieldSpec is an object.
type TableFieldSpec struct {
	// Columns:
	Columns []FieldSpec `json:"columns"`
	// Tables:
	Tables []FieldSpec `json:"tables"`
}

// GetColumns returns the Columns property
func (m TableFieldSpec) GetColumns() []FieldSpec {
	return m.Columns
}

// SetColumns sets the Columns property
func (m TableFieldSpec) SetColumns(val []FieldSpec) {
	m.Columns = val
}

// GetTables returns the Tables property
func (m TableFieldSpec) GetTables() []FieldSpec {
	return m.Tables
}

// SetTables sets the Tables property
func (m TableFieldSpec) SetTables(val []FieldSpec) {
	m.Tables = val
}
