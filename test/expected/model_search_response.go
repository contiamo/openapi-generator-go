// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// SearchResponse is an object.
type SearchResponse struct {
	// Children: List of search responses that holds matches that are direct children of this resource
	Children []SearchResponse `json:"children"`
	// Kind: A list of supported resource kinds
	Kind ResourceKind `json:"kind"`
	// Meta:
	Meta struct {
		DescriptionHighlight string  `json:"descriptionHighlight,omitempty"`
		Rank                 float32 `json:"rank,omitempty"`
	} `json:"meta"`
	// Resource:
	Resource interface{} `json:"resource"`
}

// GetChildren returns the Children property
func (m SearchResponse) GetChildren() []SearchResponse {
	return m.Children
}

// SetChildren sets the Children property
func (m SearchResponse) SetChildren(val []SearchResponse) {
	m.Children = val
}

// GetKind returns the Kind property
func (m SearchResponse) GetKind() ResourceKind {
	return m.Kind
}

// SetKind sets the Kind property
func (m SearchResponse) SetKind(val ResourceKind) {
	m.Kind = val
}

// GetMeta returns the Meta property
func (m SearchResponse) GetMeta() struct {
	DescriptionHighlight string  `json:"descriptionHighlight,omitempty"`
	Rank                 float32 `json:"rank,omitempty"`
} {
	return m.Meta
}

// SetMeta sets the Meta property
func (m SearchResponse) SetMeta(val struct {
	DescriptionHighlight string  `json:"descriptionHighlight,omitempty"`
	Rank                 float32 `json:"rank,omitempty"`
}) {
	m.Meta = val
}

// GetResource returns the Resource property
func (m SearchResponse) GetResource() interface{} {
	return m.Resource
}

// SetResource sets the Resource property
func (m SearchResponse) SetResource(val interface{}) {
	m.Resource = val
}
