// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Test
//     Version: 0.1.0
package generatortest

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Shape is an object.
type Shape struct {
	// Coordinates:
	Coordinates []Line `json:"coordinates,omitempty"`
	// Type:
	Type string `json:"type,omitempty"`
}

// Validate implements basic validation for this model
func (m Shape) Validate() error {
	return validation.Errors{
		"coordinates": validation.Validate(
			m.Coordinates,
		),
	}.Filter()
}

// GetCoordinates returns the Coordinates property
func (m Shape) GetCoordinates() []Line {
	return m.Coordinates
}

// SetCoordinates sets the Coordinates property
func (m *Shape) SetCoordinates(val []Line) {
	m.Coordinates = val
}

// GetType returns the Type property
func (m Shape) GetType() string {
	return m.Type
}

// SetType sets the Type property
func (m *Shape) SetType(val string) {
	m.Type = val
}
