// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// ResourceStatistics is an object. Various statistics about the resource
type ResourceStatistics struct {
	// ChildCount: How many child resources are connected to it
	ChildCount int32 `json:"childCount"`
}

// GetChildCount returns the ChildCount property
func (m ResourceStatistics) GetChildCount() int32 {
	return m.ChildCount
}

// SetChildCount sets the ChildCount property
func (m ResourceStatistics) SetChildCount(val int32) {
	m.ChildCount = val
}
