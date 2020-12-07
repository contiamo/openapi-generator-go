// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// FileDiscoveryResponse is an object.
type FileDiscoveryResponse struct {
	// Options: FileOptions determine how the file will be opened and parsed
	Options FileOptions `json:"options"`
	// Preview: FileSchemaPreview provides a preview of how a file will be parsed with a given schema, representing an array of rows of a table
	Preview [][]SchemaPreviewCell `json:"preview"`
	// Schema: Schema of a file or of the corresponding managed table which will be created
	Schema ManagedTableSchema `json:"schema"`
	// SchemaErrors:
	SchemaErrors *interface{} `json:"schemaErrors"`
	// TargetSchema:
	TargetSchema *interface{} `json:"targetSchema"`
}

// GetOptions returns the Options property
func (m FileDiscoveryResponse) GetOptions() FileOptions {
	return m.Options
}

// SetOptions sets the Options property
func (m FileDiscoveryResponse) SetOptions(val FileOptions) {
	m.Options = val
}

// GetPreview returns the Preview property
func (m FileDiscoveryResponse) GetPreview() [][]SchemaPreviewCell {
	return m.Preview
}

// SetPreview sets the Preview property
func (m FileDiscoveryResponse) SetPreview(val [][]SchemaPreviewCell) {
	m.Preview = val
}

// GetSchema returns the Schema property
func (m FileDiscoveryResponse) GetSchema() ManagedTableSchema {
	return m.Schema
}

// SetSchema sets the Schema property
func (m FileDiscoveryResponse) SetSchema(val ManagedTableSchema) {
	m.Schema = val
}

// GetSchemaErrors returns the SchemaErrors property
func (m FileDiscoveryResponse) GetSchemaErrors() *interface{} {
	return m.SchemaErrors
}

// SetSchemaErrors sets the SchemaErrors property
func (m FileDiscoveryResponse) SetSchemaErrors(val *interface{}) {
	m.SchemaErrors = val
}

// GetTargetSchema returns the TargetSchema property
func (m FileDiscoveryResponse) GetTargetSchema() *interface{} {
	return m.TargetSchema
}

// SetTargetSchema sets the TargetSchema property
func (m FileDiscoveryResponse) SetTargetSchema(val *interface{}) {
	m.TargetSchema = val
}
