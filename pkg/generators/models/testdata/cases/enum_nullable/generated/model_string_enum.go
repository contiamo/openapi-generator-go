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

// StringEnum is an enum.
type StringEnum string

// Validate implements basic validation for this model
func (m StringEnum) Validate() error {
	return InKnownStringEnum.Validate(m)
}

var (
	StringEnumBar StringEnum = "bar"
	StringEnumFoo StringEnum = "foo"

	// KnownStringEnum is the list of valid StringEnum
	KnownStringEnum = []StringEnum{
		StringEnumBar,
		StringEnumFoo,
	}
	// KnownStringEnumString is the list of valid StringEnum as string
	KnownStringEnumString = []string{
		string(StringEnumBar),
		string(StringEnumFoo),
	}

	// InKnownStringEnum is an ozzo-validator for StringEnum
	InKnownStringEnum = validation.In(
		StringEnumBar,
		StringEnumFoo,
	)
)
