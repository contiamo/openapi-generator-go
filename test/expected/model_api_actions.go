// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// APIActions is an object. List of actions allowed to perform on the API
type APIActions struct {
	// Delete: If true the user is allowed to delete the resource
	Delete bool `json:"delete"`
	// EditMetadata: If true the user is allowed to edit the metadata of the resource
	EditMetadata bool `json:"editMetadata"`
	// Grant: If true the user is allowed to grant permissions for the resource to other users
	Grant bool `json:"grant"`
	// Read: If true the user is allowed to read the resource and its connections/links
	Read bool `json:"read"`
	// Role: Aggregated role name.  This can be used to indicate the total role level that is represented by the users available
	// actions.
	//
	// The value `none` indicates that the user has no direct permissions on this resource. In general, this will only be returned if the
	// resource is `public`, indicating that the user can read the resource metadata, but can not otherwise perform any other actions.
	//
	// The value `admin` indicates the user is a project admin, and therefore is allowed all available actions, overriding any other specific
	// permissions that may have been directly granted to the user or a group they are in.
	//
	// The other allowed values correspond to a specific `Role`. This indicates that the user or a group they are in was directly granted this role.
	// In this case, the API returns the role name with the highest access levels.
	Role string `json:"role"`
}

// GetDelete returns the Delete property
func (m APIActions) GetDelete() bool {
	return m.Delete
}

// SetDelete sets the Delete property
func (m APIActions) SetDelete(val bool) {
	m.Delete = val
}

// GetEditMetadata returns the EditMetadata property
func (m APIActions) GetEditMetadata() bool {
	return m.EditMetadata
}

// SetEditMetadata sets the EditMetadata property
func (m APIActions) SetEditMetadata(val bool) {
	m.EditMetadata = val
}

// GetGrant returns the Grant property
func (m APIActions) GetGrant() bool {
	return m.Grant
}

// SetGrant sets the Grant property
func (m APIActions) SetGrant(val bool) {
	m.Grant = val
}

// GetRead returns the Read property
func (m APIActions) GetRead() bool {
	return m.Read
}

// SetRead sets the Read property
func (m APIActions) SetRead(val bool) {
	m.Read = val
}

// GetRole returns the Role property
func (m APIActions) GetRole() string {
	return m.Role
}

// SetRole sets the Role property
func (m APIActions) SetRole(val string) {
	m.Role = val
}
