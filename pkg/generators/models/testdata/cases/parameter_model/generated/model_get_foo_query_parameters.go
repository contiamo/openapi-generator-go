// Code generated by openapi-generator-go DO NOT EDIT.
//
// Source:
//
//	Title: Test
//	Version: 0.1.0
package generatortest

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

// GetFooQueryParameters is an object.
type GetFooQueryParameters struct {
	// Plus1:
	Plus1 int32 `json:"+1,omitempty" mapstructure:"+1,omitempty"`
	// Minus1:
	Minus1 int32 `json:"-1,omitempty" mapstructure:"-1,omitempty"`
	// Id:
	Id string `json:"id" mapstructure:"id"`
	// Page: The current set of paged results to display, based on a 1-based array index
	Page int32 `json:"page,omitempty" mapstructure:"page,omitempty"`
	// Param1:
	Param1 string `json:"param1,omitempty" mapstructure:"param1,omitempty"`
	// Param2:
	Param2 int32 `json:"param2,omitempty" mapstructure:"param2,omitempty"`
	// Param3:
	Param3 []string `json:"param3,omitempty" mapstructure:"param3,omitempty"`
	// Param4: A referenced status
	Param4 ReferencedStatus `json:"param4,omitempty" mapstructure:"param4,omitempty"`
}

// Validate implements basic validation for this model
func (m GetFooQueryParameters) Validate() error {
	return validation.Errors{
		"id": validation.Validate(
			m.Id, validation.Required, is.UUID,
		),
		"page": validation.Validate(
			m.Page, validation.Min(int32(1)),
		),
		"param1": validation.Validate(
			m.Param1, is.UUID,
		),
		"param2": validation.Validate(
			m.Param2, validation.Min(int32(0)), validation.Max(int32(10)),
		),
		"param3": validation.Validate(
			m.Param3,
		),
		"param4": validation.Validate(
			m.Param4,
		),
	}.Filter()
}

// GetPlus1 returns the Plus1 property
func (m GetFooQueryParameters) GetPlus1() int32 {
	return m.Plus1
}

// SetPlus1 sets the Plus1 property
func (m *GetFooQueryParameters) SetPlus1(val int32) {
	m.Plus1 = val
}

// GetMinus1 returns the Minus1 property
func (m GetFooQueryParameters) GetMinus1() int32 {
	return m.Minus1
}

// SetMinus1 sets the Minus1 property
func (m *GetFooQueryParameters) SetMinus1(val int32) {
	m.Minus1 = val
}

// GetId returns the Id property
func (m GetFooQueryParameters) GetId() string {
	return m.Id
}

// SetId sets the Id property
func (m *GetFooQueryParameters) SetId(val string) {
	m.Id = val
}

// GetPage returns the Page property
func (m GetFooQueryParameters) GetPage() int32 {
	return m.Page
}

// SetPage sets the Page property
func (m *GetFooQueryParameters) SetPage(val int32) {
	m.Page = val
}

// GetParam1 returns the Param1 property
func (m GetFooQueryParameters) GetParam1() string {
	return m.Param1
}

// SetParam1 sets the Param1 property
func (m *GetFooQueryParameters) SetParam1(val string) {
	m.Param1 = val
}

// GetParam2 returns the Param2 property
func (m GetFooQueryParameters) GetParam2() int32 {
	return m.Param2
}

// SetParam2 sets the Param2 property
func (m *GetFooQueryParameters) SetParam2(val int32) {
	m.Param2 = val
}

// GetParam3 returns the Param3 property
func (m GetFooQueryParameters) GetParam3() []string {
	return m.Param3
}

// SetParam3 sets the Param3 property
func (m *GetFooQueryParameters) SetParam3(val []string) {
	m.Param3 = val
}

// GetParam4 returns the Param4 property
func (m GetFooQueryParameters) GetParam4() ReferencedStatus {
	return m.Param4
}

// SetParam4 sets the Param4 property
func (m *GetFooQueryParameters) SetParam4(val ReferencedStatus) {
	m.Param4 = val
}
