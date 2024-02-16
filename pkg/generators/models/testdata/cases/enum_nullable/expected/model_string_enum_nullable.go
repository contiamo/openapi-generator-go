// Code generated by openapi-generator-go DO NOT EDIT.
//
// Source:
//
//	Title: Test
//	Version: 0.1.0
package generatortest

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// StringEnumNullable is an enum.
type StringEnumNullable string

// Validate implements basic validation for this model
func (m StringEnumNullable) Validate() error {
	return InKnownStringEnumNullable.Validate(m)
}

var (
	StringEnumNullableBar StringEnumNullable = "bar"
	StringEnumNullableFoo StringEnumNullable = "foo"

	// KnownStringEnumNullable is the list of valid StringEnumNullable
	KnownStringEnumNullable = []StringEnumNullable{
		StringEnumNullableBar,
		StringEnumNullableFoo,
	}
	// KnownStringEnumNullableString is the list of valid StringEnumNullable as string
	KnownStringEnumNullableString = []string{
		string(StringEnumNullableBar),
		string(StringEnumNullableFoo),
	}

	// InKnownStringEnumNullable is an ozzo-validator for StringEnumNullable
	InKnownStringEnumNullable = validation.In(
		StringEnumNullableBar,
		StringEnumNullableFoo,
	)
)