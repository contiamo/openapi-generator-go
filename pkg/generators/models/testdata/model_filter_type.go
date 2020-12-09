// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package testpkg

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

// FilterType is an enum.
type FilterType string

var (
	FilterTypeHideAll      FilterType = "hideAll"
	FilterTypeHideOnlyExpr FilterType = "hideOnlyExpr"
	FilterTypeShowAll      FilterType = "showAll"
	FilterTypeShowOnlyExpr FilterType = "showOnlyExpr"

	// KnownFilterType is the list of valid FilterType
	KnownFilterType = []FilterType{
		FilterTypeHideAll,
		FilterTypeHideOnlyExpr,
		FilterTypeShowAll,
		FilterTypeShowOnlyExpr,
	}
	// KnownFilterTypeString is the list of valid FilterType as string
	KnownFilterTypeString = []string{
		string(FilterTypeHideAll),
		string(FilterTypeHideOnlyExpr),
		string(FilterTypeShowAll),
		string(FilterTypeShowOnlyExpr),
	}

	// InKnownFilterType is an ozzo-validator for FilterType
	InKnownFilterType = validation.In(
		FilterTypeHideAll,
		FilterTypeHideOnlyExpr,
		FilterTypeShowAll,
		FilterTypeShowOnlyExpr,
	)
)
