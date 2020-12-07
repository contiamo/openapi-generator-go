// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// FilePreviewResponse is an object.
type FilePreviewResponse struct {
	// Preview: FileSchemaPreview provides a preview of how a file will be parsed with a given schema, representing an array of rows of a table
	Preview [][]SchemaPreviewCell `json:"preview"`
}

// GetPreview returns the Preview property
func (m FilePreviewResponse) GetPreview() [][]SchemaPreviewCell {
	return m.Preview
}

// SetPreview sets the Preview property
func (m FilePreviewResponse) SetPreview(val [][]SchemaPreviewCell) {
	m.Preview = val
}
