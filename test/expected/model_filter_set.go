// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// FilterSet is an object. A set of filters which should all be applied
type FilterSet struct {
	// Filters:
	Filters []Filter `json:"filters"`
}

// GetFilters returns the Filters property
func (m FilterSet) GetFilters() []Filter {
	return m.Filters
}

// SetFilters sets the Filters property
func (m FilterSet) SetFilters(val []Filter) {
	m.Filters = val
}
