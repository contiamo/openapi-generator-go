package templates

import (
	"fmt"
)

// MapFunc returns a map from the given key-value pairs.
func MapFunc(pairs ...any) (map[string]any, error) {
	if len(pairs)%2 != 0 {
		return nil, fmt.Errorf("expecting even number of arguments, but got: %d", len(pairs))
	}

	m := make(map[string]any, len(pairs)/2)
	for i := 0; i < len(pairs); i += 2 {
		key, ok := pairs[i].(string)
		if !ok {
			return nil, fmt.Errorf("expecting type string as map key, got: %#v as map key", pairs[i])
		}
		m[key] = pairs[i+1]
	}
	return m, nil
}
