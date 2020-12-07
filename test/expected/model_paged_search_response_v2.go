// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// PagedSearchResponseV2 is an object.
type PagedSearchResponseV2 struct {
	// Data:
	Data []SearchResponse `json:"data"`
	// FieldSpecs: Map from resource kind to fieldspec array
	FieldSpecs map[string][]FieldSpec `json:"fieldSpecs"`
	// Page: Contains the pagination metadata for a response
	Page PageInfo `json:"page"`
}

// GetData returns the Data property
func (m PagedSearchResponseV2) GetData() []SearchResponse {
	return m.Data
}

// SetData sets the Data property
func (m PagedSearchResponseV2) SetData(val []SearchResponse) {
	m.Data = val
}

// GetFieldSpecs returns the FieldSpecs property
func (m PagedSearchResponseV2) GetFieldSpecs() map[string][]FieldSpec {
	return m.FieldSpecs
}

// SetFieldSpecs sets the FieldSpecs property
func (m PagedSearchResponseV2) SetFieldSpecs(val map[string][]FieldSpec) {
	m.FieldSpecs = val
}

// GetPage returns the Page property
func (m PagedSearchResponseV2) GetPage() PageInfo {
	return m.Page
}

// SetPage sets the Page property
func (m PagedSearchResponseV2) SetPage(val PageInfo) {
	m.Page = val
}
