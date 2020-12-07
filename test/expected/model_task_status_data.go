// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// TaskStatusData is an object.
type TaskStatusData struct {
	// DetectedRowCount: The number of rows that were detected in the import
	DetectedRowCount int32 `json:"detectedRowCount"`
	// ErrorRowCount: The number of rows that errored during import
	ErrorRowCount int32 `json:"errorRowCount"`
	// FileSize:
	FileSize string `json:"fileSize"`
	// ImportedRowCount: The number of rows successfully imported
	ImportedRowCount int32 `json:"importedRowCount"`
}

// GetDetectedRowCount returns the DetectedRowCount property
func (m TaskStatusData) GetDetectedRowCount() int32 {
	return m.DetectedRowCount
}

// SetDetectedRowCount sets the DetectedRowCount property
func (m TaskStatusData) SetDetectedRowCount(val int32) {
	m.DetectedRowCount = val
}

// GetErrorRowCount returns the ErrorRowCount property
func (m TaskStatusData) GetErrorRowCount() int32 {
	return m.ErrorRowCount
}

// SetErrorRowCount sets the ErrorRowCount property
func (m TaskStatusData) SetErrorRowCount(val int32) {
	m.ErrorRowCount = val
}

// GetFileSize returns the FileSize property
func (m TaskStatusData) GetFileSize() string {
	return m.FileSize
}

// SetFileSize sets the FileSize property
func (m TaskStatusData) SetFileSize(val string) {
	m.FileSize = val
}

// GetImportedRowCount returns the ImportedRowCount property
func (m TaskStatusData) GetImportedRowCount() int32 {
	return m.ImportedRowCount
}

// SetImportedRowCount sets the ImportedRowCount property
func (m TaskStatusData) SetImportedRowCount(val int32) {
	m.ImportedRowCount = val
}
