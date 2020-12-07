// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// APIStatistics is an object. Various statistics about the API
type APIStatistics struct {
	// DependantCount: How many resources depend on this API
	DependantCount int32 `json:"dependantCount"`
	// DependencyCount: How many dependencies this API has
	DependencyCount int32 `json:"dependencyCount"`
}

// GetDependantCount returns the DependantCount property
func (m APIStatistics) GetDependantCount() int32 {
	return m.DependantCount
}

// SetDependantCount sets the DependantCount property
func (m APIStatistics) SetDependantCount(val int32) {
	m.DependantCount = val
}

// GetDependencyCount returns the DependencyCount property
func (m APIStatistics) GetDependencyCount() int32 {
	return m.DependencyCount
}

// SetDependencyCount sets the DependencyCount property
func (m APIStatistics) SetDependencyCount(val int32) {
	m.DependencyCount = val
}
