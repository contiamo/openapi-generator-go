// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// BIReportCreateRequest is an object.
type BIReportCreateRequest struct {
	// Description:
	Description string `json:"description,omitempty"`
	// DisplayName:
	DisplayName string `json:"displayName"`
	// ExternalEmbeddableURL:
	ExternalEmbeddableURL string `json:"externalEmbeddableURL,omitempty"`
	// ExternalId:
	ExternalId *string `json:"externalId,omitempty"`
	// ExternalThumbnailURL:
	ExternalThumbnailURL string `json:"externalThumbnailURL,omitempty"`
	// ExternalURL:
	ExternalURL string `json:"externalURL,omitempty"`
	// IsPublic:
	IsPublic bool `json:"isPublic"`
	// Type:
	Type BIReportType `json:"type,omitempty"`
}

// GetDescription returns the Description property
func (m BIReportCreateRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m BIReportCreateRequest) SetDescription(val string) {
	m.Description = val
}

// GetDisplayName returns the DisplayName property
func (m BIReportCreateRequest) GetDisplayName() string {
	return m.DisplayName
}

// SetDisplayName sets the DisplayName property
func (m BIReportCreateRequest) SetDisplayName(val string) {
	m.DisplayName = val
}

// GetExternalEmbeddableURL returns the ExternalEmbeddableURL property
func (m BIReportCreateRequest) GetExternalEmbeddableURL() string {
	return m.ExternalEmbeddableURL
}

// SetExternalEmbeddableURL sets the ExternalEmbeddableURL property
func (m BIReportCreateRequest) SetExternalEmbeddableURL(val string) {
	m.ExternalEmbeddableURL = val
}

// GetExternalId returns the ExternalId property
func (m BIReportCreateRequest) GetExternalId() *string {
	return m.ExternalId
}

// SetExternalId sets the ExternalId property
func (m BIReportCreateRequest) SetExternalId(val *string) {
	m.ExternalId = val
}

// GetExternalThumbnailURL returns the ExternalThumbnailURL property
func (m BIReportCreateRequest) GetExternalThumbnailURL() string {
	return m.ExternalThumbnailURL
}

// SetExternalThumbnailURL sets the ExternalThumbnailURL property
func (m BIReportCreateRequest) SetExternalThumbnailURL(val string) {
	m.ExternalThumbnailURL = val
}

// GetExternalURL returns the ExternalURL property
func (m BIReportCreateRequest) GetExternalURL() string {
	return m.ExternalURL
}

// SetExternalURL sets the ExternalURL property
func (m BIReportCreateRequest) SetExternalURL(val string) {
	m.ExternalURL = val
}

// GetIsPublic returns the IsPublic property
func (m BIReportCreateRequest) GetIsPublic() bool {
	return m.IsPublic
}

// SetIsPublic sets the IsPublic property
func (m BIReportCreateRequest) SetIsPublic(val bool) {
	m.IsPublic = val
}

// GetType returns the Type property
func (m BIReportCreateRequest) GetType() BIReportType {
	return m.Type
}

// SetType sets the Type property
func (m BIReportCreateRequest) SetType(val BIReportType) {
	m.Type = val
}
