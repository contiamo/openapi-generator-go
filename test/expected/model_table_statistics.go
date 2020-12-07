// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// TableStatistics is an object. Various statistics about the table
type TableStatistics struct {
	// ColumnCount: How many columns the table
	ColumnCount int32 `json:"columnCount"`
}

// GetColumnCount returns the ColumnCount property
func (m TableStatistics) GetColumnCount() int32 {
	return m.ColumnCount
}

// SetColumnCount sets the ColumnCount property
func (m TableStatistics) SetColumnCount(val int32) {
	m.ColumnCount = val
}
