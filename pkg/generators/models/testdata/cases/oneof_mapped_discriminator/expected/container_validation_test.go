package generatortest

import (
	"encoding/json"
	"testing"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/stretchr/testify/require"
)

func TestContainerValidation(t *testing.T) {
	rawBytes := `{}`
	var container Container

	err := json.Unmarshal([]byte(rawBytes), &container)
	require.NoError(t, err)

	err = container.Validate()
	require.Equal(t, err, validation.Errors{
		"error": EmptyErrorError,
	})
}
