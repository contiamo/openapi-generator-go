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

// BaseEntity is an object. Contains shared properties for all the entities
type BaseEntity struct {
	// Id:
	Id string `json:"id"`
	// Name:
	Name string `json:"name"`
}

// Validate implements basic validation for this model
func (m BaseEntity) Validate() error {
	return validation.Errors{
		"id": validation.Validate(
			m.Id, validation.NotNil, is.UUID,
		),
	}.Filter()
}

// GetId returns the Id property
func (m BaseEntity) GetId() string {
	return m.Id
}

// SetId sets the Id property
func (m BaseEntity) SetId(val string) {
	m.Id = val
}

// GetName returns the Name property
func (m BaseEntity) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m BaseEntity) SetName(val string) {
	m.Name = val
}
