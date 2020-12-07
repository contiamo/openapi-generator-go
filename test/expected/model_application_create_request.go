// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// ApplicationCreateRequest is an object.
type ApplicationCreateRequest struct {
	// Description:
	Description string `json:"description,omitempty"`
	// DisplayName:
	DisplayName string `json:"displayName"`
	// ExternalId:
	ExternalId *string `json:"externalId,omitempty"`
	// IsPublic:
	IsPublic bool `json:"isPublic"`
	// Type:
	Type ApplicationType `json:"type,omitempty"`
}

// GetDescription returns the Description property
func (m ApplicationCreateRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m ApplicationCreateRequest) SetDescription(val string) {
	m.Description = val
}

// GetDisplayName returns the DisplayName property
func (m ApplicationCreateRequest) GetDisplayName() string {
	return m.DisplayName
}

// SetDisplayName sets the DisplayName property
func (m ApplicationCreateRequest) SetDisplayName(val string) {
	m.DisplayName = val
}

// GetExternalId returns the ExternalId property
func (m ApplicationCreateRequest) GetExternalId() *string {
	return m.ExternalId
}

// SetExternalId sets the ExternalId property
func (m ApplicationCreateRequest) SetExternalId(val *string) {
	m.ExternalId = val
}

// GetIsPublic returns the IsPublic property
func (m ApplicationCreateRequest) GetIsPublic() bool {
	return m.IsPublic
}

// SetIsPublic sets the IsPublic property
func (m ApplicationCreateRequest) SetIsPublic(val bool) {
	m.IsPublic = val
}

// GetType returns the Type property
func (m ApplicationCreateRequest) GetType() ApplicationType {
	return m.Type
}

// SetType sets the Type property
func (m ApplicationCreateRequest) SetType(val ApplicationType) {
	m.Type = val
}
