// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// GlossaryItemActions is an object. List of actions allowed to perform on the glossary item
type GlossaryItemActions struct {
	// CreateChild: If true the user is allowed to create child glossary items
	CreateChild bool `json:"createChild,omitempty"`
	// Delete: If true the user is allowed to delete the glossary item
	Delete bool `json:"delete"`
	// EditMetadata: If true the user is allowed to edit the metadata of the glossary item
	EditMetadata bool `json:"editMetadata"`
	// Grant: If true the user is allowed to grant permissions for the glossary item to other users
	Grant bool `json:"grant"`
	// Read: If true the user is allowed to read the glossary item and its connections/links
	Read bool `json:"read"`
	// Role: Aggregated role name. This can be used to indicate the total role level that is represented by the users available
	// actions.
	//
	// The value `none` indicates that the user has no direct permissions on the glossary item. In general, this will only be returned if the
	// resource is `public`, indicating that the user can read the resource metadata, but can not otherwise perform any other actions.
	//
	// The value `admin` indicates the user is a project admin, and therefore is allowed all available actions, overriding any other specific
	// permissions that may have been directly granted to the user or a group they are in.
	//
	// The other allowed values correspond to a specific `Role`. This indicates that the user or a group they are in was directly granted this role.
	// In this case, the API returns the role name with the highest access levels.
	Role string `json:"role"`
}

// GetCreateChild returns the CreateChild property
func (m GlossaryItemActions) GetCreateChild() bool {
	return m.CreateChild
}

// SetCreateChild sets the CreateChild property
func (m GlossaryItemActions) SetCreateChild(val bool) {
	m.CreateChild = val
}

// GetDelete returns the Delete property
func (m GlossaryItemActions) GetDelete() bool {
	return m.Delete
}

// SetDelete sets the Delete property
func (m GlossaryItemActions) SetDelete(val bool) {
	m.Delete = val
}

// GetEditMetadata returns the EditMetadata property
func (m GlossaryItemActions) GetEditMetadata() bool {
	return m.EditMetadata
}

// SetEditMetadata sets the EditMetadata property
func (m GlossaryItemActions) SetEditMetadata(val bool) {
	m.EditMetadata = val
}

// GetGrant returns the Grant property
func (m GlossaryItemActions) GetGrant() bool {
	return m.Grant
}

// SetGrant sets the Grant property
func (m GlossaryItemActions) SetGrant(val bool) {
	m.Grant = val
}

// GetRead returns the Read property
func (m GlossaryItemActions) GetRead() bool {
	return m.Read
}

// SetRead sets the Read property
func (m GlossaryItemActions) SetRead(val bool) {
	m.Read = val
}

// GetRole returns the Role property
func (m GlossaryItemActions) GetRole() string {
	return m.Role
}

// SetRole sets the Role property
func (m GlossaryItemActions) SetRole(val string) {
	m.Role = val
}
