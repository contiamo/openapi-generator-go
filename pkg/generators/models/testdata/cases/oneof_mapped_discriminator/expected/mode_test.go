package generatortest

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSerialization(t *testing.T) {

	t.Run("generic error", func(t *testing.T) {

		rawMsg := `{"kind": "generic", "code": "err101", "message": "this is an error", "traceId": "ab1123"}`
		resp := Error{}

		err := json.Unmarshal([]byte(rawMsg), &resp)
		require.NoError(t, err)

		typed, err := resp.AsGenericError()
		require.NoError(t, err)

		require.Equal(t, "this is an error", typed.Message)
		require.Equal(t, "err101", typed.Code)
		require.Equal(t, "ab1123", typed.TraceId)

		out, err := json.Marshal(resp)
		require.NoError(t, err)
		require.JSONEq(t, rawMsg, string(out))
	})

	t.Run("external error", func(t *testing.T) {

		rawMsg := `{"kind": "external", "service": "google", "traceId": "ab1123"}`
		resp := Error{}

		err := json.Unmarshal([]byte(rawMsg), &resp)
		require.NoError(t, err)

		typed, err := resp.AsExternalError()
		require.NoError(t, err)

		require.Equal(t, "google", typed.Service)
		require.Equal(t, "ab1123", typed.TraceId)

		out, err := json.Marshal(resp)
		require.NoError(t, err)
		require.JSONEq(t, rawMsg, string(out))
	})

	t.Run("external error", func(t *testing.T) {

		rawMsg := `{"kind": "unknown", "service": "google", "traceId": "ab1123"}`
		resp := Error{}

		err := json.Unmarshal([]byte(rawMsg), &resp)
		require.EqualError(t, err, "kind: invalid value.")
	})

}

func TestValidation(t *testing.T) {
	rawMsg := `{"kind": "external", "service": "", "traceId": ""}`
	resp := Error{}

	err := json.Unmarshal([]byte(rawMsg), &resp)
	require.NoError(t, err)

	err = resp.Validate()
	require.EqualError(t, err, "traceId: cannot be blank.")
}
