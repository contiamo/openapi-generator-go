// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// PageInfo is an object. Contains the pagination metadata for a response
type PageInfo struct {
	// Current: The current page number using 1-based array indexing
	Current int32 `json:"current"`
	// ItemCount: Total number of items
	ItemCount int32 `json:"itemCount"`
	// ItemsPerPage: Maximum items that can be on the page. They may be different from
	// the requested number of times
	ItemsPerPage int32 `json:"itemsPerPage"`
	// UnfilteredItemCount: Item count if filters were not applied
	UnfilteredItemCount int32 `json:"unfilteredItemCount"`
}

// GetCurrent returns the Current property
func (m PageInfo) GetCurrent() int32 {
	return m.Current
}

// SetCurrent sets the Current property
func (m PageInfo) SetCurrent(val int32) {
	m.Current = val
}

// GetItemCount returns the ItemCount property
func (m PageInfo) GetItemCount() int32 {
	return m.ItemCount
}

// SetItemCount sets the ItemCount property
func (m PageInfo) SetItemCount(val int32) {
	m.ItemCount = val
}

// GetItemsPerPage returns the ItemsPerPage property
func (m PageInfo) GetItemsPerPage() int32 {
	return m.ItemsPerPage
}

// SetItemsPerPage sets the ItemsPerPage property
func (m PageInfo) SetItemsPerPage(val int32) {
	m.ItemsPerPage = val
}

// GetUnfilteredItemCount returns the UnfilteredItemCount property
func (m PageInfo) GetUnfilteredItemCount() int32 {
	return m.UnfilteredItemCount
}

// SetUnfilteredItemCount sets the UnfilteredItemCount property
func (m PageInfo) SetUnfilteredItemCount(val int32) {
	m.UnfilteredItemCount = val
}
