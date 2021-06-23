package generatortest

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBaz(t *testing.T) {
	data := []byte(`{"bar": 123}`)
	var m Baz
	require.NoError(t, json.Unmarshal(data, &m))
	foo, err := m.AsFoo()
	require.NoError(t, err)
	require.Equal(t, 123., foo.Bar)
}
