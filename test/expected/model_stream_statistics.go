// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// StreamStatistics is an object. Various statistics about the stream
type StreamStatistics struct {
	// DependantCount: How many resources depend on this stream
	DependantCount int32 `json:"dependantCount"`
	// DependencyCount: How many dependencies this resources stream has
	DependencyCount int32 `json:"dependencyCount"`
}

// GetDependantCount returns the DependantCount property
func (m StreamStatistics) GetDependantCount() int32 {
	return m.DependantCount
}

// SetDependantCount sets the DependantCount property
func (m StreamStatistics) SetDependantCount(val int32) {
	m.DependantCount = val
}

// GetDependencyCount returns the DependencyCount property
func (m StreamStatistics) GetDependencyCount() int32 {
	return m.DependencyCount
}

// SetDependencyCount sets the DependencyCount property
func (m StreamStatistics) SetDependencyCount(val int32) {
	m.DependencyCount = val
}
