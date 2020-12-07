// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

import (
	"time"
)

// TableSearchResponse is an object.
type TableSearchResponse struct {
	// CreatedAt:
	CreatedAt time.Time `json:"createdAt"`
	// DataSourceId:
	DataSourceId string `json:"dataSourceId"`
	// Description:
	Description string `json:"description"`
	// Highlight: contains the matching parts of the description
	Highlight string `json:"highlight,omitempty"`
	// Id:
	Id string `json:"id"`
	// LastUpdateErrorMessage: contains an error message indicating why this table schema could not be updated
	LastUpdateErrorMessage string `json:"lastUpdateErrorMessage,omitempty"`
	// Name:
	Name string `json:"name"`
	// Properties:
	Properties map[string]interface{} `json:"properties"`
	// Sql:
	Sql string `json:"sql,omitempty"`
	// Type:
	Type TableType `json:"type"`
	// UpdatedAt:
	UpdatedAt time.Time `json:"updatedAt"`
}

// GetCreatedAt returns the CreatedAt property
func (m TableSearchResponse) GetCreatedAt() time.Time {
	return m.CreatedAt
}

// SetCreatedAt sets the CreatedAt property
func (m TableSearchResponse) SetCreatedAt(val time.Time) {
	m.CreatedAt = val
}

// GetDataSourceId returns the DataSourceId property
func (m TableSearchResponse) GetDataSourceId() string {
	return m.DataSourceId
}

// SetDataSourceId sets the DataSourceId property
func (m TableSearchResponse) SetDataSourceId(val string) {
	m.DataSourceId = val
}

// GetDescription returns the Description property
func (m TableSearchResponse) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m TableSearchResponse) SetDescription(val string) {
	m.Description = val
}

// GetHighlight returns the Highlight property
func (m TableSearchResponse) GetHighlight() string {
	return m.Highlight
}

// SetHighlight sets the Highlight property
func (m TableSearchResponse) SetHighlight(val string) {
	m.Highlight = val
}

// GetId returns the Id property
func (m TableSearchResponse) GetId() string {
	return m.Id
}

// SetId sets the Id property
func (m TableSearchResponse) SetId(val string) {
	m.Id = val
}

// GetLastUpdateErrorMessage returns the LastUpdateErrorMessage property
func (m TableSearchResponse) GetLastUpdateErrorMessage() string {
	return m.LastUpdateErrorMessage
}

// SetLastUpdateErrorMessage sets the LastUpdateErrorMessage property
func (m TableSearchResponse) SetLastUpdateErrorMessage(val string) {
	m.LastUpdateErrorMessage = val
}

// GetName returns the Name property
func (m TableSearchResponse) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m TableSearchResponse) SetName(val string) {
	m.Name = val
}

// GetProperties returns the Properties property
func (m TableSearchResponse) GetProperties() map[string]interface{} {
	return m.Properties
}

// SetProperties sets the Properties property
func (m TableSearchResponse) SetProperties(val map[string]interface{}) {
	m.Properties = val
}

// GetSql returns the Sql property
func (m TableSearchResponse) GetSql() string {
	return m.Sql
}

// SetSql sets the Sql property
func (m TableSearchResponse) SetSql(val string) {
	m.Sql = val
}

// GetType returns the Type property
func (m TableSearchResponse) GetType() TableType {
	return m.Type
}

// SetType sets the Type property
func (m TableSearchResponse) SetType(val TableType) {
	m.Type = val
}

// GetUpdatedAt returns the UpdatedAt property
func (m TableSearchResponse) GetUpdatedAt() time.Time {
	return m.UpdatedAt
}

// SetUpdatedAt sets the UpdatedAt property
func (m TableSearchResponse) SetUpdatedAt(val time.Time) {
	m.UpdatedAt = val
}
