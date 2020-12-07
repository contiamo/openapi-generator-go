// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

// QueryFileFormat is an enum.
type QueryFileFormat string

var (
	QueryFileFormatCsv  QueryFileFormat = "csv"
	QueryFileFormatXlsx QueryFileFormat = "xlsx"

	// KnownQueryFileFormat is the list of valid QueryFileFormat
	KnownQueryFileFormat = []QueryFileFormat{
		QueryFileFormatCsv,
		QueryFileFormatXlsx,
	}
	// KnownQueryFileFormatString is the list of valid QueryFileFormat as string
	KnownQueryFileFormatString = []string{
		string(QueryFileFormatCsv),
		string(QueryFileFormatXlsx),
	}

	// InKnownQueryFileFormat is an ozzo-validator for QueryFileFormat
	InKnownQueryFileFormat = validation.In(
		QueryFileFormatCsv,
		QueryFileFormatXlsx,
	)
)
