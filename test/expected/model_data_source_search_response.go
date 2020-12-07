// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

import (
	"time"
)

// DataSourceSearchResponse is an object.
type DataSourceSearchResponse struct {
	// CreatedAt:
	CreatedAt time.Time `json:"createdAt"`
	// Description:
	Description string `json:"description"`
	// DisplayName:
	DisplayName string `json:"displayName,omitempty"`
	// Highlight: contains the matching parts of the description
	Highlight string `json:"highlight,omitempty"`
	// Icon: Name of an icon that is assigned to the data source
	Icon string `json:"icon"`
	// Id:
	Id string `json:"id"`
	// IsPublic:
	IsPublic bool `json:"isPublic"`
	// MatchingTableCount: Total number of matching tables in this data source
	MatchingTableCount int32 `json:"matchingTableCount,omitempty"`
	// Name:
	Name string `json:"name"`
	// ProjectId:
	ProjectId string `json:"projectId"`
	// Properties:
	Properties map[string]interface{} `json:"properties"`
	// ResourceActions: List of actions allowed to perform on the data source
	ResourceActions DataSourceActions `json:"resourceActions"`
	// Stats: Various statistics about the data source
	Stats DataSourceStatistics `json:"stats,omitempty"`
	// Tables: First N matching tables, N equals tableCount query parameter
	Tables []TableSearchResponse `json:"tables,omitempty"`
	// Technology: Type of a data source
	Technology DataSourceTechnology `json:"technology,omitempty"`
	// Type: Type of a data source
	Type DataSourceType `json:"type"`
	// UpdatedAt:
	UpdatedAt time.Time `json:"updatedAt"`
}

// GetCreatedAt returns the CreatedAt property
func (m DataSourceSearchResponse) GetCreatedAt() time.Time {
	return m.CreatedAt
}

// SetCreatedAt sets the CreatedAt property
func (m DataSourceSearchResponse) SetCreatedAt(val time.Time) {
	m.CreatedAt = val
}

// GetDescription returns the Description property
func (m DataSourceSearchResponse) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m DataSourceSearchResponse) SetDescription(val string) {
	m.Description = val
}

// GetDisplayName returns the DisplayName property
func (m DataSourceSearchResponse) GetDisplayName() string {
	return m.DisplayName
}

// SetDisplayName sets the DisplayName property
func (m DataSourceSearchResponse) SetDisplayName(val string) {
	m.DisplayName = val
}

// GetHighlight returns the Highlight property
func (m DataSourceSearchResponse) GetHighlight() string {
	return m.Highlight
}

// SetHighlight sets the Highlight property
func (m DataSourceSearchResponse) SetHighlight(val string) {
	m.Highlight = val
}

// GetIcon returns the Icon property
func (m DataSourceSearchResponse) GetIcon() string {
	return m.Icon
}

// SetIcon sets the Icon property
func (m DataSourceSearchResponse) SetIcon(val string) {
	m.Icon = val
}

// GetId returns the Id property
func (m DataSourceSearchResponse) GetId() string {
	return m.Id
}

// SetId sets the Id property
func (m DataSourceSearchResponse) SetId(val string) {
	m.Id = val
}

// GetIsPublic returns the IsPublic property
func (m DataSourceSearchResponse) GetIsPublic() bool {
	return m.IsPublic
}

// SetIsPublic sets the IsPublic property
func (m DataSourceSearchResponse) SetIsPublic(val bool) {
	m.IsPublic = val
}

// GetMatchingTableCount returns the MatchingTableCount property
func (m DataSourceSearchResponse) GetMatchingTableCount() int32 {
	return m.MatchingTableCount
}

// SetMatchingTableCount sets the MatchingTableCount property
func (m DataSourceSearchResponse) SetMatchingTableCount(val int32) {
	m.MatchingTableCount = val
}

// GetName returns the Name property
func (m DataSourceSearchResponse) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m DataSourceSearchResponse) SetName(val string) {
	m.Name = val
}

// GetProjectId returns the ProjectId property
func (m DataSourceSearchResponse) GetProjectId() string {
	return m.ProjectId
}

// SetProjectId sets the ProjectId property
func (m DataSourceSearchResponse) SetProjectId(val string) {
	m.ProjectId = val
}

// GetProperties returns the Properties property
func (m DataSourceSearchResponse) GetProperties() map[string]interface{} {
	return m.Properties
}

// SetProperties sets the Properties property
func (m DataSourceSearchResponse) SetProperties(val map[string]interface{}) {
	m.Properties = val
}

// GetResourceActions returns the ResourceActions property
func (m DataSourceSearchResponse) GetResourceActions() DataSourceActions {
	return m.ResourceActions
}

// SetResourceActions sets the ResourceActions property
func (m DataSourceSearchResponse) SetResourceActions(val DataSourceActions) {
	m.ResourceActions = val
}

// GetStats returns the Stats property
func (m DataSourceSearchResponse) GetStats() DataSourceStatistics {
	return m.Stats
}

// SetStats sets the Stats property
func (m DataSourceSearchResponse) SetStats(val DataSourceStatistics) {
	m.Stats = val
}

// GetTables returns the Tables property
func (m DataSourceSearchResponse) GetTables() []TableSearchResponse {
	return m.Tables
}

// SetTables sets the Tables property
func (m DataSourceSearchResponse) SetTables(val []TableSearchResponse) {
	m.Tables = val
}

// GetTechnology returns the Technology property
func (m DataSourceSearchResponse) GetTechnology() DataSourceTechnology {
	return m.Technology
}

// SetTechnology sets the Technology property
func (m DataSourceSearchResponse) SetTechnology(val DataSourceTechnology) {
	m.Technology = val
}

// GetType returns the Type property
func (m DataSourceSearchResponse) GetType() DataSourceType {
	return m.Type
}

// SetType sets the Type property
func (m DataSourceSearchResponse) SetType(val DataSourceType) {
	m.Type = val
}

// GetUpdatedAt returns the UpdatedAt property
func (m DataSourceSearchResponse) GetUpdatedAt() time.Time {
	return m.UpdatedAt
}

// SetUpdatedAt sets the UpdatedAt property
func (m DataSourceSearchResponse) SetUpdatedAt(val time.Time) {
	m.UpdatedAt = val
}
