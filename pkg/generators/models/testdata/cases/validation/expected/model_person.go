// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Test
//     Version: 0.1.0
package generatortest

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

// Person is an object.
type Person struct {
	// Age:
	Age float32 `json:"age,omitempty"`
	// Gender:
	Gender Gender `json:"gender"`
	// Name:
	Name string `json:"name"`
}

// Validate implements basic validation for this model
func (m Person) Validate() error {
	return validation.Errors{
		"age": validation.Validate(
			m.Age, validation.Min(0), validation.Max(120),
		),
		"gender": validation.Validate(
			m.Gender, validation.Required, InKnownGender,
		),
		"name": validation.Validate(
			m.Name, validation.Required, validation.Length(2, 32),
		),
	}.Filter()
}

// GetAge returns the Age property
func (m Person) GetAge() float32 {
	return m.Age
}

// SetAge sets the Age property
func (m Person) SetAge(val float32) {
	m.Age = val
}

// GetGender returns the Gender property
func (m Person) GetGender() Gender {
	return m.Gender
}

// SetGender sets the Gender property
func (m Person) SetGender(val Gender) {
	m.Gender = val
}

// GetName returns the Name property
func (m Person) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m Person) SetName(val string) {
	m.Name = val
}
