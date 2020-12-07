// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

// TableType is an enum.
type TableType string

var (
	TableTypeMaterializationTable TableType = "MaterializationTable"
	TableTypeRegularTable         TableType = "RegularTable"
	TableTypeViewTable            TableType = "ViewTable"

	// KnownTableType is the list of valid TableType
	KnownTableType = []TableType{
		TableTypeMaterializationTable,
		TableTypeRegularTable,
		TableTypeViewTable,
	}
	// KnownTableTypeString is the list of valid TableType as string
	KnownTableTypeString = []string{
		string(TableTypeMaterializationTable),
		string(TableTypeRegularTable),
		string(TableTypeViewTable),
	}

	// InKnownTableType is an ozzo-validator for TableType
	InKnownTableType = validation.In(
		TableTypeMaterializationTable,
		TableTypeRegularTable,
		TableTypeViewTable,
	)
)
