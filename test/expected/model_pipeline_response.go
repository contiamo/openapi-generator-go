// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

import (
	"time"
)

// PipelineResponse is an object.
type PipelineResponse struct {
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
	// IsOrphaned: Indicates that this pipeline couldn't be synced from an external system
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
	// ResourceActions: List of actions allowed to perform on the pipeline
	ResourceActions PipelineActions `json:"resourceActions"`
	// Stats: Various statistics about the pipeline
	Stats PipelineStatistics `json:"stats"`
	// Type:
	Type PipelineType `json:"type"`
	// UpdatedAt:
	UpdatedAt time.Time `json:"updatedAt"`
}

// GetCreatedAt returns the CreatedAt property
func (m PipelineResponse) GetCreatedAt() time.Time {
	return m.CreatedAt
}

// SetCreatedAt sets the CreatedAt property
func (m PipelineResponse) SetCreatedAt(val time.Time) {
	m.CreatedAt = val
}

// GetDescription returns the Description property
func (m PipelineResponse) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m PipelineResponse) SetDescription(val string) {
	m.Description = val
}

// GetDisplayName returns the DisplayName property
func (m PipelineResponse) GetDisplayName() string {
	return m.DisplayName
}

// SetDisplayName sets the DisplayName property
func (m PipelineResponse) SetDisplayName(val string) {
	m.DisplayName = val
}

// GetDocumentation returns the Documentation property
func (m PipelineResponse) GetDocumentation() string {
	return m.Documentation
}

// SetDocumentation sets the Documentation property
func (m PipelineResponse) SetDocumentation(val string) {
	m.Documentation = val
}

// GetExternalId returns the ExternalId property
func (m PipelineResponse) GetExternalId() *string {
	return m.ExternalId
}

// SetExternalId sets the ExternalId property
func (m PipelineResponse) SetExternalId(val *string) {
	m.ExternalId = val
}

// GetId returns the Id property
func (m PipelineResponse) GetId() string {
	return m.Id
}

// SetId sets the Id property
func (m PipelineResponse) SetId(val string) {
	m.Id = val
}

// GetIsOrphaned returns the IsOrphaned property
func (m PipelineResponse) GetIsOrphaned() bool {
	return m.IsOrphaned
}

// SetIsOrphaned sets the IsOrphaned property
func (m PipelineResponse) SetIsOrphaned(val bool) {
	m.IsOrphaned = val
}

// GetIsPublic returns the IsPublic property
func (m PipelineResponse) GetIsPublic() bool {
	return m.IsPublic
}

// SetIsPublic sets the IsPublic property
func (m PipelineResponse) SetIsPublic(val bool) {
	m.IsPublic = val
}

// GetName returns the Name property
func (m PipelineResponse) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m PipelineResponse) SetName(val string) {
	m.Name = val
}

// GetOrphanedAt returns the OrphanedAt property
func (m PipelineResponse) GetOrphanedAt() *time.Time {
	return m.OrphanedAt
}

// SetOrphanedAt sets the OrphanedAt property
func (m PipelineResponse) SetOrphanedAt(val *time.Time) {
	m.OrphanedAt = val
}

// GetProjectId returns the ProjectId property
func (m PipelineResponse) GetProjectId() string {
	return m.ProjectId
}

// SetProjectId sets the ProjectId property
func (m PipelineResponse) SetProjectId(val string) {
	m.ProjectId = val
}

// GetProperties returns the Properties property
func (m PipelineResponse) GetProperties() map[string]interface{} {
	return m.Properties
}

// SetProperties sets the Properties property
func (m PipelineResponse) SetProperties(val map[string]interface{}) {
	m.Properties = val
}

// GetResourceActions returns the ResourceActions property
func (m PipelineResponse) GetResourceActions() PipelineActions {
	return m.ResourceActions
}

// SetResourceActions sets the ResourceActions property
func (m PipelineResponse) SetResourceActions(val PipelineActions) {
	m.ResourceActions = val
}

// GetStats returns the Stats property
func (m PipelineResponse) GetStats() PipelineStatistics {
	return m.Stats
}

// SetStats sets the Stats property
func (m PipelineResponse) SetStats(val PipelineStatistics) {
	m.Stats = val
}

// GetType returns the Type property
func (m PipelineResponse) GetType() PipelineType {
	return m.Type
}

// SetType sets the Type property
func (m PipelineResponse) SetType(val PipelineType) {
	m.Type = val
}

// GetUpdatedAt returns the UpdatedAt property
func (m PipelineResponse) GetUpdatedAt() time.Time {
	return m.UpdatedAt
}

// SetUpdatedAt sets the UpdatedAt property
func (m PipelineResponse) SetUpdatedAt(val time.Time) {
	m.UpdatedAt = val
}
