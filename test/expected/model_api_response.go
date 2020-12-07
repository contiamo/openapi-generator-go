// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

import (
	"time"
)

// APIResponse is an object.
type APIResponse struct {
	// CreatedAt:
	CreatedAt time.Time `json:"createdAt"`
	// Description:
	Description string `json:"description"`
	// DisplayName:
	DisplayName string `json:"displayName"`
	// Documentation: Long form markdown field that can be used to document the resource
	Documentation string `json:"documentation"`
	// ExternalId:
	ExternalId *string `json:"externalId"`
	// Id:
	Id string `json:"id"`
	// IsOrphaned: Indicates that this API couldn't be synced from an external system
	IsOrphaned bool `json:"isOrphaned"`
	// IsPublic:
	IsPublic bool `json:"isPublic"`
	// Name:
	Name string `json:"name"`
	// OrphanedAt:
	OrphanedAt *time.Time `json:"orphanedAt"`
	// ProjectId:
	ProjectId string `json:"projectId"`
	// Properties:
	Properties map[string]interface{} `json:"properties"`
	// ResourceActions: List of actions allowed to perform on the API
	ResourceActions APIActions `json:"resourceActions"`
	// Stats: Various statistics about the API
	Stats APIStatistics `json:"stats"`
	// Type:
	Type APIType `json:"type"`
	// UpdatedAt:
	UpdatedAt time.Time `json:"updatedAt"`
}

// GetCreatedAt returns the CreatedAt property
func (m APIResponse) GetCreatedAt() time.Time {
	return m.CreatedAt
}

// SetCreatedAt sets the CreatedAt property
func (m APIResponse) SetCreatedAt(val time.Time) {
	m.CreatedAt = val
}

// GetDescription returns the Description property
func (m APIResponse) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m APIResponse) SetDescription(val string) {
	m.Description = val
}

// GetDisplayName returns the DisplayName property
func (m APIResponse) GetDisplayName() string {
	return m.DisplayName
}

// SetDisplayName sets the DisplayName property
func (m APIResponse) SetDisplayName(val string) {
	m.DisplayName = val
}

// GetDocumentation returns the Documentation property
func (m APIResponse) GetDocumentation() string {
	return m.Documentation
}

// SetDocumentation sets the Documentation property
func (m APIResponse) SetDocumentation(val string) {
	m.Documentation = val
}

// GetExternalId returns the ExternalId property
func (m APIResponse) GetExternalId() *string {
	return m.ExternalId
}

// SetExternalId sets the ExternalId property
func (m APIResponse) SetExternalId(val *string) {
	m.ExternalId = val
}

// GetId returns the Id property
func (m APIResponse) GetId() string {
	return m.Id
}

// SetId sets the Id property
func (m APIResponse) SetId(val string) {
	m.Id = val
}

// GetIsOrphaned returns the IsOrphaned property
func (m APIResponse) GetIsOrphaned() bool {
	return m.IsOrphaned
}

// SetIsOrphaned sets the IsOrphaned property
func (m APIResponse) SetIsOrphaned(val bool) {
	m.IsOrphaned = val
}

// GetIsPublic returns the IsPublic property
func (m APIResponse) GetIsPublic() bool {
	return m.IsPublic
}

// SetIsPublic sets the IsPublic property
func (m APIResponse) SetIsPublic(val bool) {
	m.IsPublic = val
}

// GetName returns the Name property
func (m APIResponse) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m APIResponse) SetName(val string) {
	m.Name = val
}

// GetOrphanedAt returns the OrphanedAt property
func (m APIResponse) GetOrphanedAt() *time.Time {
	return m.OrphanedAt
}

// SetOrphanedAt sets the OrphanedAt property
func (m APIResponse) SetOrphanedAt(val *time.Time) {
	m.OrphanedAt = val
}

// GetProjectId returns the ProjectId property
func (m APIResponse) GetProjectId() string {
	return m.ProjectId
}

// SetProjectId sets the ProjectId property
func (m APIResponse) SetProjectId(val string) {
	m.ProjectId = val
}

// GetProperties returns the Properties property
func (m APIResponse) GetProperties() map[string]interface{} {
	return m.Properties
}

// SetProperties sets the Properties property
func (m APIResponse) SetProperties(val map[string]interface{}) {
	m.Properties = val
}

// GetResourceActions returns the ResourceActions property
func (m APIResponse) GetResourceActions() APIActions {
	return m.ResourceActions
}

// SetResourceActions sets the ResourceActions property
func (m APIResponse) SetResourceActions(val APIActions) {
	m.ResourceActions = val
}

// GetStats returns the Stats property
func (m APIResponse) GetStats() APIStatistics {
	return m.Stats
}

// SetStats sets the Stats property
func (m APIResponse) SetStats(val APIStatistics) {
	m.Stats = val
}

// GetType returns the Type property
func (m APIResponse) GetType() APIType {
	return m.Type
}

// SetType sets the Type property
func (m APIResponse) SetType(val APIType) {
	m.Type = val
}

// GetUpdatedAt returns the UpdatedAt property
func (m APIResponse) GetUpdatedAt() time.Time {
	return m.UpdatedAt
}

// SetUpdatedAt sets the UpdatedAt property
func (m APIResponse) SetUpdatedAt(val time.Time) {
	m.UpdatedAt = val
}
