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

// ColumnMetadata is an object. Metadata for single column
type ColumnMetadata struct {
	// Comment: Column description
	Comment string `json:"comment,omitempty" mapstructure:"comment,omitempty"`
	// Name: Column name
	Name string `json:"name" mapstructure:"name"`
	// Type: Type metadata
	Type ColumnTypeMetadata `json:"type" mapstructure:"type"`
}

// Validate implements basic validation for this model
func (m ColumnMetadata) Validate() error {
	return validation.Errors{
		"type": validation.Validate(
			m.Type, validation.NotNil,
		),
	}.Filter()
}

// GetComment returns the Comment property
func (m ColumnMetadata) GetComment() string {
	return m.Comment
}

// SetComment sets the Comment property
func (m *ColumnMetadata) SetComment(val string) {
	m.Comment = val
}

// GetName returns the Name property
func (m ColumnMetadata) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *ColumnMetadata) SetName(val string) {
	m.Name = val
}

// GetType returns the Type property
func (m ColumnMetadata) GetType() ColumnTypeMetadata {
	return m.Type
}

// SetType sets the Type property
func (m *ColumnMetadata) SetType(val ColumnTypeMetadata) {
	m.Type = val
}
