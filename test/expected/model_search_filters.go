// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// SearchFilters is an object. Filters to be applied indexed by entity type
type SearchFilters struct {
	// Datasource:
	Datasource []Filter `json:"datasource,omitempty"`
	// Table:
	Table []Filter `json:"table,omitempty"`
	// Usecase:
	Usecase []Filter `json:"usecase,omitempty"`
}

// GetDatasource returns the Datasource property
func (m SearchFilters) GetDatasource() []Filter {
	return m.Datasource
}

// SetDatasource sets the Datasource property
func (m SearchFilters) SetDatasource(val []Filter) {
	m.Datasource = val
}

// GetTable returns the Table property
func (m SearchFilters) GetTable() []Filter {
	return m.Table
}

// SetTable sets the Table property
func (m SearchFilters) SetTable(val []Filter) {
	m.Table = val
}

// GetUsecase returns the Usecase property
func (m SearchFilters) GetUsecase() []Filter {
	return m.Usecase
}

// SetUsecase sets the Usecase property
func (m SearchFilters) SetUsecase(val []Filter) {
	m.Usecase = val
}
