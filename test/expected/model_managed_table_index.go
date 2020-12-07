// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// ManagedTableIndex is an object.
type ManagedTableIndex struct {
	// Columns:
	Columns []string `json:"columns"`
	// IndexType: Supported index types
	IndexType IndexTypes `json:"indexType"`
}

// GetColumns returns the Columns property
func (m ManagedTableIndex) GetColumns() []string {
	return m.Columns
}

// SetColumns sets the Columns property
func (m ManagedTableIndex) SetColumns(val []string) {
	m.Columns = val
}

// GetIndexType returns the IndexType property
func (m ManagedTableIndex) GetIndexType() IndexTypes {
	return m.IndexType
}

// SetIndexType sets the IndexType property
func (m ManagedTableIndex) SetIndexType(val IndexTypes) {
	m.IndexType = val
}
