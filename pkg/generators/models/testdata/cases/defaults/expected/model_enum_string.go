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

// EnumString is an enum.
type EnumString string

// DefaultEnumString is the default value for EnumString
var DefaultEnumString = EnumString("bar")

// Validate implements basic validation for this model
func (m EnumString) Validate() error {
	return InKnownEnumString.Validate(m)
}

var (
	EnumStringBar EnumString = "bar"
	EnumStringFoo EnumString = "foo"

	// KnownEnumString is the list of valid EnumString
	KnownEnumString = []EnumString{
		EnumStringBar,
		EnumStringFoo,
	}

	// KnownEnumStringString is the list of valid EnumString as string
	KnownEnumStringString = []string{
		string(EnumStringBar),
		string(EnumStringFoo),
	}

	// InKnownEnumString is an ozzo-validator for EnumString
	InKnownEnumString = validation.In(
		EnumStringBar,
		EnumStringFoo,
	)
)
