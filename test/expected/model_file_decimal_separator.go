// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

// FileDecimalSeparator is an enum. FileDecimalSeparator determines how numeric values are parsed, e.g. '10,50'.
type FileDecimalSeparator string

var (
	FileDecimalSeparatorNumericComma  FileDecimalSeparator = "numericComma"
	FileDecimalSeparatorNumericPeriod FileDecimalSeparator = "numericPeriod"

	// KnownFileDecimalSeparator is the list of valid FileDecimalSeparator
	KnownFileDecimalSeparator = []FileDecimalSeparator{
		FileDecimalSeparatorNumericComma,
		FileDecimalSeparatorNumericPeriod,
	}
	// KnownFileDecimalSeparatorString is the list of valid FileDecimalSeparator as string
	KnownFileDecimalSeparatorString = []string{
		string(FileDecimalSeparatorNumericComma),
		string(FileDecimalSeparatorNumericPeriod),
	}

	// InKnownFileDecimalSeparator is an ozzo-validator for FileDecimalSeparator
	InKnownFileDecimalSeparator = validation.In(
		FileDecimalSeparatorNumericComma,
		FileDecimalSeparatorNumericPeriod,
	)
)
