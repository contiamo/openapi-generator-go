// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// GlossaryItemCreateRequest is an object.
type GlossaryItemCreateRequest struct {
	// Description:
	Description string `json:"description,omitempty"`
	// DisplayName:
	DisplayName string `json:"displayName"`
	// ExternalId:
	ExternalId *string `json:"externalId,omitempty"`
	// GlossaryParentId:
	GlossaryParentId *string `json:"glossaryParentId,omitempty"`
	// IsPublic:
	IsPublic bool `json:"isPublic"`
}

// GetDescription returns the Description property
func (m GlossaryItemCreateRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m GlossaryItemCreateRequest) SetDescription(val string) {
	m.Description = val
}

// GetDisplayName returns the DisplayName property
func (m GlossaryItemCreateRequest) GetDisplayName() string {
	return m.DisplayName
}

// SetDisplayName sets the DisplayName property
func (m GlossaryItemCreateRequest) SetDisplayName(val string) {
	m.DisplayName = val
}

// GetExternalId returns the ExternalId property
func (m GlossaryItemCreateRequest) GetExternalId() *string {
	return m.ExternalId
}

// SetExternalId sets the ExternalId property
func (m GlossaryItemCreateRequest) SetExternalId(val *string) {
	m.ExternalId = val
}

// GetGlossaryParentId returns the GlossaryParentId property
func (m GlossaryItemCreateRequest) GetGlossaryParentId() *string {
	return m.GlossaryParentId
}

// SetGlossaryParentId sets the GlossaryParentId property
func (m GlossaryItemCreateRequest) SetGlossaryParentId(val *string) {
	m.GlossaryParentId = val
}

// GetIsPublic returns the IsPublic property
func (m GlossaryItemCreateRequest) GetIsPublic() bool {
	return m.IsPublic
}

// SetIsPublic sets the IsPublic property
func (m GlossaryItemCreateRequest) SetIsPublic(val bool) {
	m.IsPublic = val
}
