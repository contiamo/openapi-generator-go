// This file is auto-generated, DO NOT EDIT.
//
// Source:
//
//	Title: Test
//	Version: 0.1.0
package generatortest

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Person is an object.
type Person struct {
	// Age:
	Age int32 `json:"age,omitempty"`
	// Name:
	Name string `json:"name,omitempty"`
}

// Validate implements basic validation for this model
func (m Person) Validate() error {
	return validation.Errors{}.Filter()
}

// GetAge returns the Age property
func (m Person) GetAge() int32 {
	return m.Age
}

// SetAge sets the Age property
func (m *Person) SetAge(val int32) {
	m.Age = val
}

// GetName returns the Name property
func (m Person) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *Person) SetName(val string) {
	m.Name = val
}
