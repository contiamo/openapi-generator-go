package generatortest

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestContainerValidation(t *testing.T) {
	rawBytes := `{}`
	var container Container

	err := json.Unmarshal([]byte(rawBytes), &container)
	require.NoError(t, err)

	err = container.Validate()
	require.NoError(t, err)
}
