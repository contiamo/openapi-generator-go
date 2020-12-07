// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// ProjectResourceStatisticsItem is an object.
type ProjectResourceStatisticsItem struct {
	// Children:
	Children []ProjectResourceStatisticsItem `json:"children"`
	// Count: How many resources of this kind are there in the project
	Count int32 `json:"count"`
	// Kind: A list of supported resource kinds
	Kind ResourceKind `json:"kind"`
}

// GetChildren returns the Children property
func (m ProjectResourceStatisticsItem) GetChildren() []ProjectResourceStatisticsItem {
	return m.Children
}

// SetChildren sets the Children property
func (m ProjectResourceStatisticsItem) SetChildren(val []ProjectResourceStatisticsItem) {
	m.Children = val
}

// GetCount returns the Count property
func (m ProjectResourceStatisticsItem) GetCount() int32 {
	return m.Count
}

// SetCount sets the Count property
func (m ProjectResourceStatisticsItem) SetCount(val int32) {
	m.Count = val
}

// GetKind returns the Kind property
func (m ProjectResourceStatisticsItem) GetKind() ResourceKind {
	return m.Kind
}

// SetKind sets the Kind property
func (m ProjectResourceStatisticsItem) SetKind(val ResourceKind) {
	m.Kind = val
}
