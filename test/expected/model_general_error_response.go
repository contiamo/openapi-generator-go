// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// GeneralErrorResponse is an object. General error response that usually has a very generic message
type GeneralErrorResponse struct {
	// Errors:
	Errors []GeneralError `json:"errors"`
}

// GetErrors returns the Errors property
func (m GeneralErrorResponse) GetErrors() []GeneralError {
	return m.Errors
}

// SetErrors sets the Errors property
func (m GeneralErrorResponse) SetErrors(val []GeneralError) {
	m.Errors = val
}
