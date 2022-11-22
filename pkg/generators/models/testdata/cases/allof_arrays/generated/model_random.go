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

// Random is an object.
type Random struct {
	// List:
	List []Item `json:"list"`
}

// Validate implements basic validation for this model
func (m Random) Validate() error {
	return validation.Errors{
		"list": validation.Validate(
			m.List, validation.NotNil,
		),
	}.Filter()
}

// GetList returns the List property
func (m Random) GetList() []Item {
	return m.List
}

// SetList sets the List property
func (m *Random) SetList(val []Item) {
	m.List = val
}
