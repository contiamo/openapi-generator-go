// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// RelationResponse is an object. A wrapper type that contains the kind of the resource and the resource itself
type RelationResponse struct {
	// Kind: A list of supported resource kinds
	Kind ResourceKind `json:"kind"`
	// Resource:
	Resource interface{} `json:"resource"`
}

// GetKind returns the Kind property
func (m RelationResponse) GetKind() ResourceKind {
	return m.Kind
}

// SetKind sets the Kind property
func (m RelationResponse) SetKind(val ResourceKind) {
	m.Kind = val
}

// GetResource returns the Resource property
func (m RelationResponse) GetResource() interface{} {
	return m.Resource
}

// SetResource sets the Resource property
func (m RelationResponse) SetResource(val interface{}) {
	m.Resource = val
}
