// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// MaterializationSpec is an object.
type MaterializationSpec struct {
	// Cron: The cron schedule for updating the materialization (https://en.wikipedia.org/wiki/Cron#Overview).
	// An empty string indicates a one-time-job.
	Cron string `json:"cron,omitempty"`
	// Enabled:
	Enabled bool `json:"enabled"`
	// ExpiresUnit: The time unit associated with expiresIn; can only be null if expiresIn is 0.
	ExpiresUnit *MaterializationTimeUnits `json:"expiresUnit,omitempty"`
	// ExpiresValue: Indicates how long before the materialization expires without being refreshed.
	// A value of zero indicates the result does not expire.
	ExpiresValue int32 `json:"expiresValue,omitempty"`
}

// GetCron returns the Cron property
func (m MaterializationSpec) GetCron() string {
	return m.Cron
}

// SetCron sets the Cron property
func (m MaterializationSpec) SetCron(val string) {
	m.Cron = val
}

// GetEnabled returns the Enabled property
func (m MaterializationSpec) GetEnabled() bool {
	return m.Enabled
}

// SetEnabled sets the Enabled property
func (m MaterializationSpec) SetEnabled(val bool) {
	m.Enabled = val
}

// GetExpiresUnit returns the ExpiresUnit property
func (m MaterializationSpec) GetExpiresUnit() *MaterializationTimeUnits {
	return m.ExpiresUnit
}

// SetExpiresUnit sets the ExpiresUnit property
func (m MaterializationSpec) SetExpiresUnit(val *MaterializationTimeUnits) {
	m.ExpiresUnit = val
}

// GetExpiresValue returns the ExpiresValue property
func (m MaterializationSpec) GetExpiresValue() int32 {
	return m.ExpiresValue
}

// SetExpiresValue sets the ExpiresValue property
func (m MaterializationSpec) SetExpiresValue(val int32) {
	m.ExpiresValue = val
}
