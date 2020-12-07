// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// GlossaryItemProperties is an object.
type GlossaryItemProperties struct {
	// DisplayNamePath: Array of display names representing the ancestor chain of this glossary item, this should be used to create the "B2B → Customer → Address" label in the front-end
	DisplayNamePath []string `json:"displayNamePath"`
	// GlossaryParentId:
	GlossaryParentId *string `json:"glossaryParentId"`
	// Path: Array of ids representing the ancestor chain ot this glossary item
	Path []string `json:"path"`
}

// GetDisplayNamePath returns the DisplayNamePath property
func (m GlossaryItemProperties) GetDisplayNamePath() []string {
	return m.DisplayNamePath
}

// SetDisplayNamePath sets the DisplayNamePath property
func (m GlossaryItemProperties) SetDisplayNamePath(val []string) {
	m.DisplayNamePath = val
}

// GetGlossaryParentId returns the GlossaryParentId property
func (m GlossaryItemProperties) GetGlossaryParentId() *string {
	return m.GlossaryParentId
}

// SetGlossaryParentId sets the GlossaryParentId property
func (m GlossaryItemProperties) SetGlossaryParentId(val *string) {
	m.GlossaryParentId = val
}

// GetPath returns the Path property
func (m GlossaryItemProperties) GetPath() []string {
	return m.Path
}

// SetPath sets the Path property
func (m GlossaryItemProperties) SetPath(val []string) {
	m.Path = val
}
