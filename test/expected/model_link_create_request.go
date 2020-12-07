// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// LinkCreateRequest is an object.
type LinkCreateRequest struct {
	// SourceId:
	SourceId string `json:"sourceId"`
	// TargetId:
	TargetId string `json:"targetId"`
	// Type:
	Type LinkType `json:"type"`
}

// GetSourceId returns the SourceId property
func (m LinkCreateRequest) GetSourceId() string {
	return m.SourceId
}

// SetSourceId sets the SourceId property
func (m LinkCreateRequest) SetSourceId(val string) {
	m.SourceId = val
}

// GetTargetId returns the TargetId property
func (m LinkCreateRequest) GetTargetId() string {
	return m.TargetId
}

// SetTargetId sets the TargetId property
func (m LinkCreateRequest) SetTargetId(val string) {
	m.TargetId = val
}

// GetType returns the Type property
func (m LinkCreateRequest) GetType() LinkType {
	return m.Type
}

// SetType sets the Type property
func (m LinkCreateRequest) SetType(val LinkType) {
	m.Type = val
}
