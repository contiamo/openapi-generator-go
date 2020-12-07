// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

import (
	"time"
)

// StreamResponse is an object.
type StreamResponse struct {
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
	// IsOrphaned: Indicates that this stream couldn't be synced from an external system
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
	// ResourceActions: List of actions allowed to perform on the stream
	ResourceActions StreamActions `json:"resourceActions"`
	// Stats: Various statistics about the stream
	Stats StreamStatistics `json:"stats"`
	// Type:
	Type StreamType `json:"type"`
	// UpdatedAt:
	UpdatedAt time.Time `json:"updatedAt"`
}

// GetCreatedAt returns the CreatedAt property
func (m StreamResponse) GetCreatedAt() time.Time {
	return m.CreatedAt
}

// SetCreatedAt sets the CreatedAt property
func (m StreamResponse) SetCreatedAt(val time.Time) {
	m.CreatedAt = val
}

// GetDescription returns the Description property
func (m StreamResponse) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m StreamResponse) SetDescription(val string) {
	m.Description = val
}

// GetDisplayName returns the DisplayName property
func (m StreamResponse) GetDisplayName() string {
	return m.DisplayName
}

// SetDisplayName sets the DisplayName property
func (m StreamResponse) SetDisplayName(val string) {
	m.DisplayName = val
}

// GetDocumentation returns the Documentation property
func (m StreamResponse) GetDocumentation() string {
	return m.Documentation
}

// SetDocumentation sets the Documentation property
func (m StreamResponse) SetDocumentation(val string) {
	m.Documentation = val
}

// GetExternalId returns the ExternalId property
func (m StreamResponse) GetExternalId() *string {
	return m.ExternalId
}

// SetExternalId sets the ExternalId property
func (m StreamResponse) SetExternalId(val *string) {
	m.ExternalId = val
}

// GetId returns the Id property
func (m StreamResponse) GetId() string {
	return m.Id
}

// SetId sets the Id property
func (m StreamResponse) SetId(val string) {
	m.Id = val
}

// GetIsOrphaned returns the IsOrphaned property
func (m StreamResponse) GetIsOrphaned() bool {
	return m.IsOrphaned
}

// SetIsOrphaned sets the IsOrphaned property
func (m StreamResponse) SetIsOrphaned(val bool) {
	m.IsOrphaned = val
}

// GetIsPublic returns the IsPublic property
func (m StreamResponse) GetIsPublic() bool {
	return m.IsPublic
}

// SetIsPublic sets the IsPublic property
func (m StreamResponse) SetIsPublic(val bool) {
	m.IsPublic = val
}

// GetName returns the Name property
func (m StreamResponse) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m StreamResponse) SetName(val string) {
	m.Name = val
}

// GetOrphanedAt returns the OrphanedAt property
func (m StreamResponse) GetOrphanedAt() *time.Time {
	return m.OrphanedAt
}

// SetOrphanedAt sets the OrphanedAt property
func (m StreamResponse) SetOrphanedAt(val *time.Time) {
	m.OrphanedAt = val
}

// GetProjectId returns the ProjectId property
func (m StreamResponse) GetProjectId() string {
	return m.ProjectId
}

// SetProjectId sets the ProjectId property
func (m StreamResponse) SetProjectId(val string) {
	m.ProjectId = val
}

// GetProperties returns the Properties property
func (m StreamResponse) GetProperties() map[string]interface{} {
	return m.Properties
}

// SetProperties sets the Properties property
func (m StreamResponse) SetProperties(val map[string]interface{}) {
	m.Properties = val
}

// GetResourceActions returns the ResourceActions property
func (m StreamResponse) GetResourceActions() StreamActions {
	return m.ResourceActions
}

// SetResourceActions sets the ResourceActions property
func (m StreamResponse) SetResourceActions(val StreamActions) {
	m.ResourceActions = val
}

// GetStats returns the Stats property
func (m StreamResponse) GetStats() StreamStatistics {
	return m.Stats
}

// SetStats sets the Stats property
func (m StreamResponse) SetStats(val StreamStatistics) {
	m.Stats = val
}

// GetType returns the Type property
func (m StreamResponse) GetType() StreamType {
	return m.Type
}

// SetType sets the Type property
func (m StreamResponse) SetType(val StreamType) {
	m.Type = val
}

// GetUpdatedAt returns the UpdatedAt property
func (m StreamResponse) GetUpdatedAt() time.Time {
	return m.UpdatedAt
}

// SetUpdatedAt sets the UpdatedAt property
func (m StreamResponse) SetUpdatedAt(val time.Time) {
	m.UpdatedAt = val
}
