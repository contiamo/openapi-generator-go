// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// DataSourceStatistics is an object. Various statistics about the data source
type DataSourceStatistics struct {
	// TableCount: How many tables the data source has
	TableCount int32 `json:"tableCount"`
}

// GetTableCount returns the TableCount property
func (m DataSourceStatistics) GetTableCount() int32 {
	return m.TableCount
}

// SetTableCount sets the TableCount property
func (m DataSourceStatistics) SetTableCount(val int32) {
	m.TableCount = val
}
