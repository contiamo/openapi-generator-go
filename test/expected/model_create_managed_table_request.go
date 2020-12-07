// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

import (
	"time"
)

// CreateManagedTableRequest is an object.
type CreateManagedTableRequest struct {
	// DefaultInsertMode: Defines how to apply rows to a table.
	// upsert: insert or update rows from the upload to the table
	// append: insert rows from the upload to the table.
	DefaultInsertMode IngestInsertModeType `json:"defaultInsertMode"`
	// Description:
	Description string `json:"description"`
	// ExpiresAt: Table will be deleted at the specified time
	ExpiresAt *time.Time `json:"expiresAt"`
	// Name:
	Name string `json:"name"`
	// Schema: Schema of a file or of the corresponding managed table which will be created
	Schema ManagedTableSchema `json:"schema"`
}

// GetDefaultInsertMode returns the DefaultInsertMode property
func (m CreateManagedTableRequest) GetDefaultInsertMode() IngestInsertModeType {
	return m.DefaultInsertMode
}

// SetDefaultInsertMode sets the DefaultInsertMode property
func (m CreateManagedTableRequest) SetDefaultInsertMode(val IngestInsertModeType) {
	m.DefaultInsertMode = val
}

// GetDescription returns the Description property
func (m CreateManagedTableRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m CreateManagedTableRequest) SetDescription(val string) {
	m.Description = val
}

// GetExpiresAt returns the ExpiresAt property
func (m CreateManagedTableRequest) GetExpiresAt() *time.Time {
	return m.ExpiresAt
}

// SetExpiresAt sets the ExpiresAt property
func (m CreateManagedTableRequest) SetExpiresAt(val *time.Time) {
	m.ExpiresAt = val
}

// GetName returns the Name property
func (m CreateManagedTableRequest) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m CreateManagedTableRequest) SetName(val string) {
	m.Name = val
}

// GetSchema returns the Schema property
func (m CreateManagedTableRequest) GetSchema() ManagedTableSchema {
	return m.Schema
}

// SetSchema sets the Schema property
func (m CreateManagedTableRequest) SetSchema(val ManagedTableSchema) {
	m.Schema = val
}
