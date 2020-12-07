// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// UseCaseCreateRequest is an object.
type UseCaseCreateRequest struct {
	// Description:
	Description string `json:"description,omitempty"`
	// DisplayName:
	DisplayName string `json:"displayName"`
	// ExternalId:
	ExternalId *string `json:"externalId,omitempty"`
	// IsPublic:
	IsPublic bool `json:"isPublic"`
}

// GetDescription returns the Description property
func (m UseCaseCreateRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m UseCaseCreateRequest) SetDescription(val string) {
	m.Description = val
}

// GetDisplayName returns the DisplayName property
func (m UseCaseCreateRequest) GetDisplayName() string {
	return m.DisplayName
}

// SetDisplayName sets the DisplayName property
func (m UseCaseCreateRequest) SetDisplayName(val string) {
	m.DisplayName = val
}

// GetExternalId returns the ExternalId property
func (m UseCaseCreateRequest) GetExternalId() *string {
	return m.ExternalId
}

// SetExternalId sets the ExternalId property
func (m UseCaseCreateRequest) SetExternalId(val *string) {
	m.ExternalId = val
}

// GetIsPublic returns the IsPublic property
func (m UseCaseCreateRequest) GetIsPublic() bool {
	return m.IsPublic
}

// SetIsPublic sets the IsPublic property
func (m UseCaseCreateRequest) SetIsPublic(val bool) {
	m.IsPublic = val
}
