// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// ModelStatistics is an object. Various statistics about the model
type ModelStatistics struct {
	// DependantCount: How many resources depend on this model
	DependantCount int32 `json:"dependantCount"`
	// DependencyCount: How many dependencies this resources model has
	DependencyCount int32 `json:"dependencyCount"`
}

// GetDependantCount returns the DependantCount property
func (m ModelStatistics) GetDependantCount() int32 {
	return m.DependantCount
}

// SetDependantCount sets the DependantCount property
func (m ModelStatistics) SetDependantCount(val int32) {
	m.DependantCount = val
}

// GetDependencyCount returns the DependencyCount property
func (m ModelStatistics) GetDependencyCount() int32 {
	return m.DependencyCount
}

// SetDependencyCount sets the DependencyCount property
func (m ModelStatistics) SetDependencyCount(val int32) {
	m.DependencyCount = val
}
