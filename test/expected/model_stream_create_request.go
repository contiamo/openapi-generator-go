// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// StreamCreateRequest is an object.
type StreamCreateRequest struct {
	// Description:
	Description string `json:"description,omitempty"`
	// DisplayName:
	DisplayName string `json:"displayName"`
	// ExternalId:
	ExternalId *string `json:"externalId,omitempty"`
	// IsPublic:
	IsPublic bool `json:"isPublic"`
	// Type:
	Type StreamType `json:"type,omitempty"`
}

// GetDescription returns the Description property
func (m StreamCreateRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m StreamCreateRequest) SetDescription(val string) {
	m.Description = val
}

// GetDisplayName returns the DisplayName property
func (m StreamCreateRequest) GetDisplayName() string {
	return m.DisplayName
}

// SetDisplayName sets the DisplayName property
func (m StreamCreateRequest) SetDisplayName(val string) {
	m.DisplayName = val
}

// GetExternalId returns the ExternalId property
func (m StreamCreateRequest) GetExternalId() *string {
	return m.ExternalId
}

// SetExternalId sets the ExternalId property
func (m StreamCreateRequest) SetExternalId(val *string) {
	m.ExternalId = val
}

// GetIsPublic returns the IsPublic property
func (m StreamCreateRequest) GetIsPublic() bool {
	return m.IsPublic
}

// SetIsPublic sets the IsPublic property
func (m StreamCreateRequest) SetIsPublic(val bool) {
	m.IsPublic = val
}

// GetType returns the Type property
func (m StreamCreateRequest) GetType() StreamType {
	return m.Type
}

// SetType sets the Type property
func (m StreamCreateRequest) SetType(val StreamType) {
	m.Type = val
}
