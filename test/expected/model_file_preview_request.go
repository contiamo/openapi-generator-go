// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// FilePreviewRequest is an object.
type FilePreviewRequest struct {
	// Options: FileOptions determine how the file will be opened and parsed
	Options FileOptions `json:"options"`
	// Schema: Schema of a file or of the corresponding managed table which will be created
	Schema ManagedTableSchema `json:"schema"`
}

// GetOptions returns the Options property
func (m FilePreviewRequest) GetOptions() FileOptions {
	return m.Options
}

// SetOptions sets the Options property
func (m FilePreviewRequest) SetOptions(val FileOptions) {
	m.Options = val
}

// GetSchema returns the Schema property
func (m FilePreviewRequest) GetSchema() ManagedTableSchema {
	return m.Schema
}

// SetSchema sets the Schema property
func (m FilePreviewRequest) SetSchema(val ManagedTableSchema) {
	m.Schema = val
}
