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
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"regexp"
	"time"
)

// connectionNamePattern is the validation pattern for Connection.Name
var connectionNamePattern = regexp.MustCompile(`^([a-zA-z_]+[a-z0-9_]*){2}$`)

// Connection is an object.
type Connection struct {
	// CreatedAt:
	CreatedAt time.Time `json:"createdAt" mapstructure:"createdAt"`
	// Id:
	Id string `json:"id" mapstructure:"id"`
	// Name:
	Name string `json:"name" mapstructure:"name"`
	// Properties:
	Properties ConnectionProperties `json:"properties" mapstructure:"properties"`
	// Technology: The connection technology is either the technology value of the related data source or the integration type
	Technology ConnectionTechnology `json:"technology" mapstructure:"technology"`
	// UpdatedAt:
	UpdatedAt time.Time `json:"updatedAt" mapstructure:"updatedAt"`
}

// NewConnection instantiates a new Connection with default values overriding them as follows:
// 1. Default values specified in the Connection schema
// 2. Default values specified per Connection property
func NewConnection() *Connection {
	m := &Connection{}

	return m
}

// UnmarshalJSON implements the json.Unmarshaler interface for Connection. It set the default values for the Connection type
func (m *Connection) UnmarshalJSON(data []byte) error {
	// Set default values
	*m = *NewConnection()

	// Unmarshal using an alias to avoid an infinite loop
	type alias Connection
	err := json.Unmarshal(data, (*alias)(m))
	if err != nil {
		return err
	}
	return nil
}

// Validate implements basic validation for this model
func (m Connection) Validate() error {
	errors := validation.Errors{
		"createdAt": validation.Validate(
			m.CreatedAt, validation.Required, validation.Date(time.RFC3339),
		),
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
			m.Technology, validation.Required,
		),
		"updatedAt": validation.Validate(
			m.UpdatedAt, validation.Required, validation.Date(time.RFC3339),
		),
	}
	return errors.Filter()
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
