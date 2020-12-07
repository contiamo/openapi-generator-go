// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

// SearchOrder is an enum.
type SearchOrder string

var (
	SearchOrderCreatedAt   SearchOrder = "createdAt"
	SearchOrderDisplayName SearchOrder = "displayName"
	SearchOrderName        SearchOrder = "name"
	SearchOrderResourceId  SearchOrder = "resourceId"
	SearchOrderUpdatedAt   SearchOrder = "updatedAt"

	// KnownSearchOrder is the list of valid SearchOrder
	KnownSearchOrder = []SearchOrder{
		SearchOrderCreatedAt,
		SearchOrderDisplayName,
		SearchOrderName,
		SearchOrderResourceId,
		SearchOrderUpdatedAt,
	}
	// KnownSearchOrderString is the list of valid SearchOrder as string
	KnownSearchOrderString = []string{
		string(SearchOrderCreatedAt),
		string(SearchOrderDisplayName),
		string(SearchOrderName),
		string(SearchOrderResourceId),
		string(SearchOrderUpdatedAt),
	}

	// InKnownSearchOrder is an ozzo-validator for SearchOrder
	InKnownSearchOrder = validation.In(
		SearchOrderCreatedAt,
		SearchOrderDisplayName,
		SearchOrderName,
		SearchOrderResourceId,
		SearchOrderUpdatedAt,
	)
)
