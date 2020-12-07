// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// UploadResponse is an object.
type UploadResponse struct {
	// CollectionId:
	CollectionId string `json:"collectionId"`
	// FileId:
	FileId string `json:"fileId"`
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
	// UploadedBytes:
	UploadedBytes int32 `json:"uploadedBytes"`
}

// GetCollectionId returns the CollectionId property
func (m UploadResponse) GetCollectionId() string {
	return m.CollectionId
}

// SetCollectionId sets the CollectionId property
func (m UploadResponse) SetCollectionId(val string) {
	m.CollectionId = val
}

// GetFileId returns the FileId property
func (m UploadResponse) GetFileId() string {
	return m.FileId
}

// SetFileId sets the FileId property
func (m UploadResponse) SetFileId(val string) {
	m.FileId = val
}

// GetOptions returns the Options property
func (m UploadResponse) GetOptions() FileOptions {
	return m.Options
}

// SetOptions sets the Options property
func (m UploadResponse) SetOptions(val FileOptions) {
	m.Options = val
}

// GetPreview returns the Preview property
func (m UploadResponse) GetPreview() [][]SchemaPreviewCell {
	return m.Preview
}

// SetPreview sets the Preview property
func (m UploadResponse) SetPreview(val [][]SchemaPreviewCell) {
	m.Preview = val
}

// GetSchema returns the Schema property
func (m UploadResponse) GetSchema() ManagedTableSchema {
	return m.Schema
}

// SetSchema sets the Schema property
func (m UploadResponse) SetSchema(val ManagedTableSchema) {
	m.Schema = val
}

// GetSchemaErrors returns the SchemaErrors property
func (m UploadResponse) GetSchemaErrors() *interface{} {
	return m.SchemaErrors
}

// SetSchemaErrors sets the SchemaErrors property
func (m UploadResponse) SetSchemaErrors(val *interface{}) {
	m.SchemaErrors = val
}

// GetTargetSchema returns the TargetSchema property
func (m UploadResponse) GetTargetSchema() *interface{} {
	return m.TargetSchema
}

// SetTargetSchema sets the TargetSchema property
func (m UploadResponse) SetTargetSchema(val *interface{}) {
	m.TargetSchema = val
}

// GetUploadedBytes returns the UploadedBytes property
func (m UploadResponse) GetUploadedBytes() int32 {
	return m.UploadedBytes
}

// SetUploadedBytes sets the UploadedBytes property
func (m UploadResponse) SetUploadedBytes(val int32) {
	m.UploadedBytes = val
}
