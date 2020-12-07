// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

// ResourceKind is an enum. A list of supported resource kinds
type ResourceKind string

var (
	ResourceKindApi          ResourceKind = "api"
	ResourceKindApplication  ResourceKind = "application"
	ResourceKindBireport     ResourceKind = "bireport"
	ResourceKindColumn       ResourceKind = "column"
	ResourceKindDatasource   ResourceKind = "datasource"
	ResourceKindGlossaryitem ResourceKind = "glossaryitem"
	ResourceKindModel        ResourceKind = "model"
	ResourceKindPipeline     ResourceKind = "pipeline"
	ResourceKindStream       ResourceKind = "stream"
	ResourceKindTable        ResourceKind = "table"
	ResourceKindUsecase      ResourceKind = "usecase"

	// KnownResourceKind is the list of valid ResourceKind
	KnownResourceKind = []ResourceKind{
		ResourceKindApi,
		ResourceKindApplication,
		ResourceKindBireport,
		ResourceKindColumn,
		ResourceKindDatasource,
		ResourceKindGlossaryitem,
		ResourceKindModel,
		ResourceKindPipeline,
		ResourceKindStream,
		ResourceKindTable,
		ResourceKindUsecase,
	}
	// KnownResourceKindString is the list of valid ResourceKind as string
	KnownResourceKindString = []string{
		string(ResourceKindApi),
		string(ResourceKindApplication),
		string(ResourceKindBireport),
		string(ResourceKindColumn),
		string(ResourceKindDatasource),
		string(ResourceKindGlossaryitem),
		string(ResourceKindModel),
		string(ResourceKindPipeline),
		string(ResourceKindStream),
		string(ResourceKindTable),
		string(ResourceKindUsecase),
	}

	// InKnownResourceKind is an ozzo-validator for ResourceKind
	InKnownResourceKind = validation.In(
		ResourceKindApi,
		ResourceKindApplication,
		ResourceKindBireport,
		ResourceKindColumn,
		ResourceKindDatasource,
		ResourceKindGlossaryitem,
		ResourceKindModel,
		ResourceKindPipeline,
		ResourceKindStream,
		ResourceKindTable,
		ResourceKindUsecase,
	)
)
