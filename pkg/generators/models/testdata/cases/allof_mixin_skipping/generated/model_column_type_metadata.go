// This file is auto-generated, DO NOT EDIT.
//
// Source:
//
//	Title: Test
//	Version: 0.1.0
package generatortest

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// ColumnTypeMetadata is an object. Type metadata
type ColumnTypeMetadata struct {
	// Columns: List of columns if this type is structural
	Columns []ColumnMetadata `json:"columns,omitempty" mapstructure:"columns,omitempty"`
	// ItemType: Type metadata
	ItemType *ColumnTypeMetadata `json:"itemType,omitempty" mapstructure:"itemType,omitempty"`
	// Nullable: Column nullability
	Nullable Nullability `json:"nullable" mapstructure:"nullable"`
	// OriginalName: Original column type as given by data source
	OriginalName string `json:"originalName" mapstructure:"originalName"`
	// Type: Normalized column type. If type cannot be determined or is not compatible, then 'other'.
	Type ColumnType `json:"type" mapstructure:"type"`
}

// Validate implements basic validation for this model
func (m ColumnTypeMetadata) Validate() error {
	return validation.Errors{
		"columns": validation.Validate(
			m.Columns,
		),
		"itemType": validation.Validate(
			m.ItemType,
		),
		"nullable": validation.Validate(
			m.Nullable, validation.Required,
		),
		"type": validation.Validate(
			m.Type, validation.Required,
		),
	}.Filter()
}

// GetColumns returns the Columns property
func (m ColumnTypeMetadata) GetColumns() []ColumnMetadata {
	return m.Columns
}

// SetColumns sets the Columns property
func (m *ColumnTypeMetadata) SetColumns(val []ColumnMetadata) {
	m.Columns = val
}

// GetItemType returns the ItemType property
func (m ColumnTypeMetadata) GetItemType() *ColumnTypeMetadata {
	return m.ItemType
}

// SetItemType sets the ItemType property
func (m *ColumnTypeMetadata) SetItemType(val *ColumnTypeMetadata) {
	m.ItemType = val
}

// GetNullable returns the Nullable property
func (m ColumnTypeMetadata) GetNullable() Nullability {
	return m.Nullable
}

// SetNullable sets the Nullable property
func (m *ColumnTypeMetadata) SetNullable(val Nullability) {
	m.Nullable = val
}

// GetOriginalName returns the OriginalName property
func (m ColumnTypeMetadata) GetOriginalName() string {
	return m.OriginalName
}

// SetOriginalName sets the OriginalName property
func (m *ColumnTypeMetadata) SetOriginalName(val string) {
	m.OriginalName = val
}

// GetType returns the Type property
func (m ColumnTypeMetadata) GetType() ColumnType {
	return m.Type
}

// SetType sets the Type property
func (m *ColumnTypeMetadata) SetType(val ColumnType) {
	m.Type = val
}
