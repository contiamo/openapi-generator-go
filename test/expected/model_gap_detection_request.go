// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// GapDetectionRequest is an object.
type GapDetectionRequest struct {
	// Column:
	Column string `json:"column"`
	// ProjectId:
	ProjectId string `json:"projectId"`
	// Sql:
	Sql string `json:"sql"`
}

// GetColumn returns the Column property
func (m GapDetectionRequest) GetColumn() string {
	return m.Column
}

// SetColumn sets the Column property
func (m GapDetectionRequest) SetColumn(val string) {
	m.Column = val
}

// GetProjectId returns the ProjectId property
func (m GapDetectionRequest) GetProjectId() string {
	return m.ProjectId
}

// SetProjectId sets the ProjectId property
func (m GapDetectionRequest) SetProjectId(val string) {
	m.ProjectId = val
}

// GetSql returns the Sql property
func (m GapDetectionRequest) GetSql() string {
	return m.Sql
}

// SetSql sets the Sql property
func (m GapDetectionRequest) SetSql(val string) {
	m.Sql = val
}
