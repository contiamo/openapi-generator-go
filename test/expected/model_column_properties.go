// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// ColumnProperties is an object.
type ColumnProperties struct {
	// Entities: A list of detected entities in the column
	Entities []string `json:"entities"`
	// Nullable:
	Nullable bool `json:"nullable"`
	// OriginalType:
	OriginalType string `json:"originalType"`
	// Type: The data type of the table column as it was discovered from the data source.
	Type TableColumnType `json:"type"`
}

// GetEntities returns the Entities property
func (m ColumnProperties) GetEntities() []string {
	return m.Entities
}

// SetEntities sets the Entities property
func (m ColumnProperties) SetEntities(val []string) {
	m.Entities = val
}

// GetNullable returns the Nullable property
func (m ColumnProperties) GetNullable() bool {
	return m.Nullable
}

// SetNullable sets the Nullable property
func (m ColumnProperties) SetNullable(val bool) {
	m.Nullable = val
}

// GetOriginalType returns the OriginalType property
func (m ColumnProperties) GetOriginalType() string {
	return m.OriginalType
}

// SetOriginalType sets the OriginalType property
func (m ColumnProperties) SetOriginalType(val string) {
	m.OriginalType = val
}

// GetType returns the Type property
func (m ColumnProperties) GetType() TableColumnType {
	return m.Type
}

// SetType sets the Type property
func (m ColumnProperties) SetType(val TableColumnType) {
	m.Type = val
}
