package templates

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTernary(t *testing.T) {
	require.Equal(t, "a", Ternary("a", "b", true))
	require.Equal(t, "b", Ternary("a", "b", false))

	require.Equal(t, 1, Ternary(1, 2, true))
	require.Equal(t, 2, Ternary(1, 2, false))

	require.Equal(t, 1.0, Ternary(1.0, 2.0, true))
	require.Equal(t, 2.0, Ternary(1.0, 2.0, false))

	require.Equal(t, true, Ternary(true, false, true))
	require.Equal(t, false, Ternary(true, false, false))

	require.Equal(t, "a", Ternary("a", 1, true))
	require.Equal(t, 1, Ternary("a", 1, false))
}
