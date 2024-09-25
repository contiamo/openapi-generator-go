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

// Top is an object.
type Top struct {
	// Arr:
	Arr []Sub1 `json:"arr,omitempty" mapstructure:"arr,omitempty"`
	// Boo: Type alias for a value type
	Boo Sub3 `json:"boo,omitempty" mapstructure:"boo,omitempty"`
	// Obj:
	Obj *Sub1 `json:"obj,omitempty" mapstructure:"obj,omitempty"`
}

// NewTop instantiates a new Top with default values overriding them as follows:
// 1. Default values specified in the Top schema
// 2. Default values specified per Top property
func NewTop() *Top {
	m := &Top{
		Obj: NewSub1(),
	}

	return m
}

// UnmarshalJSON implements the json.Unmarshaler interface for Top. It set the default values for the Top type
func (m *Top) UnmarshalJSON(data []byte) error {
	// Set default values
	*m = *NewTop()

	// Unmarshal using an alias to avoid an infinite loop
	type alias Top
	err := json.Unmarshal(data, (*alias)(m))
	if err != nil {
		return err
	}
	return nil
}

// Validate implements basic validation for this model
func (m Top) Validate() error {
	errors := validation.Errors{
		"arr": validation.Validate(
			m.Arr,
		),
		"boo": validation.Validate(
			m.Boo,
		),
		"obj": validation.Validate(
			m.Obj,
		),
	}
	return errors.Filter()
}

// GetArr returns the Arr property
func (m Top) GetArr() []Sub1 {
	return m.Arr
}

// SetArr sets the Arr property
func (m *Top) SetArr(val []Sub1) {
	m.Arr = val
}

// GetBoo returns the Boo property
func (m Top) GetBoo() Sub3 {
	return m.Boo
}

// SetBoo sets the Boo property
func (m *Top) SetBoo(val Sub3) {
	m.Boo = val
}

// GetObj returns the Obj property
func (m Top) GetObj() *Sub1 {
	return m.Obj
}

// SetObj sets the Obj property
func (m *Top) SetObj(val *Sub1) {
	m.Obj = val
}
