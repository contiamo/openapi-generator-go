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

// Random is an object.
type Random struct {
	// List:
	List []Item `json:"list" mapstructure:"list"`
}

// NewRandom instantiates a new Random with default values overriding them as follows:
// 1. Default values specified in the Random schema
// 2. Default values specified per Random property
func NewRandom() *Random {
	m := &Random{}

	return m
}

// UnmarshalJSON implements the json.Unmarshaler interface for Random. It set the default values for the Random type
func (m *Random) UnmarshalJSON(data []byte) error {
	// Set default values
	*m = *NewRandom()

	// Unmarshal using an alias to avoid an infinite loop
	type alias Random
	err := json.Unmarshal(data, (*alias)(m))
	if err != nil {
		return err
	}
	return nil
}

// Validate implements basic validation for this model
func (m Random) Validate() error {
	errors := validation.Errors{
		"list": validation.Validate(
			m.List, validation.NotNil,
		),
	}
	return errors.Filter()
}

// GetList returns the List property
func (m Random) GetList() []Item {
	return m.List
}

// SetList sets the List property
func (m *Random) SetList(val []Item) {
	m.List = val
}
