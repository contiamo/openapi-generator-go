// Code generated by openapi-generator-go DO NOT EDIT.
//
// Source:
//
//	Title: Test oneOf discriminator support
//	Version: 0.1.0
package generatortest

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// FieldErrorKind is an enum.
type FieldErrorKind string

// Validate implements basic validation for this model
func (m FieldErrorKind) Validate() error {
	return InKnownFieldErrorKind.Validate(m)
}

var (
	FieldErrorKindFieldError FieldErrorKind = "FieldError"

	// KnownFieldErrorKind is the list of valid FieldErrorKind
	KnownFieldErrorKind = []FieldErrorKind{
		FieldErrorKindFieldError,
	}

	// KnownFieldErrorKindString is the list of valid FieldErrorKind as string
	KnownFieldErrorKindString = []string{
		string(FieldErrorKindFieldError),
	}

	// InKnownFieldErrorKind is an ozzo-validator for FieldErrorKind
	InKnownFieldErrorKind = validation.In(
		FieldErrorKindFieldError,
	)
)