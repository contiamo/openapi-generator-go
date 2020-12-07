// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// GlossaryItemStatistics is an object. Various statistics about the glossary item
type GlossaryItemStatistics struct {
	// LinkCount: How many resources are linked to this glossary item
	LinkCount int32 `json:"linkCount"`
}

// GetLinkCount returns the LinkCount property
func (m GlossaryItemStatistics) GetLinkCount() int32 {
	return m.LinkCount
}

// SetLinkCount sets the LinkCount property
func (m GlossaryItemStatistics) SetLinkCount(val int32) {
	m.LinkCount = val
}
