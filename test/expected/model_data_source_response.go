// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

import (
	"time"
)

// DataSourceResponse is an object.
type DataSourceResponse struct {
	// CreatedAt:
	CreatedAt time.Time `json:"createdAt"`
	// Description:
	Description string `json:"description"`
	// DisplayName:
	DisplayName string `json:"displayName,omitempty"`
	// Documentation: Long form markdown field that can be used to document the resource
	Documentation string `json:"documentation"`
	// ExternalId:
	ExternalId *string `json:"externalId"`
	// Icon: Name of an icon that is assigned to the data source
	Icon string `json:"icon"`
	// Id:
	Id string `json:"id"`
	// IsOrphaned: Indicates that this datasource couldn't be synced from an external system
	IsOrphaned bool `json:"isOrphaned"`
	// IsPublic:
	IsPublic bool `json:"isPublic"`
	// LastUpdateErrorMessage: contains an error message indicating why this data source schema could not be updated
	LastUpdateErrorMessage string `json:"lastUpdateErrorMessage,omitempty"`
	// Name:
	Name string `json:"name"`
	// OrphanedAt:
	OrphanedAt *time.Time `json:"orphanedAt"`
	// ProjectId:
	ProjectId string `json:"projectId"`
	// Properties:
	Properties map[string]interface{} `json:"properties"`
	// ResourceActions: List of actions allowed to perform on the data source
	ResourceActions DataSourceActions `json:"resourceActions"`
	// Stats: Various statistics about the data source
	Stats DataSourceStatistics `json:"stats,omitempty"`
	// SyncTask:
	SyncTask *DataSourceSyncTask `json:"syncTask"`
	// Technology: Type of a data source
	Technology DataSourceTechnology `json:"technology,omitempty"`
	// Type: Type of a data source
	Type DataSourceType `json:"type"`
	// UpdatedAt:
	UpdatedAt time.Time `json:"updatedAt"`
}

// GetCreatedAt returns the CreatedAt property
func (m DataSourceResponse) GetCreatedAt() time.Time {
	return m.CreatedAt
}

// SetCreatedAt sets the CreatedAt property
func (m DataSourceResponse) SetCreatedAt(val time.Time) {
	m.CreatedAt = val
}

// GetDescription returns the Description property
func (m DataSourceResponse) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m DataSourceResponse) SetDescription(val string) {
	m.Description = val
}

// GetDisplayName returns the DisplayName property
func (m DataSourceResponse) GetDisplayName() string {
	return m.DisplayName
}

// SetDisplayName sets the DisplayName property
func (m DataSourceResponse) SetDisplayName(val string) {
	m.DisplayName = val
}

// GetDocumentation returns the Documentation property
func (m DataSourceResponse) GetDocumentation() string {
	return m.Documentation
}

// SetDocumentation sets the Documentation property
func (m DataSourceResponse) SetDocumentation(val string) {
	m.Documentation = val
}

// GetExternalId returns the ExternalId property
func (m DataSourceResponse) GetExternalId() *string {
	return m.ExternalId
}

// SetExternalId sets the ExternalId property
func (m DataSourceResponse) SetExternalId(val *string) {
	m.ExternalId = val
}

// GetIcon returns the Icon property
func (m DataSourceResponse) GetIcon() string {
	return m.Icon
}

// SetIcon sets the Icon property
func (m DataSourceResponse) SetIcon(val string) {
	m.Icon = val
}

// GetId returns the Id property
func (m DataSourceResponse) GetId() string {
	return m.Id
}

// SetId sets the Id property
func (m DataSourceResponse) SetId(val string) {
	m.Id = val
}

// GetIsOrphaned returns the IsOrphaned property
func (m DataSourceResponse) GetIsOrphaned() bool {
	return m.IsOrphaned
}

// SetIsOrphaned sets the IsOrphaned property
func (m DataSourceResponse) SetIsOrphaned(val bool) {
	m.IsOrphaned = val
}

// GetIsPublic returns the IsPublic property
func (m DataSourceResponse) GetIsPublic() bool {
	return m.IsPublic
}

// SetIsPublic sets the IsPublic property
func (m DataSourceResponse) SetIsPublic(val bool) {
	m.IsPublic = val
}

// GetLastUpdateErrorMessage returns the LastUpdateErrorMessage property
func (m DataSourceResponse) GetLastUpdateErrorMessage() string {
	return m.LastUpdateErrorMessage
}

// SetLastUpdateErrorMessage sets the LastUpdateErrorMessage property
func (m DataSourceResponse) SetLastUpdateErrorMessage(val string) {
	m.LastUpdateErrorMessage = val
}

// GetName returns the Name property
func (m DataSourceResponse) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m DataSourceResponse) SetName(val string) {
	m.Name = val
}

// GetOrphanedAt returns the OrphanedAt property
func (m DataSourceResponse) GetOrphanedAt() *time.Time {
	return m.OrphanedAt
}

// SetOrphanedAt sets the OrphanedAt property
func (m DataSourceResponse) SetOrphanedAt(val *time.Time) {
	m.OrphanedAt = val
}

// GetProjectId returns the ProjectId property
func (m DataSourceResponse) GetProjectId() string {
	return m.ProjectId
}

// SetProjectId sets the ProjectId property
func (m DataSourceResponse) SetProjectId(val string) {
	m.ProjectId = val
}

// GetProperties returns the Properties property
func (m DataSourceResponse) GetProperties() map[string]interface{} {
	return m.Properties
}

// SetProperties sets the Properties property
func (m DataSourceResponse) SetProperties(val map[string]interface{}) {
	m.Properties = val
}

// GetResourceActions returns the ResourceActions property
func (m DataSourceResponse) GetResourceActions() DataSourceActions {
	return m.ResourceActions
}

// SetResourceActions sets the ResourceActions property
func (m DataSourceResponse) SetResourceActions(val DataSourceActions) {
	m.ResourceActions = val
}

// GetStats returns the Stats property
func (m DataSourceResponse) GetStats() DataSourceStatistics {
	return m.Stats
}

// SetStats sets the Stats property
func (m DataSourceResponse) SetStats(val DataSourceStatistics) {
	m.Stats = val
}

// GetSyncTask returns the SyncTask property
func (m DataSourceResponse) GetSyncTask() *DataSourceSyncTask {
	return m.SyncTask
}

// SetSyncTask sets the SyncTask property
func (m DataSourceResponse) SetSyncTask(val *DataSourceSyncTask) {
	m.SyncTask = val
}

// GetTechnology returns the Technology property
func (m DataSourceResponse) GetTechnology() DataSourceTechnology {
	return m.Technology
}

// SetTechnology sets the Technology property
func (m DataSourceResponse) SetTechnology(val DataSourceTechnology) {
	m.Technology = val
}

// GetType returns the Type property
func (m DataSourceResponse) GetType() DataSourceType {
	return m.Type
}

// SetType sets the Type property
func (m DataSourceResponse) SetType(val DataSourceType) {
	m.Type = val
}

// GetUpdatedAt returns the UpdatedAt property
func (m DataSourceResponse) GetUpdatedAt() time.Time {
	return m.UpdatedAt
}

// SetUpdatedAt sets the UpdatedAt property
func (m DataSourceResponse) SetUpdatedAt(val time.Time) {
	m.UpdatedAt = val
}
