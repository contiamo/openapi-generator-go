// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// GlossaryTree is an object.
type GlossaryTree struct {
	// Children:
	Children []GlossaryTree `json:"children"`
	// Resource:
	Resource GlossaryItemResponse `json:"resource"`
}

// GetChildren returns the Children property
func (m GlossaryTree) GetChildren() []GlossaryTree {
	return m.Children
}

// SetChildren sets the Children property
func (m GlossaryTree) SetChildren(val []GlossaryTree) {
	m.Children = val
}

// GetResource returns the Resource property
func (m GlossaryTree) GetResource() GlossaryItemResponse {
	return m.Resource
}

// SetResource sets the Resource property
func (m GlossaryTree) SetResource(val GlossaryItemResponse) {
	m.Resource = val
}
