// Code generated by openapi-generator-go DO NOT EDIT.
//
// Source:
//
//	Title: Test
//	Version: 0.1.0
package generatortest

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// ArtistLeftFoot is an enum.
type ArtistLeftFoot string

// DefaultArtistLeftFoot is the default value for ArtistLeftFoot
var DefaultArtistLeftFoot = ArtistLeftFoot("blue")

// Validate implements basic validation for this model
func (m ArtistLeftFoot) Validate() error {
	return InKnownArtistLeftFoot.Validate(m)
}

var (
	ArtistLeftFootBlue  ArtistLeftFoot = "blue"
	ArtistLeftFootGreen ArtistLeftFoot = "green"
	ArtistLeftFootRed   ArtistLeftFoot = "red"

	// KnownArtistLeftFoot is the list of valid ArtistLeftFoot
	KnownArtistLeftFoot = []ArtistLeftFoot{
		ArtistLeftFootBlue,
		ArtistLeftFootGreen,
		ArtistLeftFootRed,
	}

	// KnownArtistLeftFootString is the list of valid ArtistLeftFoot as string
	KnownArtistLeftFootString = []string{
		string(ArtistLeftFootBlue),
		string(ArtistLeftFootGreen),
		string(ArtistLeftFootRed),
	}

	// InKnownArtistLeftFoot is an ozzo-validator for ArtistLeftFoot
	InKnownArtistLeftFoot = validation.In(
		ArtistLeftFootBlue,
		ArtistLeftFootGreen,
		ArtistLeftFootRed,
	)
)