// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

// QueryHistoryRecordType is an enum.
type QueryHistoryRecordType string

var (
	QueryHistoryRecordTypeCatalog QueryHistoryRecordType = "Catalog"
	QueryHistoryRecordTypeNative  QueryHistoryRecordType = "Native"

	// KnownQueryHistoryRecordType is the list of valid QueryHistoryRecordType
	KnownQueryHistoryRecordType = []QueryHistoryRecordType{
		QueryHistoryRecordTypeCatalog,
		QueryHistoryRecordTypeNative,
	}
	// KnownQueryHistoryRecordTypeString is the list of valid QueryHistoryRecordType as string
	KnownQueryHistoryRecordTypeString = []string{
		string(QueryHistoryRecordTypeCatalog),
		string(QueryHistoryRecordTypeNative),
	}

	// InKnownQueryHistoryRecordType is an ozzo-validator for QueryHistoryRecordType
	InKnownQueryHistoryRecordType = validation.In(
		QueryHistoryRecordTypeCatalog,
		QueryHistoryRecordTypeNative,
	)
)
