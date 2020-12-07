// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

// FieldType is an enum. The data type of the field:
//
// - `text` is an arbitrary string
//
// - `multitext` is an arbitrary array of strings
//
// - `select` a value which can be selected from a list
//
// - `multiselect` is a `select` with multiple selected values, represented as an array
//
// - `boolean` is `true` or `false` value
//
// - `date` is a date string in `yyyy-mm-dd` format, e.g. `2019-08-21`
//
// - `datetime` is an RFC3339 datetime string, e.g. `2019-08-21T11:54:30.369917Z`
//
// - `number` is a numeric value, described by RFC8259, section 6.
//
// - `link` is a link object with a displayed text and URL, e.g. `{"displayText": "text", "url": "http://example.com"}`. Only absolute URLs with HTTP/HTTPS protocol are allowed.
type FieldType string

var (
	FieldTypeBoolean     FieldType = "boolean"
	FieldTypeDate        FieldType = "date"
	FieldTypeDatetime    FieldType = "datetime"
	FieldTypeLink        FieldType = "link"
	FieldTypeMultiselect FieldType = "multiselect"
	FieldTypeMultitext   FieldType = "multitext"
	FieldTypeNumber      FieldType = "number"
	FieldTypeSelect      FieldType = "select"
	FieldTypeText        FieldType = "text"

	// KnownFieldType is the list of valid FieldType
	KnownFieldType = []FieldType{
		FieldTypeBoolean,
		FieldTypeDate,
		FieldTypeDatetime,
		FieldTypeLink,
		FieldTypeMultiselect,
		FieldTypeMultitext,
		FieldTypeNumber,
		FieldTypeSelect,
		FieldTypeText,
	}
	// KnownFieldTypeString is the list of valid FieldType as string
	KnownFieldTypeString = []string{
		string(FieldTypeBoolean),
		string(FieldTypeDate),
		string(FieldTypeDatetime),
		string(FieldTypeLink),
		string(FieldTypeMultiselect),
		string(FieldTypeMultitext),
		string(FieldTypeNumber),
		string(FieldTypeSelect),
		string(FieldTypeText),
	}

	// InKnownFieldType is an ozzo-validator for FieldType
	InKnownFieldType = validation.In(
		FieldTypeBoolean,
		FieldTypeDate,
		FieldTypeDatetime,
		FieldTypeLink,
		FieldTypeMultiselect,
		FieldTypeMultitext,
		FieldTypeNumber,
		FieldTypeSelect,
		FieldTypeText,
	)
)
