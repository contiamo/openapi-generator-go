// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// FileIngestRequest is an object.
type FileIngestRequest struct {
	// ColumnMapping: Defines how to map the columns from a file to be ingested
	// to the columns of the target table.
	// This allows to reorder and/or drop certain columns of the input data
	ColumnMapping []ColumnMappingEntry `json:"columnMapping,omitempty"`
	// FileId: ID of the uploaded file
	FileId string `json:"fileId"`
	// InsertMode: Defines how to apply rows to a table.
	// upsert: insert or update rows from the upload to the table
	// append: insert rows from the upload to the table.
	InsertMode IngestInsertModeType `json:"insertMode"`
	// Options: FileOptions determine how the file will be opened and parsed
	Options FileOptions `json:"options"`
	// Schema: Schema of a file or of the corresponding managed table which will be created
	Schema ManagedTableSchema `json:"schema"`
	// TableId: Uuid or name of the table
	TableId string `json:"tableId"`
}

// GetColumnMapping returns the ColumnMapping property
func (m FileIngestRequest) GetColumnMapping() []ColumnMappingEntry {
	return m.ColumnMapping
}

// SetColumnMapping sets the ColumnMapping property
func (m FileIngestRequest) SetColumnMapping(val []ColumnMappingEntry) {
	m.ColumnMapping = val
}

// GetFileId returns the FileId property
func (m FileIngestRequest) GetFileId() string {
	return m.FileId
}

// SetFileId sets the FileId property
func (m FileIngestRequest) SetFileId(val string) {
	m.FileId = val
}

// GetInsertMode returns the InsertMode property
func (m FileIngestRequest) GetInsertMode() IngestInsertModeType {
	return m.InsertMode
}

// SetInsertMode sets the InsertMode property
func (m FileIngestRequest) SetInsertMode(val IngestInsertModeType) {
	m.InsertMode = val
}

// GetOptions returns the Options property
func (m FileIngestRequest) GetOptions() FileOptions {
	return m.Options
}

// SetOptions sets the Options property
func (m FileIngestRequest) SetOptions(val FileOptions) {
	m.Options = val
}

// GetSchema returns the Schema property
func (m FileIngestRequest) GetSchema() ManagedTableSchema {
	return m.Schema
}

// SetSchema sets the Schema property
func (m FileIngestRequest) SetSchema(val ManagedTableSchema) {
	m.Schema = val
}

// GetTableId returns the TableId property
func (m FileIngestRequest) GetTableId() string {
	return m.TableId
}

// SetTableId sets the TableId property
func (m FileIngestRequest) SetTableId(val string) {
	m.TableId = val
}
