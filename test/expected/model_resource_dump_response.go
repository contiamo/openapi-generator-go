// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

import (
	"time"
)

// ResourceDumpResponse is an object.
type ResourceDumpResponse struct {
	// Children: List of responses that holds matches that are direct children of this resource
	Children []ResourceDumpResponse `json:"children,omitempty"`
	// CreatedAt:
	CreatedAt time.Time `json:"createdAt"`
	// CustomProperties:
	CustomProperties map[string]interface{} `json:"customProperties"`
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
	// IsPublic:
	IsPublic bool `json:"isPublic"`
	// Kind: A list of supported resource kinds
	Kind ResourceKind `json:"kind"`
	// Name:
	Name string `json:"name"`
	// ParentId:
	ParentId *string `json:"parentId,omitempty"`
	// ProjectId:
	ProjectId string `json:"projectId"`
	// Properties:
	Properties interface{} `json:"properties"`
	// Stats:
	Stats interface{} `json:"stats"`
	// UpdatedAt:
	UpdatedAt time.Time `json:"updatedAt"`
}

// GetChildren returns the Children property
func (m ResourceDumpResponse) GetChildren() []ResourceDumpResponse {
	return m.Children
}

// SetChildren sets the Children property
func (m ResourceDumpResponse) SetChildren(val []ResourceDumpResponse) {
	m.Children = val
}

// GetCreatedAt returns the CreatedAt property
func (m ResourceDumpResponse) GetCreatedAt() time.Time {
	return m.CreatedAt
}

// SetCreatedAt sets the CreatedAt property
func (m ResourceDumpResponse) SetCreatedAt(val time.Time) {
	m.CreatedAt = val
}

// GetCustomProperties returns the CustomProperties property
func (m ResourceDumpResponse) GetCustomProperties() map[string]interface{} {
	return m.CustomProperties
}

// SetCustomProperties sets the CustomProperties property
func (m ResourceDumpResponse) SetCustomProperties(val map[string]interface{}) {
	m.CustomProperties = val
}

// GetDescription returns the Description property
func (m ResourceDumpResponse) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m ResourceDumpResponse) SetDescription(val string) {
	m.Description = val
}

// GetDisplayName returns the DisplayName property
func (m ResourceDumpResponse) GetDisplayName() string {
	return m.DisplayName
}

// SetDisplayName sets the DisplayName property
func (m ResourceDumpResponse) SetDisplayName(val string) {
	m.DisplayName = val
}

// GetDocumentation returns the Documentation property
func (m ResourceDumpResponse) GetDocumentation() string {
	return m.Documentation
}

// SetDocumentation sets the Documentation property
func (m ResourceDumpResponse) SetDocumentation(val string) {
	m.Documentation = val
}

// GetExternalId returns the ExternalId property
func (m ResourceDumpResponse) GetExternalId() *string {
	return m.ExternalId
}

// SetExternalId sets the ExternalId property
func (m ResourceDumpResponse) SetExternalId(val *string) {
	m.ExternalId = val
}

// GetIcon returns the Icon property
func (m ResourceDumpResponse) GetIcon() string {
	return m.Icon
}

// SetIcon sets the Icon property
func (m ResourceDumpResponse) SetIcon(val string) {
	m.Icon = val
}

// GetId returns the Id property
func (m ResourceDumpResponse) GetId() string {
	return m.Id
}

// SetId sets the Id property
func (m ResourceDumpResponse) SetId(val string) {
	m.Id = val
}

// GetIsPublic returns the IsPublic property
func (m ResourceDumpResponse) GetIsPublic() bool {
	return m.IsPublic
}

// SetIsPublic sets the IsPublic property
func (m ResourceDumpResponse) SetIsPublic(val bool) {
	m.IsPublic = val
}

// GetKind returns the Kind property
func (m ResourceDumpResponse) GetKind() ResourceKind {
	return m.Kind
}

// SetKind sets the Kind property
func (m ResourceDumpResponse) SetKind(val ResourceKind) {
	m.Kind = val
}

// GetName returns the Name property
func (m ResourceDumpResponse) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m ResourceDumpResponse) SetName(val string) {
	m.Name = val
}

// GetParentId returns the ParentId property
func (m ResourceDumpResponse) GetParentId() *string {
	return m.ParentId
}

// SetParentId sets the ParentId property
func (m ResourceDumpResponse) SetParentId(val *string) {
	m.ParentId = val
}

// GetProjectId returns the ProjectId property
func (m ResourceDumpResponse) GetProjectId() string {
	return m.ProjectId
}

// SetProjectId sets the ProjectId property
func (m ResourceDumpResponse) SetProjectId(val string) {
	m.ProjectId = val
}

// GetProperties returns the Properties property
func (m ResourceDumpResponse) GetProperties() interface{} {
	return m.Properties
}

// SetProperties sets the Properties property
func (m ResourceDumpResponse) SetProperties(val interface{}) {
	m.Properties = val
}

// GetStats returns the Stats property
func (m ResourceDumpResponse) GetStats() interface{} {
	return m.Stats
}

// SetStats sets the Stats property
func (m ResourceDumpResponse) SetStats(val interface{}) {
	m.Stats = val
}

// GetUpdatedAt returns the UpdatedAt property
func (m ResourceDumpResponse) GetUpdatedAt() time.Time {
	return m.UpdatedAt
}

// SetUpdatedAt sets the UpdatedAt property
func (m ResourceDumpResponse) SetUpdatedAt(val time.Time) {
	m.UpdatedAt = val
}
