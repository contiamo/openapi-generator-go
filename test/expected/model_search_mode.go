// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

// SearchMode is an enum. The search mode to use
// * 'searchCombined' will search for datasources and tables
// * 'searchDatasources' will only search for datasources
// * 'searchTables' will only search for tables
type SearchMode string

var (
	SearchModeSearchCombined    SearchMode = "searchCombined"
	SearchModeSearchDatasources SearchMode = "searchDatasources"
	SearchModeSearchTables      SearchMode = "searchTables"

	// KnownSearchMode is the list of valid SearchMode
	KnownSearchMode = []SearchMode{
		SearchModeSearchCombined,
		SearchModeSearchDatasources,
		SearchModeSearchTables,
	}
	// KnownSearchModeString is the list of valid SearchMode as string
	KnownSearchModeString = []string{
		string(SearchModeSearchCombined),
		string(SearchModeSearchDatasources),
		string(SearchModeSearchTables),
	}

	// InKnownSearchMode is an ozzo-validator for SearchMode
	InKnownSearchMode = validation.In(
		SearchModeSearchCombined,
		SearchModeSearchDatasources,
		SearchModeSearchTables,
	)
)
