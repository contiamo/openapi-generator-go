// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Test
//     Version: 0.1.0
package generatortest

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

// UserEntity is an object. This is a short description of a user entity used for permission listing or assignments in other services.
type UserEntity struct {
	// Email:
	Email string `json:"email"`
	// Id:
	Id string `json:"id"`
	// Name:
	Name string `json:"name"`
}

// Validate implements basic validation for this model
func (m UserEntity) Validate() error {
	return validation.Errors{
		"email": validation.Validate(
			m.Email, validation.NotNil, is.EmailFormat,
		),
		"id": validation.Validate(
			m.Id, validation.NotNil, is.UUID,
		),
	}.Filter()
}

// GetEmail returns the Email property
func (m UserEntity) GetEmail() string {
	return m.Email
}

// SetEmail sets the Email property
func (m UserEntity) SetEmail(val string) {
	m.Email = val
}

// GetId returns the Id property
func (m UserEntity) GetId() string {
	return m.Id
}

// SetId sets the Id property
func (m UserEntity) SetId(val string) {
	m.Id = val
}

// GetName returns the Name property
func (m UserEntity) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m UserEntity) SetName(val string) {
	m.Name = val
}
