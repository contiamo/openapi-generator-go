// This file is auto-generated, DO NOT EDIT.
//
// Source:
//
//	Title: Test
//	Version: 0.1.0
package generatortest

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"time"

	"github.com/go-ozzo/ozzo-validation/v4/is"

	"regexp"
)

// connectionNamePattern is the validation pattern for Connection.Name
var connectionNamePattern = regexp.MustCompile(`^([a-zA-z_]+[a-z0-9_]*){2}$`)

// Connection is an object.
type Connection struct {
	// CreatedAt:
	CreatedAt time.Time `json:"createdAt"`
	// Id:
	Id string `json:"id"`
	// Name:
	Name string `json:"name"`
	// Properties:
	Properties ConnectionProperties `json:"properties"`
	// Technology: The connection technology is either the technology value of the related data source or the integration type
	Technology ConnectionTechnology `json:"technology"`
	// UpdatedAt:
	UpdatedAt time.Time `json:"updatedAt"`
}

// Validate implements basic validation for this model
func (m Connection) Validate() error {
	return validation.Errors{
		"id": validation.Validate(
			m.Id, validation.Required, is.UUID,
		),
		"name": validation.Validate(
			m.Name, validation.Required, validation.Length(2, 255), validation.Match(connectionNamePattern),
		),
		"properties": validation.Validate(
			m.Properties, validation.NotNil,
		),
		"technology": validation.Validate(
			m.Technology, validation.NotNil,
		),
	}.Filter()
}

// GetCreatedAt returns the CreatedAt property
func (m Connection) GetCreatedAt() time.Time {
	return m.CreatedAt
}

// SetCreatedAt sets the CreatedAt property
func (m *Connection) SetCreatedAt(val time.Time) {
	m.CreatedAt = val
}

// GetId returns the Id property
func (m Connection) GetId() string {
	return m.Id
}

// SetId sets the Id property
func (m *Connection) SetId(val string) {
	m.Id = val
}

// GetName returns the Name property
func (m Connection) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *Connection) SetName(val string) {
	m.Name = val
}

// GetProperties returns the Properties property
func (m Connection) GetProperties() ConnectionProperties {
	return m.Properties
}

// SetProperties sets the Properties property
func (m *Connection) SetProperties(val ConnectionProperties) {
	m.Properties = val
}

// GetTechnology returns the Technology property
func (m Connection) GetTechnology() ConnectionTechnology {
	return m.Technology
}

// SetTechnology sets the Technology property
func (m *Connection) SetTechnology(val ConnectionTechnology) {
	m.Technology = val
}

// GetUpdatedAt returns the UpdatedAt property
func (m Connection) GetUpdatedAt() time.Time {
	return m.UpdatedAt
}

// SetUpdatedAt sets the UpdatedAt property
func (m *Connection) SetUpdatedAt(val time.Time) {
	m.UpdatedAt = val
}
