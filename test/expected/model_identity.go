// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// Identity is an object. Short description of the identity accessing the system.
type Identity struct {
	// Email: Email associated with the user. It's blank for service accounts.
	Email string `json:"email,omitempty"`
	// Id: UUID of the user or the service account.
	Id string `json:"id"`
	// Name: Full name of the user or a name of the service account.
	Name string `json:"name"`
}

// GetEmail returns the Email property
func (m Identity) GetEmail() string {
	return m.Email
}

// SetEmail sets the Email property
func (m Identity) SetEmail(val string) {
	m.Email = val
}

// GetId returns the Id property
func (m Identity) GetId() string {
	return m.Id
}

// SetId sets the Id property
func (m Identity) SetId(val string) {
	m.Id = val
}

// GetName returns the Name property
func (m Identity) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m Identity) SetName(val string) {
	m.Name = val
}
