// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

import (
	"time"
)

// UseCaseResponse is an object.
type UseCaseResponse struct {
	// CreatedAt:
	CreatedAt time.Time `json:"createdAt"`
	// Description:
	Description string `json:"description"`
	// DisplayName:
	DisplayName string `json:"displayName"`
	// Documentation: Long form markdown field that can be used to document the resource
	Documentation string `json:"documentation"`
	// ExternalId:
	ExternalId *string `json:"externalId,omitempty"`
	// Icon: Name of an icon that is assigned to the data source
	Icon string `json:"icon"`
	// Id:
	Id string `json:"id"`
	// IsOrphaned: Indicates that this usecase couldn't be synced from an external system
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
	// ResourceActions: List of actions allowed to perform on the use case
	ResourceActions UseCaseActions `json:"resourceActions"`
	// Stats: Various statistics about the use case
	Stats UseCaseStatistics `json:"stats"`
	// UpdatedAt:
	UpdatedAt time.Time `json:"updatedAt"`
}

// GetCreatedAt returns the CreatedAt property
func (m UseCaseResponse) GetCreatedAt() time.Time {
	return m.CreatedAt
}

// SetCreatedAt sets the CreatedAt property
func (m UseCaseResponse) SetCreatedAt(val time.Time) {
	m.CreatedAt = val
}

// GetDescription returns the Description property
func (m UseCaseResponse) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m UseCaseResponse) SetDescription(val string) {
	m.Description = val
}

// GetDisplayName returns the DisplayName property
func (m UseCaseResponse) GetDisplayName() string {
	return m.DisplayName
}

// SetDisplayName sets the DisplayName property
func (m UseCaseResponse) SetDisplayName(val string) {
	m.DisplayName = val
}

// GetDocumentation returns the Documentation property
func (m UseCaseResponse) GetDocumentation() string {
	return m.Documentation
}

// SetDocumentation sets the Documentation property
func (m UseCaseResponse) SetDocumentation(val string) {
	m.Documentation = val
}

// GetExternalId returns the ExternalId property
func (m UseCaseResponse) GetExternalId() *string {
	return m.ExternalId
}

// SetExternalId sets the ExternalId property
func (m UseCaseResponse) SetExternalId(val *string) {
	m.ExternalId = val
}

// GetIcon returns the Icon property
func (m UseCaseResponse) GetIcon() string {
	return m.Icon
}

// SetIcon sets the Icon property
func (m UseCaseResponse) SetIcon(val string) {
	m.Icon = val
}

// GetId returns the Id property
func (m UseCaseResponse) GetId() string {
	return m.Id
}

// SetId sets the Id property
func (m UseCaseResponse) SetId(val string) {
	m.Id = val
}

// GetIsOrphaned returns the IsOrphaned property
func (m UseCaseResponse) GetIsOrphaned() bool {
	return m.IsOrphaned
}

// SetIsOrphaned sets the IsOrphaned property
func (m UseCaseResponse) SetIsOrphaned(val bool) {
	m.IsOrphaned = val
}

// GetIsPublic returns the IsPublic property
func (m UseCaseResponse) GetIsPublic() bool {
	return m.IsPublic
}

// SetIsPublic sets the IsPublic property
func (m UseCaseResponse) SetIsPublic(val bool) {
	m.IsPublic = val
}

// GetName returns the Name property
func (m UseCaseResponse) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m UseCaseResponse) SetName(val string) {
	m.Name = val
}

// GetOrphanedAt returns the OrphanedAt property
func (m UseCaseResponse) GetOrphanedAt() *time.Time {
	return m.OrphanedAt
}

// SetOrphanedAt sets the OrphanedAt property
func (m UseCaseResponse) SetOrphanedAt(val *time.Time) {
	m.OrphanedAt = val
}

// GetProjectId returns the ProjectId property
func (m UseCaseResponse) GetProjectId() string {
	return m.ProjectId
}

// SetProjectId sets the ProjectId property
func (m UseCaseResponse) SetProjectId(val string) {
	m.ProjectId = val
}

// GetProperties returns the Properties property
func (m UseCaseResponse) GetProperties() map[string]interface{} {
	return m.Properties
}

// SetProperties sets the Properties property
func (m UseCaseResponse) SetProperties(val map[string]interface{}) {
	m.Properties = val
}

// GetResourceActions returns the ResourceActions property
func (m UseCaseResponse) GetResourceActions() UseCaseActions {
	return m.ResourceActions
}

// SetResourceActions sets the ResourceActions property
func (m UseCaseResponse) SetResourceActions(val UseCaseActions) {
	m.ResourceActions = val
}

// GetStats returns the Stats property
func (m UseCaseResponse) GetStats() UseCaseStatistics {
	return m.Stats
}

// SetStats sets the Stats property
func (m UseCaseResponse) SetStats(val UseCaseStatistics) {
	m.Stats = val
}

// GetUpdatedAt returns the UpdatedAt property
func (m UseCaseResponse) GetUpdatedAt() time.Time {
	return m.UpdatedAt
}

// SetUpdatedAt sets the UpdatedAt property
func (m UseCaseResponse) SetUpdatedAt(val time.Time) {
	m.UpdatedAt = val
}
