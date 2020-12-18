package generatortest

import (
	"testing"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestPerson(t *testing.T) {
	require.NoError(t, validation.Validate(&Person{
		Age:      120,
		Gender:   GenderDefault,
		Name:     "Foobar",
		Hostname: "localhost",
		Email:    "foo@localhost",
		Uuid:     uuid.NewV4().String(),
	}))
}
