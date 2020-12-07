// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// DataSourceCreateRequest is an object.
type DataSourceCreateRequest struct {
	// ConnectionInfo:
	ConnectionInfo ConnectionProperties `json:"connectionInfo,omitempty"`
	// Description:
	Description string `json:"description,omitempty"`
	// DisplayName:
	DisplayName string `json:"displayName,omitempty"`
	// ExternalId:
	ExternalId *string `json:"externalId,omitempty"`
	// IsPublic:
	IsPublic bool `json:"isPublic"`
	// Name:
	Name string `json:"name"`
	// Technology: Type of a data source
	Technology DataSourceTechnology `json:"technology,omitempty"`
	// Type: Type of a data source
	Type DataSourceType `json:"type"`
}

// GetConnectionInfo returns the ConnectionInfo property
func (m DataSourceCreateRequest) GetConnectionInfo() ConnectionProperties {
	return m.ConnectionInfo
}

// SetConnectionInfo sets the ConnectionInfo property
func (m DataSourceCreateRequest) SetConnectionInfo(val ConnectionProperties) {
	m.ConnectionInfo = val
}

// GetDescription returns the Description property
func (m DataSourceCreateRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m DataSourceCreateRequest) SetDescription(val string) {
	m.Description = val
}

// GetDisplayName returns the DisplayName property
func (m DataSourceCreateRequest) GetDisplayName() string {
	return m.DisplayName
}

// SetDisplayName sets the DisplayName property
func (m DataSourceCreateRequest) SetDisplayName(val string) {
	m.DisplayName = val
}

// GetExternalId returns the ExternalId property
func (m DataSourceCreateRequest) GetExternalId() *string {
	return m.ExternalId
}

// SetExternalId sets the ExternalId property
func (m DataSourceCreateRequest) SetExternalId(val *string) {
	m.ExternalId = val
}

// GetIsPublic returns the IsPublic property
func (m DataSourceCreateRequest) GetIsPublic() bool {
	return m.IsPublic
}

// SetIsPublic sets the IsPublic property
func (m DataSourceCreateRequest) SetIsPublic(val bool) {
	m.IsPublic = val
}

// GetName returns the Name property
func (m DataSourceCreateRequest) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m DataSourceCreateRequest) SetName(val string) {
	m.Name = val
}

// GetTechnology returns the Technology property
func (m DataSourceCreateRequest) GetTechnology() DataSourceTechnology {
	return m.Technology
}

// SetTechnology sets the Technology property
func (m DataSourceCreateRequest) SetTechnology(val DataSourceTechnology) {
	m.Technology = val
}

// GetType returns the Type property
func (m DataSourceCreateRequest) GetType() DataSourceType {
	return m.Type
}

// SetType sets the Type property
func (m DataSourceCreateRequest) SetType(val DataSourceType) {
	m.Type = val
}
