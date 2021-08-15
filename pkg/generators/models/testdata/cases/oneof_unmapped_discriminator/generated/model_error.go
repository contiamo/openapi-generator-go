// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Test oneOf discriminator support
//     Version: 0.1.0
package generatortest

import (
	"encoding/json"
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/mitchellh/mapstructure"
)

type errorer interface {
	GetKind() string
	Validate() error
}

// Error is a oneOf type.
type Error struct {
	data errorer
}

// MarshalJSON implements the json.Marshaller interface
func (m Error) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.data)
}

// UnmarshalJSON implements the json.Unmarshaller interface
func (m *Error) UnmarshalJSON(bs []byte) error {
	discriminator := struct {
		value string `json:"kind"`
	}{}
	err := json.Unmarshal(bs, &discriminator)
	if err != nil {
		return err
	}

	if discriminator.value == "" {
		return validation.Errors{
			"kind": fmt.Errorf("can not be empty"),
		}.Filter()
	}

	switch discriminator.value {
	case "ExternalError":
		data := ExternalError{}
		err := json.Unmarshal(bs, &data)
		if err != nil {
			return err
		}
		m.data = data
	case "FieldError":
		data := FieldError{}
		err := json.Unmarshal(bs, &data)
		if err != nil {
			return err
		}
		m.data = data
	case "GenericError":
		data := GenericError{}
		err := json.Unmarshal(bs, &data)
		if err != nil {
			return err
		}
		m.data = data
	}
	return nil
}

// FromGenericError sets the Error data.
func (m *Error) FromGenericError(data GenericError) {
	m.data = data
}

// FromFieldError sets the Error data.
func (m *Error) FromFieldError(data FieldError) {
	m.data = data
}

// FromExternalError sets the Error data.
func (m *Error) FromExternalError(data ExternalError) {
	m.data = data
}

// As converts Error to a user defined structure.
func (m Error) As(target interface{}) error {
	return mapstructure.Decode(m.data, target)
}

// AsGenericError converts Error to a GenericError
func (m Error) AsGenericError() (result GenericError, err error) {
	return result, mapstructure.Decode(m.data, &result)
}

// AsFieldError converts Error to a FieldError
func (m Error) AsFieldError() (result FieldError, err error) {
	return result, mapstructure.Decode(m.data, &result)
}

// AsExternalError converts Error to a ExternalError
func (m Error) AsExternalError() (result ExternalError, err error) {
	return result, mapstructure.Decode(m.data, &result)
}

// Validate implements basic validation for this model
func (m Error) Validate() error {
	discriminator := m.data.GetKind()
	switch discriminator {
	case "ExternalError":
		return m.Validate()
	case "FieldError":
		return m.Validate()
	case "GenericError":
		return m.Validate()
	default:
		return validation.Errors{
			"kind": fmt.Errorf("unknown kind value"),
		}.Filter()
	}
}

// IsError tests if data is one of the discriminated sub-types of Error.
func IsError(data interface{}) bool {
	_, ok := data.(errorer)
	return ok
}
