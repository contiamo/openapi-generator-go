// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

import (
	"time"
)

// ApplicationResponse is an object.
type ApplicationResponse struct {
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
	// IsOrphaned: Indicates that this application couldn't be synced from an external system
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
	// ResourceActions: List of actions allowed to perform on the application
	ResourceActions ApplicationActions `json:"resourceActions"`
	// Stats: Various statistics about the application
	Stats ApplicationStatistics `json:"stats"`
	// Type:
	Type ApplicationType `json:"type"`
	// UpdatedAt:
	UpdatedAt time.Time `json:"updatedAt"`
}

// GetCreatedAt returns the CreatedAt property
func (m ApplicationResponse) GetCreatedAt() time.Time {
	return m.CreatedAt
}

// SetCreatedAt sets the CreatedAt property
func (m ApplicationResponse) SetCreatedAt(val time.Time) {
	m.CreatedAt = val
}

// GetDescription returns the Description property
func (m ApplicationResponse) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m ApplicationResponse) SetDescription(val string) {
	m.Description = val
}

// GetDisplayName returns the DisplayName property
func (m ApplicationResponse) GetDisplayName() string {
	return m.DisplayName
}

// SetDisplayName sets the DisplayName property
func (m ApplicationResponse) SetDisplayName(val string) {
	m.DisplayName = val
}

// GetDocumentation returns the Documentation property
func (m ApplicationResponse) GetDocumentation() string {
	return m.Documentation
}

// SetDocumentation sets the Documentation property
func (m ApplicationResponse) SetDocumentation(val string) {
	m.Documentation = val
}

// GetExternalId returns the ExternalId property
func (m ApplicationResponse) GetExternalId() *string {
	return m.ExternalId
}

// SetExternalId sets the ExternalId property
func (m ApplicationResponse) SetExternalId(val *string) {
	m.ExternalId = val
}

// GetId returns the Id property
func (m ApplicationResponse) GetId() string {
	return m.Id
}

// SetId sets the Id property
func (m ApplicationResponse) SetId(val string) {
	m.Id = val
}

// GetIsOrphaned returns the IsOrphaned property
func (m ApplicationResponse) GetIsOrphaned() bool {
	return m.IsOrphaned
}

// SetIsOrphaned sets the IsOrphaned property
func (m ApplicationResponse) SetIsOrphaned(val bool) {
	m.IsOrphaned = val
}

// GetIsPublic returns the IsPublic property
func (m ApplicationResponse) GetIsPublic() bool {
	return m.IsPublic
}

// SetIsPublic sets the IsPublic property
func (m ApplicationResponse) SetIsPublic(val bool) {
	m.IsPublic = val
}

// GetName returns the Name property
func (m ApplicationResponse) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m ApplicationResponse) SetName(val string) {
	m.Name = val
}

// GetOrphanedAt returns the OrphanedAt property
func (m ApplicationResponse) GetOrphanedAt() *time.Time {
	return m.OrphanedAt
}

// SetOrphanedAt sets the OrphanedAt property
func (m ApplicationResponse) SetOrphanedAt(val *time.Time) {
	m.OrphanedAt = val
}

// GetProjectId returns the ProjectId property
func (m ApplicationResponse) GetProjectId() string {
	return m.ProjectId
}

// SetProjectId sets the ProjectId property
func (m ApplicationResponse) SetProjectId(val string) {
	m.ProjectId = val
}

// GetProperties returns the Properties property
func (m ApplicationResponse) GetProperties() map[string]interface{} {
	return m.Properties
}

// SetProperties sets the Properties property
func (m ApplicationResponse) SetProperties(val map[string]interface{}) {
	m.Properties = val
}

// GetResourceActions returns the ResourceActions property
func (m ApplicationResponse) GetResourceActions() ApplicationActions {
	return m.ResourceActions
}

// SetResourceActions sets the ResourceActions property
func (m ApplicationResponse) SetResourceActions(val ApplicationActions) {
	m.ResourceActions = val
}

// GetStats returns the Stats property
func (m ApplicationResponse) GetStats() ApplicationStatistics {
	return m.Stats
}

// SetStats sets the Stats property
func (m ApplicationResponse) SetStats(val ApplicationStatistics) {
	m.Stats = val
}

// GetType returns the Type property
func (m ApplicationResponse) GetType() ApplicationType {
	return m.Type
}

// SetType sets the Type property
func (m ApplicationResponse) SetType(val ApplicationType) {
	m.Type = val
}

// GetUpdatedAt returns the UpdatedAt property
func (m ApplicationResponse) GetUpdatedAt() time.Time {
	return m.UpdatedAt
}

// SetUpdatedAt sets the UpdatedAt property
func (m ApplicationResponse) SetUpdatedAt(val time.Time) {
	m.UpdatedAt = val
}
