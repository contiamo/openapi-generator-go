// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// PipelineCreateRequest is an object.
type PipelineCreateRequest struct {
	// Description:
	Description string `json:"description,omitempty"`
	// DisplayName:
	DisplayName string `json:"displayName"`
	// ExternalId:
	ExternalId *string `json:"externalId,omitempty"`
	// IsPublic:
	IsPublic bool `json:"isPublic"`
	// Type:
	Type PipelineType `json:"type,omitempty"`
}

// GetDescription returns the Description property
func (m PipelineCreateRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m PipelineCreateRequest) SetDescription(val string) {
	m.Description = val
}

// GetDisplayName returns the DisplayName property
func (m PipelineCreateRequest) GetDisplayName() string {
	return m.DisplayName
}

// SetDisplayName sets the DisplayName property
func (m PipelineCreateRequest) SetDisplayName(val string) {
	m.DisplayName = val
}

// GetExternalId returns the ExternalId property
func (m PipelineCreateRequest) GetExternalId() *string {
	return m.ExternalId
}

// SetExternalId sets the ExternalId property
func (m PipelineCreateRequest) SetExternalId(val *string) {
	m.ExternalId = val
}

// GetIsPublic returns the IsPublic property
func (m PipelineCreateRequest) GetIsPublic() bool {
	return m.IsPublic
}

// SetIsPublic sets the IsPublic property
func (m PipelineCreateRequest) SetIsPublic(val bool) {
	m.IsPublic = val
}

// GetType returns the Type property
func (m PipelineCreateRequest) GetType() PipelineType {
	return m.Type
}

// SetType sets the Type property
func (m PipelineCreateRequest) SetType(val PipelineType) {
	m.Type = val
}
