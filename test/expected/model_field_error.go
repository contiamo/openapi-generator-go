// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// FieldError is an object.
type FieldError struct {
	// Key:
	Key string `json:"key"`
	// Message:
	Message string `json:"message"`
	// Type: The type of the error response
	Type ErrorType `json:"type"`
}

// GetKey returns the Key property
func (m FieldError) GetKey() string {
	return m.Key
}

// SetKey sets the Key property
func (m FieldError) SetKey(val string) {
	m.Key = val
}

// GetMessage returns the Message property
func (m FieldError) GetMessage() string {
	return m.Message
}

// SetMessage sets the Message property
func (m FieldError) SetMessage(val string) {
	m.Message = val
}

// GetType returns the Type property
func (m FieldError) GetType() ErrorType {
	return m.Type
}

// SetType sets the Type property
func (m FieldError) SetType(val ErrorType) {
	m.Type = val
}
