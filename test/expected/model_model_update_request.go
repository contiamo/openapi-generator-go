// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// ModelUpdateRequest is an object.
type ModelUpdateRequest struct {
	// Description:
	Description *string `json:"description,omitempty"`
	// DisplayName:
	DisplayName *string `json:"displayName,omitempty"`
	// Documentation:
	Documentation *string `json:"documentation,omitempty"`
	// ExternalId:
	ExternalId *string `json:"externalId,omitempty"`
	// Icon:
	Icon *string `json:"icon,omitempty"`
	// IsPublic:
	IsPublic *bool `json:"isPublic,omitempty"`
	// Properties:
	Properties map[string]interface{} `json:"properties,omitempty"`
}

// GetDescription returns the Description property
func (m ModelUpdateRequest) GetDescription() *string {
	return m.Description
}

// SetDescription sets the Description property
func (m ModelUpdateRequest) SetDescription(val *string) {
	m.Description = val
}

// GetDisplayName returns the DisplayName property
func (m ModelUpdateRequest) GetDisplayName() *string {
	return m.DisplayName
}

// SetDisplayName sets the DisplayName property
func (m ModelUpdateRequest) SetDisplayName(val *string) {
	m.DisplayName = val
}

// GetDocumentation returns the Documentation property
func (m ModelUpdateRequest) GetDocumentation() *string {
	return m.Documentation
}

// SetDocumentation sets the Documentation property
func (m ModelUpdateRequest) SetDocumentation(val *string) {
	m.Documentation = val
}

// GetExternalId returns the ExternalId property
func (m ModelUpdateRequest) GetExternalId() *string {
	return m.ExternalId
}

// SetExternalId sets the ExternalId property
func (m ModelUpdateRequest) SetExternalId(val *string) {
	m.ExternalId = val
}

// GetIcon returns the Icon property
func (m ModelUpdateRequest) GetIcon() *string {
	return m.Icon
}

// SetIcon sets the Icon property
func (m ModelUpdateRequest) SetIcon(val *string) {
	m.Icon = val
}

// GetIsPublic returns the IsPublic property
func (m ModelUpdateRequest) GetIsPublic() *bool {
	return m.IsPublic
}

// SetIsPublic sets the IsPublic property
func (m ModelUpdateRequest) SetIsPublic(val *bool) {
	m.IsPublic = val
}

// GetProperties returns the Properties property
func (m ModelUpdateRequest) GetProperties() map[string]interface{} {
	return m.Properties
}

// SetProperties sets the Properties property
func (m ModelUpdateRequest) SetProperties(val map[string]interface{}) {
	m.Properties = val
}
