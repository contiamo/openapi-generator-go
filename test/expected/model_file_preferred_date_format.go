// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

// FilePreferredDateFormat is an enum. FilePreferredDateFormat determines how to parse ambiguous dates, such as '02-02-2010'.
type FilePreferredDateFormat string

var (
	FilePreferredDateFormatDayFirst   FilePreferredDateFormat = "dayFirst"
	FilePreferredDateFormatMonthFirst FilePreferredDateFormat = "monthFirst"
	FilePreferredDateFormatYearFirst  FilePreferredDateFormat = "yearFirst"

	// KnownFilePreferredDateFormat is the list of valid FilePreferredDateFormat
	KnownFilePreferredDateFormat = []FilePreferredDateFormat{
		FilePreferredDateFormatDayFirst,
		FilePreferredDateFormatMonthFirst,
		FilePreferredDateFormatYearFirst,
	}
	// KnownFilePreferredDateFormatString is the list of valid FilePreferredDateFormat as string
	KnownFilePreferredDateFormatString = []string{
		string(FilePreferredDateFormatDayFirst),
		string(FilePreferredDateFormatMonthFirst),
		string(FilePreferredDateFormatYearFirst),
	}

	// InKnownFilePreferredDateFormat is an ozzo-validator for FilePreferredDateFormat
	InKnownFilePreferredDateFormat = validation.In(
		FilePreferredDateFormatDayFirst,
		FilePreferredDateFormatMonthFirst,
		FilePreferredDateFormatYearFirst,
	)
)
