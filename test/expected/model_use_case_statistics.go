// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// UseCaseStatistics is an object. Various statistics about the use case
type UseCaseStatistics struct {
	// DependantCount: How many resources depend on this use case
	DependantCount int32 `json:"dependantCount"`
	// DependencyCount: How many dependencies this resources use case has
	DependencyCount int32 `json:"dependencyCount"`
}

// GetDependantCount returns the DependantCount property
func (m UseCaseStatistics) GetDependantCount() int32 {
	return m.DependantCount
}

// SetDependantCount sets the DependantCount property
func (m UseCaseStatistics) SetDependantCount(val int32) {
	m.DependantCount = val
}

// GetDependencyCount returns the DependencyCount property
func (m UseCaseStatistics) GetDependencyCount() int32 {
	return m.DependencyCount
}

// SetDependencyCount sets the DependencyCount property
func (m UseCaseStatistics) SetDependencyCount(val int32) {
	m.DependencyCount = val
}
