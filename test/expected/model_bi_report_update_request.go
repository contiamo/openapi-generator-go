// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// BIReportUpdateRequest is an object.
type BIReportUpdateRequest struct {
	// Description:
	Description *string `json:"description,omitempty"`
	// DisplayName:
	DisplayName *string `json:"displayName,omitempty"`
	// Documentation:
	Documentation *string `json:"documentation,omitempty"`
	// ExternalEmbeddableURL:
	ExternalEmbeddableURL *string `json:"externalEmbeddableURL,omitempty"`
	// ExternalId:
	ExternalId *string `json:"externalId,omitempty"`
	// ExternalThumbnailURL:
	ExternalThumbnailURL *string `json:"externalThumbnailURL,omitempty"`
	// ExternalURL:
	ExternalURL *string `json:"externalURL,omitempty"`
	// Icon:
	Icon *string `json:"icon,omitempty"`
	// IsPublic:
	IsPublic *bool `json:"isPublic,omitempty"`
	// Properties:
	Properties map[string]interface{} `json:"properties,omitempty"`
}

// GetDescription returns the Description property
func (m BIReportUpdateRequest) GetDescription() *string {
	return m.Description
}

// SetDescription sets the Description property
func (m BIReportUpdateRequest) SetDescription(val *string) {
	m.Description = val
}

// GetDisplayName returns the DisplayName property
func (m BIReportUpdateRequest) GetDisplayName() *string {
	return m.DisplayName
}

// SetDisplayName sets the DisplayName property
func (m BIReportUpdateRequest) SetDisplayName(val *string) {
	m.DisplayName = val
}

// GetDocumentation returns the Documentation property
func (m BIReportUpdateRequest) GetDocumentation() *string {
	return m.Documentation
}

// SetDocumentation sets the Documentation property
func (m BIReportUpdateRequest) SetDocumentation(val *string) {
	m.Documentation = val
}

// GetExternalEmbeddableURL returns the ExternalEmbeddableURL property
func (m BIReportUpdateRequest) GetExternalEmbeddableURL() *string {
	return m.ExternalEmbeddableURL
}

// SetExternalEmbeddableURL sets the ExternalEmbeddableURL property
func (m BIReportUpdateRequest) SetExternalEmbeddableURL(val *string) {
	m.ExternalEmbeddableURL = val
}

// GetExternalId returns the ExternalId property
func (m BIReportUpdateRequest) GetExternalId() *string {
	return m.ExternalId
}

// SetExternalId sets the ExternalId property
func (m BIReportUpdateRequest) SetExternalId(val *string) {
	m.ExternalId = val
}

// GetExternalThumbnailURL returns the ExternalThumbnailURL property
func (m BIReportUpdateRequest) GetExternalThumbnailURL() *string {
	return m.ExternalThumbnailURL
}

// SetExternalThumbnailURL sets the ExternalThumbnailURL property
func (m BIReportUpdateRequest) SetExternalThumbnailURL(val *string) {
	m.ExternalThumbnailURL = val
}

// GetExternalURL returns the ExternalURL property
func (m BIReportUpdateRequest) GetExternalURL() *string {
	return m.ExternalURL
}

// SetExternalURL sets the ExternalURL property
func (m BIReportUpdateRequest) SetExternalURL(val *string) {
	m.ExternalURL = val
}

// GetIcon returns the Icon property
func (m BIReportUpdateRequest) GetIcon() *string {
	return m.Icon
}

// SetIcon sets the Icon property
func (m BIReportUpdateRequest) SetIcon(val *string) {
	m.Icon = val
}

// GetIsPublic returns the IsPublic property
func (m BIReportUpdateRequest) GetIsPublic() *bool {
	return m.IsPublic
}

// SetIsPublic sets the IsPublic property
func (m BIReportUpdateRequest) SetIsPublic(val *bool) {
	m.IsPublic = val
}

// GetProperties returns the Properties property
func (m BIReportUpdateRequest) GetProperties() map[string]interface{} {
	return m.Properties
}

// SetProperties sets the Properties property
func (m BIReportUpdateRequest) SetProperties(val map[string]interface{}) {
	m.Properties = val
}
