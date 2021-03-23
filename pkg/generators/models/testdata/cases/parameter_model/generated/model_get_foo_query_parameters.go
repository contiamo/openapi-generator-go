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

// GetFooQueryParameters is an object.
type GetFooQueryParameters struct {
	// Param1:
	Param1 string `json:"param1,omitempty"`
	// Id:
	Id string `json:"id"`
	// Param2:
	Param2 int32 `json:"param2,omitempty"`
	// Param3:
	Param3 []string `json:"param3,omitempty"`
}

// Validate implements basic validation for this model
func (m GetFooQueryParameters) Validate() error {
	return validation.Errors{
		"param1": validation.Validate(
			m.Param1, is.UUID,
		),
		"id": validation.Validate(
			m.Id, validation.Required, is.UUID,
		),
		"param2": validation.Validate(
			m.Param2, validation.Min(int32(0)), validation.Max(int32(10)),
		),
		"param3": validation.Validate(
			m.Param3,
		),
	}.Filter()
}

// GetParam1 returns the Param1 property
func (m GetFooQueryParameters) GetParam1() string {
	return m.Param1
}

// SetParam1 sets the Param1 property
func (m GetFooQueryParameters) SetParam1(val string) {
	m.Param1 = val
}

// GetId returns the Id property
func (m GetFooQueryParameters) GetId() string {
	return m.Id
}

// SetId sets the Id property
func (m GetFooQueryParameters) SetId(val string) {
	m.Id = val
}

// GetParam2 returns the Param2 property
func (m GetFooQueryParameters) GetParam2() int32 {
	return m.Param2
}

// SetParam2 sets the Param2 property
func (m GetFooQueryParameters) SetParam2(val int32) {
	m.Param2 = val
}

// GetParam3 returns the Param3 property
func (m GetFooQueryParameters) GetParam3() []string {
	return m.Param3
}

// SetParam3 sets the Param3 property
func (m GetFooQueryParameters) SetParam3(val []string) {
	m.Param3 = val
}
