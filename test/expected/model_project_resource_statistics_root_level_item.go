// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// ProjectResourceStatisticsRootLevelItem is an object.
type ProjectResourceStatisticsRootLevelItem struct {
	// Category: A value that describes which resource kinds (`ResourceKind`) to include in the search response.
	Category SearchCategory `json:"category"`
	// Children:
	Children []ProjectResourceStatisticsItem `json:"children"`
	// TotalCount: How many resources are there in this category in the project
	TotalCount int32 `json:"totalCount"`
}

// GetCategory returns the Category property
func (m ProjectResourceStatisticsRootLevelItem) GetCategory() SearchCategory {
	return m.Category
}

// SetCategory sets the Category property
func (m ProjectResourceStatisticsRootLevelItem) SetCategory(val SearchCategory) {
	m.Category = val
}

// GetChildren returns the Children property
func (m ProjectResourceStatisticsRootLevelItem) GetChildren() []ProjectResourceStatisticsItem {
	return m.Children
}

// SetChildren sets the Children property
func (m ProjectResourceStatisticsRootLevelItem) SetChildren(val []ProjectResourceStatisticsItem) {
	m.Children = val
}

// GetTotalCount returns the TotalCount property
func (m ProjectResourceStatisticsRootLevelItem) GetTotalCount() int32 {
	return m.TotalCount
}

// SetTotalCount sets the TotalCount property
func (m ProjectResourceStatisticsRootLevelItem) SetTotalCount(val int32) {
	m.TotalCount = val
}
