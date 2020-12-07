// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// EntityDetectionResult is an object.
type EntityDetectionResult struct {
	// Entity:
	Entity string `json:"entity"`
	// Examples:
	Examples []string `json:"examples,omitempty"`
	// Frequency:
	Frequency float32 `json:"frequency"`
	// Score:
	Score float32 `json:"score"`
}

// GetEntity returns the Entity property
func (m EntityDetectionResult) GetEntity() string {
	return m.Entity
}

// SetEntity sets the Entity property
func (m EntityDetectionResult) SetEntity(val string) {
	m.Entity = val
}

// GetExamples returns the Examples property
func (m EntityDetectionResult) GetExamples() []string {
	return m.Examples
}

// SetExamples sets the Examples property
func (m EntityDetectionResult) SetExamples(val []string) {
	m.Examples = val
}

// GetFrequency returns the Frequency property
func (m EntityDetectionResult) GetFrequency() float32 {
	return m.Frequency
}

// SetFrequency sets the Frequency property
func (m EntityDetectionResult) SetFrequency(val float32) {
	m.Frequency = val
}

// GetScore returns the Score property
func (m EntityDetectionResult) GetScore() float32 {
	return m.Score
}

// SetScore sets the Score property
func (m EntityDetectionResult) SetScore(val float32) {
	m.Score = val
}
