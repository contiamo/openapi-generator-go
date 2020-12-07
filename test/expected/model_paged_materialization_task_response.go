// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// PagedMaterializationTaskResponse is an object.
type PagedMaterializationTaskResponse struct {
	// Data:
	Data []MaterializationTaskResponse `json:"data"`
	// Page: Contains the pagination metadata for a response
	Page PageInfo `json:"page"`
}

// GetData returns the Data property
func (m PagedMaterializationTaskResponse) GetData() []MaterializationTaskResponse {
	return m.Data
}

// SetData sets the Data property
func (m PagedMaterializationTaskResponse) SetData(val []MaterializationTaskResponse) {
	m.Data = val
}

// GetPage returns the Page property
func (m PagedMaterializationTaskResponse) GetPage() PageInfo {
	return m.Page
}

// SetPage sets the Page property
func (m PagedMaterializationTaskResponse) SetPage(val PageInfo) {
	m.Page = val
}
