// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// CreateLinkRequest is an object. Request for creating a link
type CreateLinkRequest struct {
	// SourceId:
	SourceId string `json:"sourceId"`
	// TargetId:
	TargetId string `json:"targetId"`
	// Type:
	Type LinkType `json:"type,omitempty"`
}

// GetSourceId returns the SourceId property
func (m CreateLinkRequest) GetSourceId() string {
	return m.SourceId
}

// SetSourceId sets the SourceId property
func (m CreateLinkRequest) SetSourceId(val string) {
	m.SourceId = val
}

// GetTargetId returns the TargetId property
func (m CreateLinkRequest) GetTargetId() string {
	return m.TargetId
}

// SetTargetId sets the TargetId property
func (m CreateLinkRequest) SetTargetId(val string) {
	m.TargetId = val
}

// GetType returns the Type property
func (m CreateLinkRequest) GetType() LinkType {
	return m.Type
}

// SetType sets the Type property
func (m CreateLinkRequest) SetType(val LinkType) {
	m.Type = val
}
