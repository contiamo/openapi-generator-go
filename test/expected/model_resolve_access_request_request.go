// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// ResolveAccessRequestRequest is an object. Request data when resolving a access request
// If the role is empty it defaults to `query` for data sources and `steward` for usecases.
type ResolveAccessRequestRequest struct {
	// Comment: Comment explaining why the request was granted/rejected
	Comment string `json:"comment,omitempty"`
	// Role: A valid resource role identifier
	Role Role `json:"role,omitempty"`
	// State: state of the request, either 'pending', 'granted' or 'rejected'
	State AccessRequestState `json:"state"`
}

// GetComment returns the Comment property
func (m ResolveAccessRequestRequest) GetComment() string {
	return m.Comment
}

// SetComment sets the Comment property
func (m ResolveAccessRequestRequest) SetComment(val string) {
	m.Comment = val
}

// GetRole returns the Role property
func (m ResolveAccessRequestRequest) GetRole() Role {
	return m.Role
}

// SetRole sets the Role property
func (m ResolveAccessRequestRequest) SetRole(val Role) {
	m.Role = val
}

// GetState returns the State property
func (m ResolveAccessRequestRequest) GetState() AccessRequestState {
	return m.State
}

// SetState sets the State property
func (m ResolveAccessRequestRequest) SetState(val AccessRequestState) {
	m.State = val
}
