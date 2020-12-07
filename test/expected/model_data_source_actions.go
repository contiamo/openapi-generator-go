// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// DataSourceActions is an object. List of actions allowed to perform on the data source
type DataSourceActions struct {
	// CreateTable: If true the user is allowed to create tables in the data source
	CreateTable bool `json:"createTable"`
	// Delete: If true the user is allowed to
	//
	//   1) delete the data source and its tables
	//
	//   2) Truncate a table in a managed data source
	Delete bool `json:"delete"`
	// Edit: If true the user is allowed to:
	//
	//   1) edit the data source and its tables
	//
	//   2) edit materialization preferences for a view
	Edit bool `json:"edit"`
	// EditMetadata: If true the user is allowed to edit the metadata of the data source and its tables
	EditMetadata bool `json:"editMetadata"`
	// Grant: If true the user is allowed to grant permissions for the data source to other users
	Grant bool `json:"grant"`
	// Profile: If true the user is allowed to read and trigger profiling sessions
	Profile bool `json:"profile"`
	// Query: If true the user is allowed to query the data source
	Query bool `json:"query"`
	// Read: If true the user is allowed to:
	//
	//   1) read the data source and its tables
	//
	//   2) read the list of import tasks for a managed data source
	Read bool `json:"read"`
	// Refresh: If true the user is allowed to refresh the data source's schema
	Refresh bool `json:"refresh"`
	// Role: aggregated role name.  This can be used to indicate the total role level that is represented by the users available
	// actions.
	//
	// The value `none` indicates that the user has no direct permissions on the data source. In general, this will only be returned if the data
	// source is `public`, indicating that the user can read the data source metadata, but can not otherwise perform any other actions.
	//
	// The value `admin` indicates the user is a project admin, and therefore is allowed all available actions, overriding any other specific
	// permissions that may have been directly granted to the user or a group they are in.
	//
	// The other allowed values correspond to a specific `Role`. This indicates that the user or a group they are in was directly granted this role.
	// In this case, the API returns the role name with the highest access levels.
	Role string `json:"role"`
	// Upload: If true the user is allowed to:
	//
	//   1) upload data to the managed data source
	//
	//   2) cancel an import task for a manged data source
	Upload bool `json:"upload"`
}

// GetCreateTable returns the CreateTable property
func (m DataSourceActions) GetCreateTable() bool {
	return m.CreateTable
}

// SetCreateTable sets the CreateTable property
func (m DataSourceActions) SetCreateTable(val bool) {
	m.CreateTable = val
}

// GetDelete returns the Delete property
func (m DataSourceActions) GetDelete() bool {
	return m.Delete
}

// SetDelete sets the Delete property
func (m DataSourceActions) SetDelete(val bool) {
	m.Delete = val
}

// GetEdit returns the Edit property
func (m DataSourceActions) GetEdit() bool {
	return m.Edit
}

// SetEdit sets the Edit property
func (m DataSourceActions) SetEdit(val bool) {
	m.Edit = val
}

// GetEditMetadata returns the EditMetadata property
func (m DataSourceActions) GetEditMetadata() bool {
	return m.EditMetadata
}

// SetEditMetadata sets the EditMetadata property
func (m DataSourceActions) SetEditMetadata(val bool) {
	m.EditMetadata = val
}

// GetGrant returns the Grant property
func (m DataSourceActions) GetGrant() bool {
	return m.Grant
}

// SetGrant sets the Grant property
func (m DataSourceActions) SetGrant(val bool) {
	m.Grant = val
}

// GetProfile returns the Profile property
func (m DataSourceActions) GetProfile() bool {
	return m.Profile
}

// SetProfile sets the Profile property
func (m DataSourceActions) SetProfile(val bool) {
	m.Profile = val
}

// GetQuery returns the Query property
func (m DataSourceActions) GetQuery() bool {
	return m.Query
}

// SetQuery sets the Query property
func (m DataSourceActions) SetQuery(val bool) {
	m.Query = val
}

// GetRead returns the Read property
func (m DataSourceActions) GetRead() bool {
	return m.Read
}

// SetRead sets the Read property
func (m DataSourceActions) SetRead(val bool) {
	m.Read = val
}

// GetRefresh returns the Refresh property
func (m DataSourceActions) GetRefresh() bool {
	return m.Refresh
}

// SetRefresh sets the Refresh property
func (m DataSourceActions) SetRefresh(val bool) {
	m.Refresh = val
}

// GetRole returns the Role property
func (m DataSourceActions) GetRole() string {
	return m.Role
}

// SetRole sets the Role property
func (m DataSourceActions) SetRole(val string) {
	m.Role = val
}

// GetUpload returns the Upload property
func (m DataSourceActions) GetUpload() bool {
	return m.Upload
}

// SetUpload sets the Upload property
func (m DataSourceActions) SetUpload(val bool) {
	m.Upload = val
}
