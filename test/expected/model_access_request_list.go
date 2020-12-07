// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// AccessRequestList is an object. A list of access requests for a given resource.
// If you are a owner of the resource, then this will contain all requests for this resource from all users.
// If you are NOT a owner of the resource, you will only see your own requests.
type AccessRequestList struct {
	// Data: List of requests for a given resource
	Data []AccessRequestResponse `json:"data"`
	// Page: Contains the pagination metadata for a response
	Page PageInfo `json:"page"`
}

// GetData returns the Data property
func (m AccessRequestList) GetData() []AccessRequestResponse {
	return m.Data
}

// SetData sets the Data property
func (m AccessRequestList) SetData(val []AccessRequestResponse) {
	m.Data = val
}

// GetPage returns the Page property
func (m AccessRequestList) GetPage() PageInfo {
	return m.Page
}

// SetPage sets the Page property
func (m AccessRequestList) SetPage(val PageInfo) {
	m.Page = val
}
