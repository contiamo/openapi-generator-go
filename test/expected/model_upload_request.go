// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// UploadRequest is an object. File uploaded by the client
type UploadRequest struct {
	// File:
	File string `json:"file"`
	// TableId:
	TableId string `json:"tableId,omitempty"`
}

// GetFile returns the File property
func (m UploadRequest) GetFile() string {
	return m.File
}

// SetFile sets the File property
func (m UploadRequest) SetFile(val string) {
	m.File = val
}

// GetTableId returns the TableId property
func (m UploadRequest) GetTableId() string {
	return m.TableId
}

// SetTableId sets the TableId property
func (m UploadRequest) SetTableId(val string) {
	m.TableId = val
}
