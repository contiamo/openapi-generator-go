// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Test
//     Version: 0.1.0
package generatortest

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Item is an object.
type Item struct {
	// First:
	First string `json:"first"`
}

// Validate implements basic validation for this model
func (m Item) Validate() error {
	return validation.Errors{}.Filter()
}

// GetFirst returns the First property
func (m Item) GetFirst() string {
	return m.First
}

// SetFirst sets the First property
func (m Item) SetFirst(val string) {
	m.First = val
}
