// Code generated by openapi-generator-go DO NOT EDIT.
//
// Source:
//
//	Title: Test
//	Version: 0.1.0
package generatortest

import (
	"encoding/json"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// User is an object.
type User struct {
	// DateOfBirth:
	DateOfBirth string `json:"date_of_birth,omitempty" mapstructure:"date_of_birth,omitempty"`
	// Name:
	Name string `json:"name,omitempty" mapstructure:"name,omitempty"`
	// UserId:
	UserId string `json:"user_id,omitempty" mapstructure:"user_id,omitempty"`
}

// NewUser instantiates a new User with default values overriding them as follows:
// 1. Default values specified in the User schema
// 2. Default values specified per User property
func NewUser() *User {
	m := &User{}

	return m
}

// UnmarshalJSON implements the json.Unmarshaler interface for User. It set the default values for the User type
func (m *User) UnmarshalJSON(data []byte) error {
	// Set default values
	*m = *NewUser()

	// Unmarshal using an alias to avoid an infinite loop
	type alias User
	err := json.Unmarshal(data, (*alias)(m))
	if err != nil {
		return err
	}
	return nil
}

// Validate implements basic validation for this model
func (m User) Validate() error {
	errors := validation.Errors{
		"dateOfBirth": validation.Validate(
			m.DateOfBirth, validation.Date("2006-01-02"),
		),
	}
	return errors.Filter()
}

// GetDateOfBirth returns the DateOfBirth property
func (m User) GetDateOfBirth() string {
	return m.DateOfBirth
}

// SetDateOfBirth sets the DateOfBirth property
func (m *User) SetDateOfBirth(val string) {
	m.DateOfBirth = val
}

// GetName returns the Name property
func (m User) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *User) SetName(val string) {
	m.Name = val
}

// GetUserId returns the UserId property
func (m User) GetUserId() string {
	return m.UserId
}

// SetUserId sets the UserId property
func (m *User) SetUserId(val string) {
	m.UserId = val
}
