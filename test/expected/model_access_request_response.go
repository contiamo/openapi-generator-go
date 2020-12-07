// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

import (
	"time"
)

// AccessRequestResponse is an object. Object representing that one user requested access on a specific resource
type AccessRequestResponse struct {
	// Comment: The reason why the access was requested.
	Comment string `json:"comment"`
	// CreatedAt: The timestamp of the creation of this request.
	CreatedAt time.Time `json:"createdAt"`
	// Id:
	Id string `json:"id"`
	// RequesterId: UUID of the user requesting access
	RequesterId string `json:"requesterId"`
	// ResolveComment: The reason why the access was rejected/granted.
	ResolveComment string `json:"resolveComment"`
	// ResolvedAt: The timestamp when this request has been granted.
	ResolvedAt *time.Time `json:"resolvedAt"`
	// ResourceId: UUID of the resource. Since resources represent all other entities like data sources, tables and use cases, resource endpoints can be used for those entities as well.
	ResourceId string `json:"resourceId"`
	// State: state of the request, either 'pending', 'granted' or 'rejected'
	State AccessRequestState `json:"state"`
}

// GetComment returns the Comment property
func (m AccessRequestResponse) GetComment() string {
	return m.Comment
}

// SetComment sets the Comment property
func (m AccessRequestResponse) SetComment(val string) {
	m.Comment = val
}

// GetCreatedAt returns the CreatedAt property
func (m AccessRequestResponse) GetCreatedAt() time.Time {
	return m.CreatedAt
}

// SetCreatedAt sets the CreatedAt property
func (m AccessRequestResponse) SetCreatedAt(val time.Time) {
	m.CreatedAt = val
}

// GetId returns the Id property
func (m AccessRequestResponse) GetId() string {
	return m.Id
}

// SetId sets the Id property
func (m AccessRequestResponse) SetId(val string) {
	m.Id = val
}

// GetRequesterId returns the RequesterId property
func (m AccessRequestResponse) GetRequesterId() string {
	return m.RequesterId
}

// SetRequesterId sets the RequesterId property
func (m AccessRequestResponse) SetRequesterId(val string) {
	m.RequesterId = val
}

// GetResolveComment returns the ResolveComment property
func (m AccessRequestResponse) GetResolveComment() string {
	return m.ResolveComment
}

// SetResolveComment sets the ResolveComment property
func (m AccessRequestResponse) SetResolveComment(val string) {
	m.ResolveComment = val
}

// GetResolvedAt returns the ResolvedAt property
func (m AccessRequestResponse) GetResolvedAt() *time.Time {
	return m.ResolvedAt
}

// SetResolvedAt sets the ResolvedAt property
func (m AccessRequestResponse) SetResolvedAt(val *time.Time) {
	m.ResolvedAt = val
}

// GetResourceId returns the ResourceId property
func (m AccessRequestResponse) GetResourceId() string {
	return m.ResourceId
}

// SetResourceId sets the ResourceId property
func (m AccessRequestResponse) SetResourceId(val string) {
	m.ResourceId = val
}

// GetState returns the State property
func (m AccessRequestResponse) GetState() AccessRequestState {
	return m.State
}

// SetState sets the State property
func (m AccessRequestResponse) SetState(val AccessRequestState) {
	m.State = val
}
