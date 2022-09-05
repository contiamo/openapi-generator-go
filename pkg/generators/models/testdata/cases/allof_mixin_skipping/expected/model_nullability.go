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

// Nullability is an enum. Column nullability
type Nullability string

// Validate implements basic validation for this model
func (m Nullability) Validate() error {
	return InKnownNullability.Validate(m)
}

var (
	NullabilityNoNulls         Nullability = "NoNulls"
	NullabilityNullable        Nullability = "Nullable"
	NullabilityNullableUnknown Nullability = "NullableUnknown"

	// KnownNullability is the list of valid Nullability
	KnownNullability = []Nullability{
		NullabilityNoNulls,
		NullabilityNullable,
		NullabilityNullableUnknown,
	}
	// KnownNullabilityString is the list of valid Nullability as string
	KnownNullabilityString = []string{
		string(NullabilityNoNulls),
		string(NullabilityNullable),
		string(NullabilityNullableUnknown),
	}

	// InKnownNullability is an ozzo-validator for Nullability
	InKnownNullability = validation.In(
		NullabilityNoNulls,
		NullabilityNullable,
		NullabilityNullableUnknown,
	)
)
