// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// DataSourceUpdateRequest is an object.
type DataSourceUpdateRequest struct {
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
func (m DataSourceUpdateRequest) GetDescription() *string {
	return m.Description
}

// SetDescription sets the Description property
func (m DataSourceUpdateRequest) SetDescription(val *string) {
	m.Description = val
}

// GetDisplayName returns the DisplayName property
func (m DataSourceUpdateRequest) GetDisplayName() *string {
	return m.DisplayName
}

// SetDisplayName sets the DisplayName property
func (m DataSourceUpdateRequest) SetDisplayName(val *string) {
	m.DisplayName = val
}

// GetDocumentation returns the Documentation property
func (m DataSourceUpdateRequest) GetDocumentation() *string {
	return m.Documentation
}

// SetDocumentation sets the Documentation property
func (m DataSourceUpdateRequest) SetDocumentation(val *string) {
	m.Documentation = val
}

// GetExternalId returns the ExternalId property
func (m DataSourceUpdateRequest) GetExternalId() *string {
	return m.ExternalId
}

// SetExternalId sets the ExternalId property
func (m DataSourceUpdateRequest) SetExternalId(val *string) {
	m.ExternalId = val
}

// GetIcon returns the Icon property
func (m DataSourceUpdateRequest) GetIcon() *string {
	return m.Icon
}

// SetIcon sets the Icon property
func (m DataSourceUpdateRequest) SetIcon(val *string) {
	m.Icon = val
}

// GetIsPublic returns the IsPublic property
func (m DataSourceUpdateRequest) GetIsPublic() *bool {
	return m.IsPublic
}

// SetIsPublic sets the IsPublic property
func (m DataSourceUpdateRequest) SetIsPublic(val *bool) {
	m.IsPublic = val
}

// GetProperties returns the Properties property
func (m DataSourceUpdateRequest) GetProperties() map[string]interface{} {
	return m.Properties
}

// SetProperties sets the Properties property
func (m DataSourceUpdateRequest) SetProperties(val map[string]interface{}) {
	m.Properties = val
}
