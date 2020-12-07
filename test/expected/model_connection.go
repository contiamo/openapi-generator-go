// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

import (
	"time"
)

// Connection is an object.
type Connection struct {
	// CreatedAt:
	CreatedAt time.Time `json:"createdAt"`
	// DataSourceId:
	DataSourceId string `json:"dataSourceId"`
	// Description:
	Description string `json:"description"`
	// Id:
	Id string `json:"id"`
	// LastHeartbeatError:
	LastHeartbeatError string `json:"lastHeartbeatError,omitempty"`
	// Name:
	Name string `json:"name"`
	// Properties:
	Properties ConnectionProperties `json:"properties"`
	// Technology: Type of a data source
	Technology DataSourceTechnology `json:"technology"`
	// Type:
	Type ConnectionType `json:"type"`
	// UpdatedAt:
	UpdatedAt time.Time `json:"updatedAt"`
}

// GetCreatedAt returns the CreatedAt property
func (m Connection) GetCreatedAt() time.Time {
	return m.CreatedAt
}

// SetCreatedAt sets the CreatedAt property
func (m Connection) SetCreatedAt(val time.Time) {
	m.CreatedAt = val
}

// GetDataSourceId returns the DataSourceId property
func (m Connection) GetDataSourceId() string {
	return m.DataSourceId
}

// SetDataSourceId sets the DataSourceId property
func (m Connection) SetDataSourceId(val string) {
	m.DataSourceId = val
}

// GetDescription returns the Description property
func (m Connection) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m Connection) SetDescription(val string) {
	m.Description = val
}

// GetId returns the Id property
func (m Connection) GetId() string {
	return m.Id
}

// SetId sets the Id property
func (m Connection) SetId(val string) {
	m.Id = val
}

// GetLastHeartbeatError returns the LastHeartbeatError property
func (m Connection) GetLastHeartbeatError() string {
	return m.LastHeartbeatError
}

// SetLastHeartbeatError sets the LastHeartbeatError property
func (m Connection) SetLastHeartbeatError(val string) {
	m.LastHeartbeatError = val
}

// GetName returns the Name property
func (m Connection) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m Connection) SetName(val string) {
	m.Name = val
}

// GetProperties returns the Properties property
func (m Connection) GetProperties() ConnectionProperties {
	return m.Properties
}

// SetProperties sets the Properties property
func (m Connection) SetProperties(val ConnectionProperties) {
	m.Properties = val
}

// GetTechnology returns the Technology property
func (m Connection) GetTechnology() DataSourceTechnology {
	return m.Technology
}

// SetTechnology sets the Technology property
func (m Connection) SetTechnology(val DataSourceTechnology) {
	m.Technology = val
}

// GetType returns the Type property
func (m Connection) GetType() ConnectionType {
	return m.Type
}

// SetType sets the Type property
func (m Connection) SetType(val ConnectionType) {
	m.Type = val
}

// GetUpdatedAt returns the UpdatedAt property
func (m Connection) GetUpdatedAt() time.Time {
	return m.UpdatedAt
}

// SetUpdatedAt sets the UpdatedAt property
func (m Connection) SetUpdatedAt(val time.Time) {
	m.UpdatedAt = val
}
