// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// GlossaryItemUpdateRequest is an object.
type GlossaryItemUpdateRequest struct {
	// Description:
	Description *string `json:"description,omitempty"`
	// DisplayName:
	DisplayName *string `json:"displayName,omitempty"`
	// Documentation:
	Documentation *string `json:"documentation,omitempty"`
	// ExternalId:
	ExternalId *string `json:"externalId,omitempty"`
	// GlossaryParentId:
	GlossaryParentId *string `json:"glossaryParentId,omitempty"`
	// IsPublic:
	IsPublic *bool `json:"isPublic,omitempty"`
	// Properties:
	Properties map[string]interface{} `json:"properties,omitempty"`
}

// GetDescription returns the Description property
func (m GlossaryItemUpdateRequest) GetDescription() *string {
	return m.Description
}

// SetDescription sets the Description property
func (m GlossaryItemUpdateRequest) SetDescription(val *string) {
	m.Description = val
}

// GetDisplayName returns the DisplayName property
func (m GlossaryItemUpdateRequest) GetDisplayName() *string {
	return m.DisplayName
}

// SetDisplayName sets the DisplayName property
func (m GlossaryItemUpdateRequest) SetDisplayName(val *string) {
	m.DisplayName = val
}

// GetDocumentation returns the Documentation property
func (m GlossaryItemUpdateRequest) GetDocumentation() *string {
	return m.Documentation
}

// SetDocumentation sets the Documentation property
func (m GlossaryItemUpdateRequest) SetDocumentation(val *string) {
	m.Documentation = val
}

// GetExternalId returns the ExternalId property
func (m GlossaryItemUpdateRequest) GetExternalId() *string {
	return m.ExternalId
}

// SetExternalId sets the ExternalId property
func (m GlossaryItemUpdateRequest) SetExternalId(val *string) {
	m.ExternalId = val
}

// GetGlossaryParentId returns the GlossaryParentId property
func (m GlossaryItemUpdateRequest) GetGlossaryParentId() *string {
	return m.GlossaryParentId
}

// SetGlossaryParentId sets the GlossaryParentId property
func (m GlossaryItemUpdateRequest) SetGlossaryParentId(val *string) {
	m.GlossaryParentId = val
}

// GetIsPublic returns the IsPublic property
func (m GlossaryItemUpdateRequest) GetIsPublic() *bool {
	return m.IsPublic
}

// SetIsPublic sets the IsPublic property
func (m GlossaryItemUpdateRequest) SetIsPublic(val *bool) {
	m.IsPublic = val
}

// GetProperties returns the Properties property
func (m GlossaryItemUpdateRequest) GetProperties() map[string]interface{} {
	return m.Properties
}

// SetProperties sets the Properties property
func (m GlossaryItemUpdateRequest) SetProperties(val map[string]interface{}) {
	m.Properties = val
}
