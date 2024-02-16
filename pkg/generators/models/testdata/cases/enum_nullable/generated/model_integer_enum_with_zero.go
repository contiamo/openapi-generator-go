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

// IntegerEnumWithZero is an enum.
type IntegerEnumWithZero int32

// Validate implements basic validation for this model
func (m IntegerEnumWithZero) Validate() error {
	return InKnownIntegerEnumWithZero.Validate(m)
}

var (
	IntegerEnumWithZero0 IntegerEnumWithZero = 0
	IntegerEnumWithZero1 IntegerEnumWithZero = 1
	IntegerEnumWithZero2 IntegerEnumWithZero = 2

	// KnownIntegerEnumWithZero is the list of valid IntegerEnumWithZero
	KnownIntegerEnumWithZero = []IntegerEnumWithZero{
		IntegerEnumWithZero0,
		IntegerEnumWithZero1,
		IntegerEnumWithZero2,
	}
	// KnownIntegerEnumWithZeroInt32 is the list of valid IntegerEnumWithZero as int32
	KnownIntegerEnumWithZeroInt32 = []int32{
		int32(IntegerEnumWithZero0),
		int32(IntegerEnumWithZero1),
		int32(IntegerEnumWithZero2),
	}

	// InKnownIntegerEnumWithZero is an ozzo-validator for IntegerEnumWithZero
	InKnownIntegerEnumWithZero = validation.In(
		IntegerEnumWithZero0,
		IntegerEnumWithZero1,
		IntegerEnumWithZero2,
	)
)