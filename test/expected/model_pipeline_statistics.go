// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// PipelineStatistics is an object. Various statistics about the pipeline
type PipelineStatistics struct {
	// DependantCount: How many resources depend on this pipeline
	DependantCount int32 `json:"dependantCount"`
	// DependencyCount: How many dependencies this resource has
	DependencyCount int32 `json:"dependencyCount"`
}

// GetDependantCount returns the DependantCount property
func (m PipelineStatistics) GetDependantCount() int32 {
	return m.DependantCount
}

// SetDependantCount sets the DependantCount property
func (m PipelineStatistics) SetDependantCount(val int32) {
	m.DependantCount = val
}

// GetDependencyCount returns the DependencyCount property
func (m PipelineStatistics) GetDependencyCount() int32 {
	return m.DependencyCount
}

// SetDependencyCount sets the DependencyCount property
func (m PipelineStatistics) SetDependencyCount(val int32) {
	m.DependencyCount = val
}
