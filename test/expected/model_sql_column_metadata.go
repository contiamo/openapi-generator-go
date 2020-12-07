// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// SqlColumnMetadata is an object.
type SqlColumnMetadata struct {
	// DataType:
	DataType string `json:"dataType"`
	// Name: Column name
	Name string `json:"name"`
	// Nullable: Whether this column is nullable. May currently be a stub that always returns true.
	Nullable bool `json:"nullable"`
	// OriginalType: Information about the original type as provided by the upstream data source, if we have it.
	OriginalType string `json:"originalType,omitempty"`
	// Primitive:
	Primitive string `json:"primitive"`
	// Ref:
	Ref string `json:"ref"`
	// SqlType: The data type of the table column as it was discovered from the data source.
	SqlType TableColumnType `json:"sqlType"`
	// Type:
	Type string `json:"type"`
}

// GetDataType returns the DataType property
func (m SqlColumnMetadata) GetDataType() string {
	return m.DataType
}

// SetDataType sets the DataType property
func (m SqlColumnMetadata) SetDataType(val string) {
	m.DataType = val
}

// GetName returns the Name property
func (m SqlColumnMetadata) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m SqlColumnMetadata) SetName(val string) {
	m.Name = val
}

// GetNullable returns the Nullable property
func (m SqlColumnMetadata) GetNullable() bool {
	return m.Nullable
}

// SetNullable sets the Nullable property
func (m SqlColumnMetadata) SetNullable(val bool) {
	m.Nullable = val
}

// GetOriginalType returns the OriginalType property
func (m SqlColumnMetadata) GetOriginalType() string {
	return m.OriginalType
}

// SetOriginalType sets the OriginalType property
func (m SqlColumnMetadata) SetOriginalType(val string) {
	m.OriginalType = val
}

// GetPrimitive returns the Primitive property
func (m SqlColumnMetadata) GetPrimitive() string {
	return m.Primitive
}

// SetPrimitive sets the Primitive property
func (m SqlColumnMetadata) SetPrimitive(val string) {
	m.Primitive = val
}

// GetRef returns the Ref property
func (m SqlColumnMetadata) GetRef() string {
	return m.Ref
}

// SetRef sets the Ref property
func (m SqlColumnMetadata) SetRef(val string) {
	m.Ref = val
}

// GetSqlType returns the SqlType property
func (m SqlColumnMetadata) GetSqlType() TableColumnType {
	return m.SqlType
}

// SetSqlType sets the SqlType property
func (m SqlColumnMetadata) SetSqlType(val TableColumnType) {
	m.SqlType = val
}

// GetType returns the Type property
func (m SqlColumnMetadata) GetType() string {
	return m.Type
}

// SetType sets the Type property
func (m SqlColumnMetadata) SetType(val string) {
	m.Type = val
}
