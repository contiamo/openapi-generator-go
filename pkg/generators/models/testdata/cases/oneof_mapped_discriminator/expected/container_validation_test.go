package generatortest

import (
	"encoding/json"
	"testing"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/stretchr/testify/require"
)

func TestContainerValidation(t *testing.T) {
	t.Run("validation succeeds for missing property", func(t *testing.T) {
		rawBytes := `{}`
		var container Container

		err := json.Unmarshal([]byte(rawBytes), &container)
		require.NoError(t, err)

		err = container.Validate()
		require.NoError(t, err)
	})

	t.Run("validation succeeds for valid property", func(t *testing.T) {
		rawBytes := `{"error":{"kind": "external", "service": "", "traceId": "12345"}}`
		var container Container

		err := json.Unmarshal([]byte(rawBytes), &container)
		require.NoError(t, err)

		err = container.Validate()
		require.NoError(t, err)
	})

	t.Run("validation fails for invalid property", func(t *testing.T) {
		rawBytes := `{"error":{"kind": "external", "service": "", "traceId": ""}}`
		var container Container

		err := json.Unmarshal([]byte(rawBytes), &container)
		require.NoError(t, err)

		err = container.Validate()
		require.EqualError(t, err, "error: (traceId: cannot be blank.).")
	})
}

func TestNonEmptyContainerValidation(t *testing.T) {
	t.Run("validation fails for empty payload", func(t *testing.T) {
		rawBytes := `{}`
		var container NonEmptyContainer

		err := json.Unmarshal([]byte(rawBytes), &container)
		require.NoError(t, err)

		err = container.Validate()
		require.Equal(t, err, validation.Errors{
			"error": EmptyErrorError,
		})
	})

	t.Run("validation succeeds for valid property", func(t *testing.T) {
		rawBytes := `{"error":{"kind": "external", "service": "", "traceId": "12345"}}`
		var container NonEmptyContainer

		err := json.Unmarshal([]byte(rawBytes), &container)
		require.NoError(t, err)

		err = container.Validate()
		require.NoError(t, err)
	})

	t.Run("validation succeeds for valid property", func(t *testing.T) {
		rawBytes := `{"error":{"kind": "external", "service": "", "traceId": ""}}`
		var container NonEmptyContainer

		err := json.Unmarshal([]byte(rawBytes), &container)
		require.NoError(t, err)

		err = container.Validate()
		require.EqualError(t, err, "error: (traceId: cannot be blank.).")
	})
}
