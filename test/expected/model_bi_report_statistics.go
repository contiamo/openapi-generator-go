// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// BIReportStatistics is an object. Various statistics about the BI Report
type BIReportStatistics struct {
	// DependantCount: How many resources depend on this BI Report
	DependantCount int32 `json:"dependantCount"`
	// DependencyCount: How many dependencies this resources BI Report has
	DependencyCount int32 `json:"dependencyCount"`
}

// GetDependantCount returns the DependantCount property
func (m BIReportStatistics) GetDependantCount() int32 {
	return m.DependantCount
}

// SetDependantCount sets the DependantCount property
func (m BIReportStatistics) SetDependantCount(val int32) {
	m.DependantCount = val
}

// GetDependencyCount returns the DependencyCount property
func (m BIReportStatistics) GetDependencyCount() int32 {
	return m.DependencyCount
}

// SetDependencyCount sets the DependencyCount property
func (m BIReportStatistics) SetDependencyCount(val int32) {
	m.DependencyCount = val
}
