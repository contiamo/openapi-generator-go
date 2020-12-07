// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// ConnectionTestWithStatsResult is an object.
type ConnectionTestWithStatsResult struct {
	// Message:
	Message string `json:"message,omitempty"`
	// Stats:
	Stats ConnectionStats `json:"stats,omitempty"`
	// Valid:
	Valid bool `json:"valid"`
}

// GetMessage returns the Message property
func (m ConnectionTestWithStatsResult) GetMessage() string {
	return m.Message
}

// SetMessage sets the Message property
func (m ConnectionTestWithStatsResult) SetMessage(val string) {
	m.Message = val
}

// GetStats returns the Stats property
func (m ConnectionTestWithStatsResult) GetStats() ConnectionStats {
	return m.Stats
}

// SetStats sets the Stats property
func (m ConnectionTestWithStatsResult) SetStats(val ConnectionStats) {
	m.Stats = val
}

// GetValid returns the Valid property
func (m ConnectionTestWithStatsResult) GetValid() bool {
	return m.Valid
}

// SetValid sets the Valid property
func (m ConnectionTestWithStatsResult) SetValid(val bool) {
	m.Valid = val
}
