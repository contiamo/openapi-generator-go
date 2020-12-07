// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// TableProfileResponse is an object.
type TableProfileResponse struct {
	// Profile:
	Profile *ProfileResponse `json:"profile,omitempty"`
	// Progress:
	Progress ProfilingProgress `json:"progress"`
}

// GetProfile returns the Profile property
func (m TableProfileResponse) GetProfile() *ProfileResponse {
	return m.Profile
}

// SetProfile sets the Profile property
func (m TableProfileResponse) SetProfile(val *ProfileResponse) {
	m.Profile = val
}

// GetProgress returns the Progress property
func (m TableProfileResponse) GetProgress() ProfilingProgress {
	return m.Progress
}

// SetProgress sets the Progress property
func (m TableProfileResponse) SetProgress(val ProfilingProgress) {
	m.Progress = val
}
