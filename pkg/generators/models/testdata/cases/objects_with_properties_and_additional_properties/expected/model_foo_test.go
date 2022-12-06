package generatortest

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFooJsonMarshalling(t *testing.T) {
	foo := Foo{
		FooProperties: FooProperties{
			Bar: 10,
		},
		AdditionalProperties: map[string]interface{}{
			"scalar": float64(200),
			"obj": map[string]interface{}{
				"nested": float64(100),
			},
			"arr": []interface{}{"v1", "v2"},
		},
	}

	bytes, err := json.Marshal(foo)
	require.NoError(t, err)

	var parsedFoo Foo
	err = json.Unmarshal(bytes, &parsedFoo)
	require.NoError(t, err)

	require.Equal(t, foo, parsedFoo)
}
