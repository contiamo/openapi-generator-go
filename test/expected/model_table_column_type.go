// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

// TableColumnType is an enum. The data type of the table column as it was discovered from the data source.
type TableColumnType string

var (
	TableColumnTypeBigint          TableColumnType = "bigint"
	TableColumnTypeBoolean         TableColumnType = "boolean"
	TableColumnTypeDate            TableColumnType = "date"
	TableColumnTypeDoublePrecision TableColumnType = "double precision"
	TableColumnTypeFloat           TableColumnType = "float"
	TableColumnTypeInteger         TableColumnType = "integer"
	TableColumnTypeIntegerArray    TableColumnType = "integer array"
	TableColumnTypeNumeric         TableColumnType = "numeric"
	TableColumnTypeOther           TableColumnType = "other"
	TableColumnTypeSmallint        TableColumnType = "smallint"
	TableColumnTypeTime            TableColumnType = "time"
	TableColumnTypeTimestamp       TableColumnType = "timestamp"
	TableColumnTypeVarchar         TableColumnType = "varchar"
	TableColumnTypeVarcharArray    TableColumnType = "varchar array"

	// KnownTableColumnType is the list of valid TableColumnType
	KnownTableColumnType = []TableColumnType{
		TableColumnTypeBigint,
		TableColumnTypeBoolean,
		TableColumnTypeDate,
		TableColumnTypeDoublePrecision,
		TableColumnTypeFloat,
		TableColumnTypeInteger,
		TableColumnTypeIntegerArray,
		TableColumnTypeNumeric,
		TableColumnTypeOther,
		TableColumnTypeSmallint,
		TableColumnTypeTime,
		TableColumnTypeTimestamp,
		TableColumnTypeVarchar,
		TableColumnTypeVarcharArray,
	}
	// KnownTableColumnTypeString is the list of valid TableColumnType as string
	KnownTableColumnTypeString = []string{
		string(TableColumnTypeBigint),
		string(TableColumnTypeBoolean),
		string(TableColumnTypeDate),
		string(TableColumnTypeDoublePrecision),
		string(TableColumnTypeFloat),
		string(TableColumnTypeInteger),
		string(TableColumnTypeIntegerArray),
		string(TableColumnTypeNumeric),
		string(TableColumnTypeOther),
		string(TableColumnTypeSmallint),
		string(TableColumnTypeTime),
		string(TableColumnTypeTimestamp),
		string(TableColumnTypeVarchar),
		string(TableColumnTypeVarcharArray),
	}

	// InKnownTableColumnType is an ozzo-validator for TableColumnType
	InKnownTableColumnType = validation.In(
		TableColumnTypeBigint,
		TableColumnTypeBoolean,
		TableColumnTypeDate,
		TableColumnTypeDoublePrecision,
		TableColumnTypeFloat,
		TableColumnTypeInteger,
		TableColumnTypeIntegerArray,
		TableColumnTypeNumeric,
		TableColumnTypeOther,
		TableColumnTypeSmallint,
		TableColumnTypeTime,
		TableColumnTypeTimestamp,
		TableColumnTypeVarchar,
		TableColumnTypeVarcharArray,
	)
)
