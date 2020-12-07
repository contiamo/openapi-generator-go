// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// FileDiscoveryRequest is an object.
type FileDiscoveryRequest struct {
	// Options: FileOptions determine how the file will be opened and parsed
	Options FileOptions `json:"options"`
	// TableId: Optional ID of the target table
	TableId string `json:"tableId,omitempty"`
}

// GetOptions returns the Options property
func (m FileDiscoveryRequest) GetOptions() FileOptions {
	return m.Options
}

// SetOptions sets the Options property
func (m FileDiscoveryRequest) SetOptions(val FileOptions) {
	m.Options = val
}

// GetTableId returns the TableId property
func (m FileDiscoveryRequest) GetTableId() string {
	return m.TableId
}

// SetTableId sets the TableId property
func (m FileDiscoveryRequest) SetTableId(val string) {
	m.TableId = val
}
