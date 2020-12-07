// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// GeneralError is an object.
type GeneralError struct {
	// Message:
	Message string `json:"message"`
	// Type: The type of the error response
	Type ErrorType `json:"type"`
}

// GetMessage returns the Message property
func (m GeneralError) GetMessage() string {
	return m.Message
}

// SetMessage sets the Message property
func (m GeneralError) SetMessage(val string) {
	m.Message = val
}

// GetType returns the Type property
func (m GeneralError) GetType() ErrorType {
	return m.Type
}

// SetType sets the Type property
func (m GeneralError) SetType(val ErrorType) {
	m.Type = val
}
