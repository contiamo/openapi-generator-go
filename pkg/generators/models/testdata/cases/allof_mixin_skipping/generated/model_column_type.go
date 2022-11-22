// This file is auto-generated, DO NOT EDIT.
//
// Source:
//
//	Title: Test
//	Version: 0.1.0
package generatortest

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// ColumnType is an enum. Normalized column type. If type cannot be determined or is not compatible, then 'other'.
type ColumnType string

// Validate implements basic validation for this model
func (m ColumnType) Validate() error {
	return InKnownColumnType.Validate(m)
}

var (
	ColumnTypeBigint          ColumnType = "bigint"
	ColumnTypeBoolean         ColumnType = "boolean"
	ColumnTypeDate            ColumnType = "date"
	ColumnTypeDoublePrecision ColumnType = "double precision"
	ColumnTypeFloat           ColumnType = "float"
	ColumnTypeInteger         ColumnType = "integer"
	ColumnTypeNumeric         ColumnType = "numeric"
	ColumnTypeOther           ColumnType = "other"
	ColumnTypeSmallint        ColumnType = "smallint"
	ColumnTypeTime            ColumnType = "time"
	ColumnTypeTimestamp       ColumnType = "timestamp"
	ColumnTypeVarchar         ColumnType = "varchar"
	ColumnTypeVarcharArray    ColumnType = "varchar array"

	// KnownColumnType is the list of valid ColumnType
	KnownColumnType = []ColumnType{
		ColumnTypeBigint,
		ColumnTypeBoolean,
		ColumnTypeDate,
		ColumnTypeDoublePrecision,
		ColumnTypeFloat,
		ColumnTypeInteger,
		ColumnTypeNumeric,
		ColumnTypeOther,
		ColumnTypeSmallint,
		ColumnTypeTime,
		ColumnTypeTimestamp,
		ColumnTypeVarchar,
		ColumnTypeVarcharArray,
	}
	// KnownColumnTypeString is the list of valid ColumnType as string
	KnownColumnTypeString = []string{
		string(ColumnTypeBigint),
		string(ColumnTypeBoolean),
		string(ColumnTypeDate),
		string(ColumnTypeDoublePrecision),
		string(ColumnTypeFloat),
		string(ColumnTypeInteger),
		string(ColumnTypeNumeric),
		string(ColumnTypeOther),
		string(ColumnTypeSmallint),
		string(ColumnTypeTime),
		string(ColumnTypeTimestamp),
		string(ColumnTypeVarchar),
		string(ColumnTypeVarcharArray),
	}

	// InKnownColumnType is an ozzo-validator for ColumnType
	InKnownColumnType = validation.In(
		ColumnTypeBigint,
		ColumnTypeBoolean,
		ColumnTypeDate,
		ColumnTypeDoublePrecision,
		ColumnTypeFloat,
		ColumnTypeInteger,
		ColumnTypeNumeric,
		ColumnTypeOther,
		ColumnTypeSmallint,
		ColumnTypeTime,
		ColumnTypeTimestamp,
		ColumnTypeVarchar,
		ColumnTypeVarcharArray,
	)
)
