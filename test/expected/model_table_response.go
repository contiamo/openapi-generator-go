// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

import (
	"time"
)

// TableResponse is an object.
type TableResponse struct {
	// CreatedAt:
	CreatedAt time.Time `json:"createdAt"`
	// CurrentColumnSetId:
	CurrentColumnSetId string `json:"currentColumnSetId"`
	// DataSourceId:
	DataSourceId string `json:"dataSourceId"`
	// Description:
	Description string `json:"description"`
	// Documentation: Long form markdown field that can be used to document the resource
	Documentation string `json:"documentation"`
	// Id:
	Id string `json:"id"`
	// IsOrphaned: Indicates that this table couldn't be synced from an external system
	IsOrphaned bool `json:"isOrphaned"`
	// LastUpdateErrorMessage: contains an error message indicating why this table schema could not be updated
	LastUpdateErrorMessage string `json:"lastUpdateErrorMessage,omitempty"`
	// Name:
	Name string `json:"name"`
	// OrphanedAt:
	OrphanedAt *time.Time `json:"orphanedAt"`
	// Properties:
	Properties map[string]interface{} `json:"properties"`
	// Sql:
	Sql string `json:"sql,omitempty"`
	// Stats: Various statistics about the table
	Stats TableStatistics `json:"stats"`
	// Type:
	Type TableType `json:"type"`
	// UpdatedAt:
	UpdatedAt time.Time `json:"updatedAt"`
}

// GetCreatedAt returns the CreatedAt property
func (m TableResponse) GetCreatedAt() time.Time {
	return m.CreatedAt
}

// SetCreatedAt sets the CreatedAt property
func (m TableResponse) SetCreatedAt(val time.Time) {
	m.CreatedAt = val
}

// GetCurrentColumnSetId returns the CurrentColumnSetId property
func (m TableResponse) GetCurrentColumnSetId() string {
	return m.CurrentColumnSetId
}

// SetCurrentColumnSetId sets the CurrentColumnSetId property
func (m TableResponse) SetCurrentColumnSetId(val string) {
	m.CurrentColumnSetId = val
}

// GetDataSourceId returns the DataSourceId property
func (m TableResponse) GetDataSourceId() string {
	return m.DataSourceId
}

// SetDataSourceId sets the DataSourceId property
func (m TableResponse) SetDataSourceId(val string) {
	m.DataSourceId = val
}

// GetDescription returns the Description property
func (m TableResponse) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m TableResponse) SetDescription(val string) {
	m.Description = val
}

// GetDocumentation returns the Documentation property
func (m TableResponse) GetDocumentation() string {
	return m.Documentation
}

// SetDocumentation sets the Documentation property
func (m TableResponse) SetDocumentation(val string) {
	m.Documentation = val
}

// GetId returns the Id property
func (m TableResponse) GetId() string {
	return m.Id
}

// SetId sets the Id property
func (m TableResponse) SetId(val string) {
	m.Id = val
}

// GetIsOrphaned returns the IsOrphaned property
func (m TableResponse) GetIsOrphaned() bool {
	return m.IsOrphaned
}

// SetIsOrphaned sets the IsOrphaned property
func (m TableResponse) SetIsOrphaned(val bool) {
	m.IsOrphaned = val
}

// GetLastUpdateErrorMessage returns the LastUpdateErrorMessage property
func (m TableResponse) GetLastUpdateErrorMessage() string {
	return m.LastUpdateErrorMessage
}

// SetLastUpdateErrorMessage sets the LastUpdateErrorMessage property
func (m TableResponse) SetLastUpdateErrorMessage(val string) {
	m.LastUpdateErrorMessage = val
}

// GetName returns the Name property
func (m TableResponse) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m TableResponse) SetName(val string) {
	m.Name = val
}

// GetOrphanedAt returns the OrphanedAt property
func (m TableResponse) GetOrphanedAt() *time.Time {
	return m.OrphanedAt
}

// SetOrphanedAt sets the OrphanedAt property
func (m TableResponse) SetOrphanedAt(val *time.Time) {
	m.OrphanedAt = val
}

// GetProperties returns the Properties property
func (m TableResponse) GetProperties() map[string]interface{} {
	return m.Properties
}

// SetProperties sets the Properties property
func (m TableResponse) SetProperties(val map[string]interface{}) {
	m.Properties = val
}

// GetSql returns the Sql property
func (m TableResponse) GetSql() string {
	return m.Sql
}

// SetSql sets the Sql property
func (m TableResponse) SetSql(val string) {
	m.Sql = val
}

// GetStats returns the Stats property
func (m TableResponse) GetStats() TableStatistics {
	return m.Stats
}

// SetStats sets the Stats property
func (m TableResponse) SetStats(val TableStatistics) {
	m.Stats = val
}

// GetType returns the Type property
func (m TableResponse) GetType() TableType {
	return m.Type
}

// SetType sets the Type property
func (m TableResponse) SetType(val TableType) {
	m.Type = val
}

// GetUpdatedAt returns the UpdatedAt property
func (m TableResponse) GetUpdatedAt() time.Time {
	return m.UpdatedAt
}

// SetUpdatedAt sets the UpdatedAt property
func (m TableResponse) SetUpdatedAt(val time.Time) {
	m.UpdatedAt = val
}
