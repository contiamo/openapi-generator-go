// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

import (
	"time"
)

// ModelResponse is an object.
type ModelResponse struct {
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
	// IsOrphaned: Indicates that this model couldn't be synced from an external system
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
	// ResourceActions: List of actions allowed to perform on the model
	ResourceActions ModelActions `json:"resourceActions"`
	// Stats: Various statistics about the model
	Stats ModelStatistics `json:"stats"`
	// Type:
	Type ModelType `json:"type"`
	// UpdatedAt:
	UpdatedAt time.Time `json:"updatedAt"`
}

// GetCreatedAt returns the CreatedAt property
func (m ModelResponse) GetCreatedAt() time.Time {
	return m.CreatedAt
}

// SetCreatedAt sets the CreatedAt property
func (m ModelResponse) SetCreatedAt(val time.Time) {
	m.CreatedAt = val
}

// GetDescription returns the Description property
func (m ModelResponse) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m ModelResponse) SetDescription(val string) {
	m.Description = val
}

// GetDisplayName returns the DisplayName property
func (m ModelResponse) GetDisplayName() string {
	return m.DisplayName
}

// SetDisplayName sets the DisplayName property
func (m ModelResponse) SetDisplayName(val string) {
	m.DisplayName = val
}

// GetDocumentation returns the Documentation property
func (m ModelResponse) GetDocumentation() string {
	return m.Documentation
}

// SetDocumentation sets the Documentation property
func (m ModelResponse) SetDocumentation(val string) {
	m.Documentation = val
}

// GetExternalId returns the ExternalId property
func (m ModelResponse) GetExternalId() *string {
	return m.ExternalId
}

// SetExternalId sets the ExternalId property
func (m ModelResponse) SetExternalId(val *string) {
	m.ExternalId = val
}

// GetId returns the Id property
func (m ModelResponse) GetId() string {
	return m.Id
}

// SetId sets the Id property
func (m ModelResponse) SetId(val string) {
	m.Id = val
}

// GetIsOrphaned returns the IsOrphaned property
func (m ModelResponse) GetIsOrphaned() bool {
	return m.IsOrphaned
}

// SetIsOrphaned sets the IsOrphaned property
func (m ModelResponse) SetIsOrphaned(val bool) {
	m.IsOrphaned = val
}

// GetIsPublic returns the IsPublic property
func (m ModelResponse) GetIsPublic() bool {
	return m.IsPublic
}

// SetIsPublic sets the IsPublic property
func (m ModelResponse) SetIsPublic(val bool) {
	m.IsPublic = val
}

// GetName returns the Name property
func (m ModelResponse) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m ModelResponse) SetName(val string) {
	m.Name = val
}

// GetOrphanedAt returns the OrphanedAt property
func (m ModelResponse) GetOrphanedAt() *time.Time {
	return m.OrphanedAt
}

// SetOrphanedAt sets the OrphanedAt property
func (m ModelResponse) SetOrphanedAt(val *time.Time) {
	m.OrphanedAt = val
}

// GetProjectId returns the ProjectId property
func (m ModelResponse) GetProjectId() string {
	return m.ProjectId
}

// SetProjectId sets the ProjectId property
func (m ModelResponse) SetProjectId(val string) {
	m.ProjectId = val
}

// GetProperties returns the Properties property
func (m ModelResponse) GetProperties() map[string]interface{} {
	return m.Properties
}

// SetProperties sets the Properties property
func (m ModelResponse) SetProperties(val map[string]interface{}) {
	m.Properties = val
}

// GetResourceActions returns the ResourceActions property
func (m ModelResponse) GetResourceActions() ModelActions {
	return m.ResourceActions
}

// SetResourceActions sets the ResourceActions property
func (m ModelResponse) SetResourceActions(val ModelActions) {
	m.ResourceActions = val
}

// GetStats returns the Stats property
func (m ModelResponse) GetStats() ModelStatistics {
	return m.Stats
}

// SetStats sets the Stats property
func (m ModelResponse) SetStats(val ModelStatistics) {
	m.Stats = val
}

// GetType returns the Type property
func (m ModelResponse) GetType() ModelType {
	return m.Type
}

// SetType sets the Type property
func (m ModelResponse) SetType(val ModelType) {
	m.Type = val
}

// GetUpdatedAt returns the UpdatedAt property
func (m ModelResponse) GetUpdatedAt() time.Time {
	return m.UpdatedAt
}

// SetUpdatedAt sets the UpdatedAt property
func (m ModelResponse) SetUpdatedAt(val time.Time) {
	m.UpdatedAt = val
}
