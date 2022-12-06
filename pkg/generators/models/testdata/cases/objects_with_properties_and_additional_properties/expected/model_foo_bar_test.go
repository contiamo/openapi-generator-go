package generatortest

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

type ChildStruct = struct {
	Test string `json:"test,omitempty"`
}

func TestFooBarJsonMarshalling(t *testing.T) {
	foobar := FooBar{
		FooBarProperties: FooBarProperties{
			Baz: "test",
		},
		AdditionalProperties: map[string][]ChildStruct{
			"arr1": []ChildStruct{
				ChildStruct{
					Test: "t1",
				},
				ChildStruct{
					Test: "t2",
				},
			},
			"arr2": []ChildStruct{
				ChildStruct{
					Test: "t3",
				},
				ChildStruct{
					Test: "t4",
				},
			},
		},
	}

	bytes, err := json.Marshal(foobar)
	require.NoError(t, err)

	var parsedFoobar FooBar
	err = json.Unmarshal(bytes, &parsedFoobar)
	require.NoError(t, err)

	require.Equal(t, foobar, parsedFoobar)
}
