// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// ApplicationStatistics is an object. Various statistics about the application
type ApplicationStatistics struct {
	// DependantCount: How many resources depend on this application
	DependantCount int32 `json:"dependantCount"`
	// DependencyCount: How many dependencies this resources application has
	DependencyCount int32 `json:"dependencyCount"`
}

// GetDependantCount returns the DependantCount property
func (m ApplicationStatistics) GetDependantCount() int32 {
	return m.DependantCount
}

// SetDependantCount sets the DependantCount property
func (m ApplicationStatistics) SetDependantCount(val int32) {
	m.DependantCount = val
}

// GetDependencyCount returns the DependencyCount property
func (m ApplicationStatistics) GetDependencyCount() int32 {
	return m.DependencyCount
}

// SetDependencyCount sets the DependencyCount property
func (m ApplicationStatistics) SetDependencyCount(val int32) {
	m.DependencyCount = val
}
