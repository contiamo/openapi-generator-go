// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

import (
	"time"
)

// ColumnResponse is an object.
type ColumnResponse struct {
	// CreatedAt:
	CreatedAt time.Time `json:"createdAt"`
	// Description:
	Description string `json:"description"`
	// Documentation: Long form markdown field that can be used to document the resource
	Documentation string `json:"documentation"`
	// Entities: A list of detected entities in the column
	Entities []string `json:"entities"`
	// Id:
	Id string `json:"id"`
	// IsOrphaned: Indicates that this column couldn't be synced from an external system
	IsOrphaned bool `json:"isOrphaned"`
	// Name:
	Name string `json:"name"`
	// Nullable:
	Nullable bool `json:"nullable"`
	// OriginalType: The column type as named by the data source
	OriginalType string `json:"originalType"`
	// OrphanedAt:
	OrphanedAt *time.Time `json:"orphanedAt"`
	// Properties:
	Properties map[string]interface{} `json:"properties"`
	// TableId:
	TableId string `json:"tableId"`
	// Type: The data type of the table column as it was discovered from the data source.
	Type TableColumnType `json:"type"`
	// UpdatedAt:
	UpdatedAt time.Time `json:"updatedAt"`
}

// GetCreatedAt returns the CreatedAt property
func (m ColumnResponse) GetCreatedAt() time.Time {
	return m.CreatedAt
}

// SetCreatedAt sets the CreatedAt property
func (m ColumnResponse) SetCreatedAt(val time.Time) {
	m.CreatedAt = val
}

// GetDescription returns the Description property
func (m ColumnResponse) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m ColumnResponse) SetDescription(val string) {
	m.Description = val
}

// GetDocumentation returns the Documentation property
func (m ColumnResponse) GetDocumentation() string {
	return m.Documentation
}

// SetDocumentation sets the Documentation property
func (m ColumnResponse) SetDocumentation(val string) {
	m.Documentation = val
}

// GetEntities returns the Entities property
func (m ColumnResponse) GetEntities() []string {
	return m.Entities
}

// SetEntities sets the Entities property
func (m ColumnResponse) SetEntities(val []string) {
	m.Entities = val
}

// GetId returns the Id property
func (m ColumnResponse) GetId() string {
	return m.Id
}

// SetId sets the Id property
func (m ColumnResponse) SetId(val string) {
	m.Id = val
}

// GetIsOrphaned returns the IsOrphaned property
func (m ColumnResponse) GetIsOrphaned() bool {
	return m.IsOrphaned
}

// SetIsOrphaned sets the IsOrphaned property
func (m ColumnResponse) SetIsOrphaned(val bool) {
	m.IsOrphaned = val
}

// GetName returns the Name property
func (m ColumnResponse) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m ColumnResponse) SetName(val string) {
	m.Name = val
}

// GetNullable returns the Nullable property
func (m ColumnResponse) GetNullable() bool {
	return m.Nullable
}

// SetNullable sets the Nullable property
func (m ColumnResponse) SetNullable(val bool) {
	m.Nullable = val
}

// GetOriginalType returns the OriginalType property
func (m ColumnResponse) GetOriginalType() string {
	return m.OriginalType
}

// SetOriginalType sets the OriginalType property
func (m ColumnResponse) SetOriginalType(val string) {
	m.OriginalType = val
}

// GetOrphanedAt returns the OrphanedAt property
func (m ColumnResponse) GetOrphanedAt() *time.Time {
	return m.OrphanedAt
}

// SetOrphanedAt sets the OrphanedAt property
func (m ColumnResponse) SetOrphanedAt(val *time.Time) {
	m.OrphanedAt = val
}

// GetProperties returns the Properties property
func (m ColumnResponse) GetProperties() map[string]interface{} {
	return m.Properties
}

// SetProperties sets the Properties property
func (m ColumnResponse) SetProperties(val map[string]interface{}) {
	m.Properties = val
}

// GetTableId returns the TableId property
func (m ColumnResponse) GetTableId() string {
	return m.TableId
}

// SetTableId sets the TableId property
func (m ColumnResponse) SetTableId(val string) {
	m.TableId = val
}

// GetType returns the Type property
func (m ColumnResponse) GetType() TableColumnType {
	return m.Type
}

// SetType sets the Type property
func (m ColumnResponse) SetType(val TableColumnType) {
	m.Type = val
}

// GetUpdatedAt returns the UpdatedAt property
func (m ColumnResponse) GetUpdatedAt() time.Time {
	return m.UpdatedAt
}

// SetUpdatedAt sets the UpdatedAt property
func (m ColumnResponse) SetUpdatedAt(val time.Time) {
	m.UpdatedAt = val
}
