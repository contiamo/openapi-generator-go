// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// FieldErrorResponse is an object. Error message that contains detailed information about certain parameters being incorrect
type FieldErrorResponse struct {
	// Errors:
	Errors []FieldError `json:"errors"`
}

// GetErrors returns the Errors property
func (m FieldErrorResponse) GetErrors() []FieldError {
	return m.Errors
}

// SetErrors sets the Errors property
func (m FieldErrorResponse) SetErrors(val []FieldError) {
	m.Errors = val
}
