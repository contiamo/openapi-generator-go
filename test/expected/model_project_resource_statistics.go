// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// ProjectResourceStatistics is an object. Contains various statistics on resources in a project
type ProjectResourceStatistics struct {
	// Data:
	Data []ProjectResourceStatisticsRootLevelItem `json:"data"`
}

// GetData returns the Data property
func (m ProjectResourceStatistics) GetData() []ProjectResourceStatisticsRootLevelItem {
	return m.Data
}

// SetData sets the Data property
func (m ProjectResourceStatistics) SetData(val []ProjectResourceStatisticsRootLevelItem) {
	m.Data = val
}
