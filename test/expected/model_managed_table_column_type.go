// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

// ManagedTableColumnType is an enum. The data type of the column in a managed table
type ManagedTableColumnType string

var (
	ManagedTableColumnTypeBool     ManagedTableColumnType = "bool"
	ManagedTableColumnTypeDate     ManagedTableColumnType = "date"
	ManagedTableColumnTypeDatetime ManagedTableColumnType = "datetime"
	ManagedTableColumnTypeDecimal  ManagedTableColumnType = "decimal"
	ManagedTableColumnTypeInteger  ManagedTableColumnType = "integer"
	ManagedTableColumnTypeText     ManagedTableColumnType = "text"
	ManagedTableColumnTypeTime     ManagedTableColumnType = "time"

	// KnownManagedTableColumnType is the list of valid ManagedTableColumnType
	KnownManagedTableColumnType = []ManagedTableColumnType{
		ManagedTableColumnTypeBool,
		ManagedTableColumnTypeDate,
		ManagedTableColumnTypeDatetime,
		ManagedTableColumnTypeDecimal,
		ManagedTableColumnTypeInteger,
		ManagedTableColumnTypeText,
		ManagedTableColumnTypeTime,
	}
	// KnownManagedTableColumnTypeString is the list of valid ManagedTableColumnType as string
	KnownManagedTableColumnTypeString = []string{
		string(ManagedTableColumnTypeBool),
		string(ManagedTableColumnTypeDate),
		string(ManagedTableColumnTypeDatetime),
		string(ManagedTableColumnTypeDecimal),
		string(ManagedTableColumnTypeInteger),
		string(ManagedTableColumnTypeText),
		string(ManagedTableColumnTypeTime),
	}

	// InKnownManagedTableColumnType is an ozzo-validator for ManagedTableColumnType
	InKnownManagedTableColumnType = validation.In(
		ManagedTableColumnTypeBool,
		ManagedTableColumnTypeDate,
		ManagedTableColumnTypeDatetime,
		ManagedTableColumnTypeDecimal,
		ManagedTableColumnTypeInteger,
		ManagedTableColumnTypeText,
		ManagedTableColumnTypeTime,
	)
)
