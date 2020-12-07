// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

import (
	"time"
)

// GlossaryItemResponse is an object.
type GlossaryItemResponse struct {
	// CreatedAt:
	CreatedAt time.Time `json:"createdAt"`
	// Description:
	Description string `json:"description"`
	// DisplayName:
	DisplayName string `json:"displayName"`
	// DisplayNamePath: Path of all parent glossary item display names starting from the top level category
	DisplayNamePath []string `json:"displayNamePath"`
	// Documentation: Long form markdown field that can be used to document the resource
	Documentation string `json:"documentation"`
	// ExternalId:
	ExternalId *string `json:"externalId"`
	// GlossaryParentId:
	GlossaryParentId *string `json:"glossaryParentId"`
	// Icon: Name of an icon that is assigned to the resource
	Icon string `json:"icon"`
	// Id:
	Id string `json:"id"`
	// IsPublic:
	IsPublic bool `json:"isPublic"`
	// Name:
	Name string `json:"name"`
	// Path: Path of all parent glossary items starting from the top level category
	Path []string `json:"path"`
	// ProjectId:
	ProjectId string `json:"projectId"`
	// Properties:
	Properties map[string]interface{} `json:"properties"`
	// ResourceActions: List of actions allowed to perform on the glossary item
	ResourceActions GlossaryItemActions `json:"resourceActions"`
	// Stats: Various statistics about the glossary item
	Stats GlossaryItemStatistics `json:"stats"`
	// UpdatedAt:
	UpdatedAt time.Time `json:"updatedAt"`
}

// GetCreatedAt returns the CreatedAt property
func (m GlossaryItemResponse) GetCreatedAt() time.Time {
	return m.CreatedAt
}

// SetCreatedAt sets the CreatedAt property
func (m GlossaryItemResponse) SetCreatedAt(val time.Time) {
	m.CreatedAt = val
}

// GetDescription returns the Description property
func (m GlossaryItemResponse) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m GlossaryItemResponse) SetDescription(val string) {
	m.Description = val
}

// GetDisplayName returns the DisplayName property
func (m GlossaryItemResponse) GetDisplayName() string {
	return m.DisplayName
}

// SetDisplayName sets the DisplayName property
func (m GlossaryItemResponse) SetDisplayName(val string) {
	m.DisplayName = val
}

// GetDisplayNamePath returns the DisplayNamePath property
func (m GlossaryItemResponse) GetDisplayNamePath() []string {
	return m.DisplayNamePath
}

// SetDisplayNamePath sets the DisplayNamePath property
func (m GlossaryItemResponse) SetDisplayNamePath(val []string) {
	m.DisplayNamePath = val
}

// GetDocumentation returns the Documentation property
func (m GlossaryItemResponse) GetDocumentation() string {
	return m.Documentation
}

// SetDocumentation sets the Documentation property
func (m GlossaryItemResponse) SetDocumentation(val string) {
	m.Documentation = val
}

// GetExternalId returns the ExternalId property
func (m GlossaryItemResponse) GetExternalId() *string {
	return m.ExternalId
}

// SetExternalId sets the ExternalId property
func (m GlossaryItemResponse) SetExternalId(val *string) {
	m.ExternalId = val
}

// GetGlossaryParentId returns the GlossaryParentId property
func (m GlossaryItemResponse) GetGlossaryParentId() *string {
	return m.GlossaryParentId
}

// SetGlossaryParentId sets the GlossaryParentId property
func (m GlossaryItemResponse) SetGlossaryParentId(val *string) {
	m.GlossaryParentId = val
}

// GetIcon returns the Icon property
func (m GlossaryItemResponse) GetIcon() string {
	return m.Icon
}

// SetIcon sets the Icon property
func (m GlossaryItemResponse) SetIcon(val string) {
	m.Icon = val
}

// GetId returns the Id property
func (m GlossaryItemResponse) GetId() string {
	return m.Id
}

// SetId sets the Id property
func (m GlossaryItemResponse) SetId(val string) {
	m.Id = val
}

// GetIsPublic returns the IsPublic property
func (m GlossaryItemResponse) GetIsPublic() bool {
	return m.IsPublic
}

// SetIsPublic sets the IsPublic property
func (m GlossaryItemResponse) SetIsPublic(val bool) {
	m.IsPublic = val
}

// GetName returns the Name property
func (m GlossaryItemResponse) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m GlossaryItemResponse) SetName(val string) {
	m.Name = val
}

// GetPath returns the Path property
func (m GlossaryItemResponse) GetPath() []string {
	return m.Path
}

// SetPath sets the Path property
func (m GlossaryItemResponse) SetPath(val []string) {
	m.Path = val
}

// GetProjectId returns the ProjectId property
func (m GlossaryItemResponse) GetProjectId() string {
	return m.ProjectId
}

// SetProjectId sets the ProjectId property
func (m GlossaryItemResponse) SetProjectId(val string) {
	m.ProjectId = val
}

// GetProperties returns the Properties property
func (m GlossaryItemResponse) GetProperties() map[string]interface{} {
	return m.Properties
}

// SetProperties sets the Properties property
func (m GlossaryItemResponse) SetProperties(val map[string]interface{}) {
	m.Properties = val
}

// GetResourceActions returns the ResourceActions property
func (m GlossaryItemResponse) GetResourceActions() GlossaryItemActions {
	return m.ResourceActions
}

// SetResourceActions sets the ResourceActions property
func (m GlossaryItemResponse) SetResourceActions(val GlossaryItemActions) {
	m.ResourceActions = val
}

// GetStats returns the Stats property
func (m GlossaryItemResponse) GetStats() GlossaryItemStatistics {
	return m.Stats
}

// SetStats sets the Stats property
func (m GlossaryItemResponse) SetStats(val GlossaryItemStatistics) {
	m.Stats = val
}

// GetUpdatedAt returns the UpdatedAt property
func (m GlossaryItemResponse) GetUpdatedAt() time.Time {
	return m.UpdatedAt
}

// SetUpdatedAt sets the UpdatedAt property
func (m GlossaryItemResponse) SetUpdatedAt(val time.Time) {
	m.UpdatedAt = val
}
