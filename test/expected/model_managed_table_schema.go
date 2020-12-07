// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// ManagedTableSchema is an object. Schema of a file or of the corresponding managed table which will be created
type ManagedTableSchema struct {
	// Columns: Definitions of columns and their types. The order is important and should match the order of columns in the file which is going to be uploaded
	Columns []ManagedTableColumn `json:"columns"`
	// Indexes: Definitions of the managed table indexes. These may be simple or compound indexes, i.e. they may contain one or more columns.
	// The appropriate index will be chosen based on the storage engine.
	Indexes []ManagedTableIndex `json:"indexes"`
}

// GetColumns returns the Columns property
func (m ManagedTableSchema) GetColumns() []ManagedTableColumn {
	return m.Columns
}

// SetColumns sets the Columns property
func (m ManagedTableSchema) SetColumns(val []ManagedTableColumn) {
	m.Columns = val
}

// GetIndexes returns the Indexes property
func (m ManagedTableSchema) GetIndexes() []ManagedTableIndex {
	return m.Indexes
}

// SetIndexes sets the Indexes property
func (m ManagedTableSchema) SetIndexes(val []ManagedTableIndex) {
	m.Indexes = val
}
