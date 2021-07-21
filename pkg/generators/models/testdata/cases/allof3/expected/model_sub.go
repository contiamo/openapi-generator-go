// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Test
//     Version: 0.1.0
package generatortest

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Sub is an object. this is a sub object
type Sub struct {
	// SubField:
	SubField string `json:"subField,omitempty"`
}

// Validate implements basic validation for this model
func (m Sub) Validate() error {
	return validation.Errors{}.Filter()
}

// GetSubField returns the SubField property
func (m Sub) GetSubField() string {
	return m.SubField
}

// SetSubField sets the SubField property
func (m *Sub) SetSubField(val string) {
	m.SubField = val
}
