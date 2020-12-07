// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

// BIReportType is an enum.
type BIReportType string

var (
	BIReportTypeContiamo BIReportType = "contiamo"
	BIReportTypeOther    BIReportType = "other"

	// KnownBIReportType is the list of valid BIReportType
	KnownBIReportType = []BIReportType{
		BIReportTypeContiamo,
		BIReportTypeOther,
	}
	// KnownBIReportTypeString is the list of valid BIReportType as string
	KnownBIReportTypeString = []string{
		string(BIReportTypeContiamo),
		string(BIReportTypeOther),
	}

	// InKnownBIReportType is an ozzo-validator for BIReportType
	InKnownBIReportType = validation.In(
		BIReportTypeContiamo,
		BIReportTypeOther,
	)
)
