// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

// ErrorType is an enum. The type of the error response
type ErrorType string

var (
	ErrorTypeFieldError   ErrorType = "FieldError"
	ErrorTypeGeneralError ErrorType = "GeneralError"

	// KnownErrorType is the list of valid ErrorType
	KnownErrorType = []ErrorType{
		ErrorTypeFieldError,
		ErrorTypeGeneralError,
	}
	// KnownErrorTypeString is the list of valid ErrorType as string
	KnownErrorTypeString = []string{
		string(ErrorTypeFieldError),
		string(ErrorTypeGeneralError),
	}

	// InKnownErrorType is an ozzo-validator for ErrorType
	InKnownErrorType = validation.In(
		ErrorTypeFieldError,
		ErrorTypeGeneralError,
	)
)
