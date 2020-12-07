// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// UseCaseUpdateRequest is an object.
type UseCaseUpdateRequest struct {
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
func (m UseCaseUpdateRequest) GetDescription() *string {
	return m.Description
}

// SetDescription sets the Description property
func (m UseCaseUpdateRequest) SetDescription(val *string) {
	m.Description = val
}

// GetDisplayName returns the DisplayName property
func (m UseCaseUpdateRequest) GetDisplayName() *string {
	return m.DisplayName
}

// SetDisplayName sets the DisplayName property
func (m UseCaseUpdateRequest) SetDisplayName(val *string) {
	m.DisplayName = val
}

// GetDocumentation returns the Documentation property
func (m UseCaseUpdateRequest) GetDocumentation() *string {
	return m.Documentation
}

// SetDocumentation sets the Documentation property
func (m UseCaseUpdateRequest) SetDocumentation(val *string) {
	m.Documentation = val
}

// GetExternalId returns the ExternalId property
func (m UseCaseUpdateRequest) GetExternalId() *string {
	return m.ExternalId
}

// SetExternalId sets the ExternalId property
func (m UseCaseUpdateRequest) SetExternalId(val *string) {
	m.ExternalId = val
}

// GetIcon returns the Icon property
func (m UseCaseUpdateRequest) GetIcon() *string {
	return m.Icon
}

// SetIcon sets the Icon property
func (m UseCaseUpdateRequest) SetIcon(val *string) {
	m.Icon = val
}

// GetIsPublic returns the IsPublic property
func (m UseCaseUpdateRequest) GetIsPublic() *bool {
	return m.IsPublic
}

// SetIsPublic sets the IsPublic property
func (m UseCaseUpdateRequest) SetIsPublic(val *bool) {
	m.IsPublic = val
}

// GetProperties returns the Properties property
func (m UseCaseUpdateRequest) GetProperties() map[string]interface{} {
	return m.Properties
}

// SetProperties sets the Properties property
func (m UseCaseUpdateRequest) SetProperties(val map[string]interface{}) {
	m.Properties = val
}
