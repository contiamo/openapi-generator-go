package templates

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTypeDisplayName(t *testing.T) {
	require.Equal(t, "Interface", TypeDisplayName("interface{}"))
	require.Equal(t, "NullableString", TypeDisplayName("*string"))
	require.Equal(t, "StringSlice", TypeDisplayName("[]string"))
	require.Equal(t, "InterfaceSlice", TypeDisplayName("[]interface{}"))
	require.Equal(t, "InterfaceSliceSlice", TypeDisplayName("[][]interface{}"))
	require.Equal(t, "NullableInterfaceSliceSlice", TypeDisplayName("*[][]interface{}"))
	require.Equal(t, "StructNamespaceName", TypeDisplayName("struct { Namespace string\nName string }"))
}
