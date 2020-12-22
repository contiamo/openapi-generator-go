package generatortest

import (
	"testing"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestPerson(t *testing.T) {
	require.NoError(t, validation.Validate(&Person{
		Age:      18,
		Gender:   GenderDefault,
		Name:     "Foo",
		Hostname: "foobar.com",
		Email:    "foo@localhost.de",
		Uuid:     uuid.NewV4().String(),
	}))
}
