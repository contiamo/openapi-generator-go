// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

import (
	"time"
)

// BIReportResponse is an object.
type BIReportResponse struct {
	// CreatedAt:
	CreatedAt time.Time `json:"createdAt"`
	// Description:
	Description string `json:"description"`
	// DisplayName:
	DisplayName string `json:"displayName"`
	// Documentation: Long form markdown field that can be used to document the resource
	Documentation string `json:"documentation"`
	// ExternalEmbeddableURL:
	ExternalEmbeddableURL string `json:"externalEmbeddableURL"`
	// ExternalId:
	ExternalId *string `json:"externalId"`
	// ExternalThumbnailURL:
	ExternalThumbnailURL string `json:"externalThumbnailURL"`
	// ExternalURL:
	ExternalURL string `json:"externalURL"`
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
	// ResourceActions: List of actions allowed to perform on the BI Report
	ResourceActions BIReportActions `json:"resourceActions"`
	// Stats: Various statistics about the BI Report
	Stats BIReportStatistics `json:"stats"`
	// Type:
	Type BIReportType `json:"type"`
	// UpdatedAt:
	UpdatedAt time.Time `json:"updatedAt"`
}

// GetCreatedAt returns the CreatedAt property
func (m BIReportResponse) GetCreatedAt() time.Time {
	return m.CreatedAt
}

// SetCreatedAt sets the CreatedAt property
func (m BIReportResponse) SetCreatedAt(val time.Time) {
	m.CreatedAt = val
}

// GetDescription returns the Description property
func (m BIReportResponse) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m BIReportResponse) SetDescription(val string) {
	m.Description = val
}

// GetDisplayName returns the DisplayName property
func (m BIReportResponse) GetDisplayName() string {
	return m.DisplayName
}

// SetDisplayName sets the DisplayName property
func (m BIReportResponse) SetDisplayName(val string) {
	m.DisplayName = val
}

// GetDocumentation returns the Documentation property
func (m BIReportResponse) GetDocumentation() string {
	return m.Documentation
}

// SetDocumentation sets the Documentation property
func (m BIReportResponse) SetDocumentation(val string) {
	m.Documentation = val
}

// GetExternalEmbeddableURL returns the ExternalEmbeddableURL property
func (m BIReportResponse) GetExternalEmbeddableURL() string {
	return m.ExternalEmbeddableURL
}

// SetExternalEmbeddableURL sets the ExternalEmbeddableURL property
func (m BIReportResponse) SetExternalEmbeddableURL(val string) {
	m.ExternalEmbeddableURL = val
}

// GetExternalId returns the ExternalId property
func (m BIReportResponse) GetExternalId() *string {
	return m.ExternalId
}

// SetExternalId sets the ExternalId property
func (m BIReportResponse) SetExternalId(val *string) {
	m.ExternalId = val
}

// GetExternalThumbnailURL returns the ExternalThumbnailURL property
func (m BIReportResponse) GetExternalThumbnailURL() string {
	return m.ExternalThumbnailURL
}

// SetExternalThumbnailURL sets the ExternalThumbnailURL property
func (m BIReportResponse) SetExternalThumbnailURL(val string) {
	m.ExternalThumbnailURL = val
}

// GetExternalURL returns the ExternalURL property
func (m BIReportResponse) GetExternalURL() string {
	return m.ExternalURL
}

// SetExternalURL sets the ExternalURL property
func (m BIReportResponse) SetExternalURL(val string) {
	m.ExternalURL = val
}

// GetIcon returns the Icon property
func (m BIReportResponse) GetIcon() string {
	return m.Icon
}

// SetIcon sets the Icon property
func (m BIReportResponse) SetIcon(val string) {
	m.Icon = val
}

// GetId returns the Id property
func (m BIReportResponse) GetId() string {
	return m.Id
}

// SetId sets the Id property
func (m BIReportResponse) SetId(val string) {
	m.Id = val
}

// GetIsOrphaned returns the IsOrphaned property
func (m BIReportResponse) GetIsOrphaned() bool {
	return m.IsOrphaned
}

// SetIsOrphaned sets the IsOrphaned property
func (m BIReportResponse) SetIsOrphaned(val bool) {
	m.IsOrphaned = val
}

// GetIsPublic returns the IsPublic property
func (m BIReportResponse) GetIsPublic() bool {
	return m.IsPublic
}

// SetIsPublic sets the IsPublic property
func (m BIReportResponse) SetIsPublic(val bool) {
	m.IsPublic = val
}

// GetName returns the Name property
func (m BIReportResponse) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m BIReportResponse) SetName(val string) {
	m.Name = val
}

// GetOrphanedAt returns the OrphanedAt property
func (m BIReportResponse) GetOrphanedAt() *time.Time {
	return m.OrphanedAt
}

// SetOrphanedAt sets the OrphanedAt property
func (m BIReportResponse) SetOrphanedAt(val *time.Time) {
	m.OrphanedAt = val
}

// GetProjectId returns the ProjectId property
func (m BIReportResponse) GetProjectId() string {
	return m.ProjectId
}

// SetProjectId sets the ProjectId property
func (m BIReportResponse) SetProjectId(val string) {
	m.ProjectId = val
}

// GetProperties returns the Properties property
func (m BIReportResponse) GetProperties() map[string]interface{} {
	return m.Properties
}

// SetProperties sets the Properties property
func (m BIReportResponse) SetProperties(val map[string]interface{}) {
	m.Properties = val
}

// GetResourceActions returns the ResourceActions property
func (m BIReportResponse) GetResourceActions() BIReportActions {
	return m.ResourceActions
}

// SetResourceActions sets the ResourceActions property
func (m BIReportResponse) SetResourceActions(val BIReportActions) {
	m.ResourceActions = val
}

// GetStats returns the Stats property
func (m BIReportResponse) GetStats() BIReportStatistics {
	return m.Stats
}

// SetStats sets the Stats property
func (m BIReportResponse) SetStats(val BIReportStatistics) {
	m.Stats = val
}

// GetType returns the Type property
func (m BIReportResponse) GetType() BIReportType {
	return m.Type
}

// SetType sets the Type property
func (m BIReportResponse) SetType(val BIReportType) {
	m.Type = val
}

// GetUpdatedAt returns the UpdatedAt property
func (m BIReportResponse) GetUpdatedAt() time.Time {
	return m.UpdatedAt
}

// SetUpdatedAt sets the UpdatedAt property
func (m BIReportResponse) SetUpdatedAt(val time.Time) {
	m.UpdatedAt = val
}
