// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// ModelCreateRequest is an object.
type ModelCreateRequest struct {
	// Description:
	Description string `json:"description,omitempty"`
	// DisplayName:
	DisplayName string `json:"displayName"`
	// ExternalId:
	ExternalId *string `json:"externalId,omitempty"`
	// IsPublic:
	IsPublic bool `json:"isPublic"`
	// Type:
	Type ModelType `json:"type,omitempty"`
}

// GetDescription returns the Description property
func (m ModelCreateRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m ModelCreateRequest) SetDescription(val string) {
	m.Description = val
}

// GetDisplayName returns the DisplayName property
func (m ModelCreateRequest) GetDisplayName() string {
	return m.DisplayName
}

// SetDisplayName sets the DisplayName property
func (m ModelCreateRequest) SetDisplayName(val string) {
	m.DisplayName = val
}

// GetExternalId returns the ExternalId property
func (m ModelCreateRequest) GetExternalId() *string {
	return m.ExternalId
}

// SetExternalId sets the ExternalId property
func (m ModelCreateRequest) SetExternalId(val *string) {
	m.ExternalId = val
}

// GetIsPublic returns the IsPublic property
func (m ModelCreateRequest) GetIsPublic() bool {
	return m.IsPublic
}

// SetIsPublic sets the IsPublic property
func (m ModelCreateRequest) SetIsPublic(val bool) {
	m.IsPublic = val
}

// GetType returns the Type property
func (m ModelCreateRequest) GetType() ModelType {
	return m.Type
}

// SetType sets the Type property
func (m ModelCreateRequest) SetType(val ModelType) {
	m.Type = val
}
