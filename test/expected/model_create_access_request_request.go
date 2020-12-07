// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// CreateAccessRequestRequest is an object. Request for creating a access request
type CreateAccessRequestRequest struct {
	// Comment: Comment why the user would like to access the resource.
	Comment string `json:"comment"`
}

// GetComment returns the Comment property
func (m CreateAccessRequestRequest) GetComment() string {
	return m.Comment
}

// SetComment sets the Comment property
func (m CreateAccessRequestRequest) SetComment(val string) {
	m.Comment = val
}
