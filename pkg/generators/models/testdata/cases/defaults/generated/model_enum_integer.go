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

// EnumInteger is an enum.
type EnumInteger int32

// DefaultEnumInteger is the default value for EnumInteger
var DefaultEnumInteger = EnumInteger(2)

// Validate implements basic validation for this model
func (m EnumInteger) Validate() error {
	return InKnownEnumInteger.Validate(m)
}

var (
	EnumInteger0 EnumInteger = 0
	EnumInteger1 EnumInteger = 1
	EnumInteger2 EnumInteger = 2

	// KnownEnumInteger is the list of valid EnumInteger
	KnownEnumInteger = []EnumInteger{
		EnumInteger0,
		EnumInteger1,
		EnumInteger2,
	}

	// KnownEnumIntegerInt32 is the list of valid EnumInteger as int32
	KnownEnumIntegerInt32 = []int32{
		int32(EnumInteger0),
		int32(EnumInteger1),
		int32(EnumInteger2),
	}

	// InKnownEnumInteger is an ozzo-validator for EnumInteger
	InKnownEnumInteger = validation.In(
		EnumInteger0,
		EnumInteger1,
		EnumInteger2,
	)
)
