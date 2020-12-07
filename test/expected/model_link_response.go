// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// LinkResponse is an object.
type LinkResponse struct {
	// LinkId:
	LinkId string `json:"linkId"`
	// SourceDisplayNamePath:
	SourceDisplayNamePath []string `json:"sourceDisplayNamePath"`
	// SourceId:
	SourceId string `json:"sourceId"`
	// SourceKind: A list of supported resource kinds
	SourceKind ResourceKind `json:"sourceKind"`
	// TargetDisplayNamePath:
	TargetDisplayNamePath []string `json:"targetDisplayNamePath"`
	// TargetId:
	TargetId string `json:"targetId"`
	// TargetKind: A list of supported resource kinds
	TargetKind ResourceKind `json:"targetKind"`
	// Type:
	Type LinkType `json:"type"`
}

// GetLinkId returns the LinkId property
func (m LinkResponse) GetLinkId() string {
	return m.LinkId
}

// SetLinkId sets the LinkId property
func (m LinkResponse) SetLinkId(val string) {
	m.LinkId = val
}

// GetSourceDisplayNamePath returns the SourceDisplayNamePath property
func (m LinkResponse) GetSourceDisplayNamePath() []string {
	return m.SourceDisplayNamePath
}

// SetSourceDisplayNamePath sets the SourceDisplayNamePath property
func (m LinkResponse) SetSourceDisplayNamePath(val []string) {
	m.SourceDisplayNamePath = val
}

// GetSourceId returns the SourceId property
func (m LinkResponse) GetSourceId() string {
	return m.SourceId
}

// SetSourceId sets the SourceId property
func (m LinkResponse) SetSourceId(val string) {
	m.SourceId = val
}

// GetSourceKind returns the SourceKind property
func (m LinkResponse) GetSourceKind() ResourceKind {
	return m.SourceKind
}

// SetSourceKind sets the SourceKind property
func (m LinkResponse) SetSourceKind(val ResourceKind) {
	m.SourceKind = val
}

// GetTargetDisplayNamePath returns the TargetDisplayNamePath property
func (m LinkResponse) GetTargetDisplayNamePath() []string {
	return m.TargetDisplayNamePath
}

// SetTargetDisplayNamePath sets the TargetDisplayNamePath property
func (m LinkResponse) SetTargetDisplayNamePath(val []string) {
	m.TargetDisplayNamePath = val
}

// GetTargetId returns the TargetId property
func (m LinkResponse) GetTargetId() string {
	return m.TargetId
}

// SetTargetId sets the TargetId property
func (m LinkResponse) SetTargetId(val string) {
	m.TargetId = val
}

// GetTargetKind returns the TargetKind property
func (m LinkResponse) GetTargetKind() ResourceKind {
	return m.TargetKind
}

// SetTargetKind sets the TargetKind property
func (m LinkResponse) SetTargetKind(val ResourceKind) {
	m.TargetKind = val
}

// GetType returns the Type property
func (m LinkResponse) GetType() LinkType {
	return m.Type
}

// SetType sets the Type property
func (m LinkResponse) SetType(val LinkType) {
	m.Type = val
}
