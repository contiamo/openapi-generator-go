// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// BIReportProperties is an object.
type BIReportProperties struct {
	// ExternalEmbeddableURL:
	ExternalEmbeddableURL string `json:"externalEmbeddableURL"`
	// ExternalThumbnailURL:
	ExternalThumbnailURL string `json:"externalThumbnailURL"`
	// ExternalURL:
	ExternalURL string `json:"externalURL"`
	// Type:
	Type BIReportType `json:"type"`
}

// GetExternalEmbeddableURL returns the ExternalEmbeddableURL property
func (m BIReportProperties) GetExternalEmbeddableURL() string {
	return m.ExternalEmbeddableURL
}

// SetExternalEmbeddableURL sets the ExternalEmbeddableURL property
func (m BIReportProperties) SetExternalEmbeddableURL(val string) {
	m.ExternalEmbeddableURL = val
}

// GetExternalThumbnailURL returns the ExternalThumbnailURL property
func (m BIReportProperties) GetExternalThumbnailURL() string {
	return m.ExternalThumbnailURL
}

// SetExternalThumbnailURL sets the ExternalThumbnailURL property
func (m BIReportProperties) SetExternalThumbnailURL(val string) {
	m.ExternalThumbnailURL = val
}

// GetExternalURL returns the ExternalURL property
func (m BIReportProperties) GetExternalURL() string {
	return m.ExternalURL
}

// SetExternalURL sets the ExternalURL property
func (m BIReportProperties) SetExternalURL(val string) {
	m.ExternalURL = val
}

// GetType returns the Type property
func (m BIReportProperties) GetType() BIReportType {
	return m.Type
}

// SetType sets the Type property
func (m BIReportProperties) SetType(val BIReportType) {
	m.Type = val
}
