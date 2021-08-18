package templates

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTypeDisplayName(t *testing.T) {
	require.Equal(t, "NullableString", TypeDisplayName("*string"))
	require.Equal(t, "SliceString", TypeDisplayName("[]string"))
	require.Equal(t, "SliceInterface", TypeDisplayName("[]interface{}"))
	require.Equal(t, "SliceSliceInterface", TypeDisplayName("[][]interface{}"))
	require.Equal(t, "NullableSliceSliceInterface", TypeDisplayName("*[][]interface{}"))
}
