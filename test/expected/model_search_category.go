// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

// SearchCategory is an enum. A value that describes which resource kinds (`ResourceKind`) to include in the search response.
type SearchCategory string

var (
	SearchCategoryApi         SearchCategory = "api"
	SearchCategoryApplication SearchCategory = "application"
	SearchCategoryBi          SearchCategory = "bi"
	SearchCategoryData        SearchCategory = "data"
	SearchCategoryGlossary    SearchCategory = "glossary"
	SearchCategoryModel       SearchCategory = "model"
	SearchCategoryPipeline    SearchCategory = "pipeline"
	SearchCategoryStream      SearchCategory = "stream"
	SearchCategoryUsecase     SearchCategory = "usecase"

	// KnownSearchCategory is the list of valid SearchCategory
	KnownSearchCategory = []SearchCategory{
		SearchCategoryApi,
		SearchCategoryApplication,
		SearchCategoryBi,
		SearchCategoryData,
		SearchCategoryGlossary,
		SearchCategoryModel,
		SearchCategoryPipeline,
		SearchCategoryStream,
		SearchCategoryUsecase,
	}
	// KnownSearchCategoryString is the list of valid SearchCategory as string
	KnownSearchCategoryString = []string{
		string(SearchCategoryApi),
		string(SearchCategoryApplication),
		string(SearchCategoryBi),
		string(SearchCategoryData),
		string(SearchCategoryGlossary),
		string(SearchCategoryModel),
		string(SearchCategoryPipeline),
		string(SearchCategoryStream),
		string(SearchCategoryUsecase),
	}

	// InKnownSearchCategory is an ozzo-validator for SearchCategory
	InKnownSearchCategory = validation.In(
		SearchCategoryApi,
		SearchCategoryApplication,
		SearchCategoryBi,
		SearchCategoryData,
		SearchCategoryGlossary,
		SearchCategoryModel,
		SearchCategoryPipeline,
		SearchCategoryStream,
		SearchCategoryUsecase,
	)
)
