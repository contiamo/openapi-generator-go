package templates

import "strings"

// TypeDisplayName returns the value with special characters replaced with
// there semantic meaning and then converted to pascal case. For example,
//     *string
// becomes
//     NullableString
func TypeDisplayName(value string) string {
	out := value
	if strings.HasPrefix(value, "*") {
		out = strings.Replace(out, "*", "nullable_", 1)
	}

	if strings.Contains(value, "interface{}") {
		out = strings.ReplaceAll(out, "interface{}", "interface")
	}

	if strings.Contains(value, "[]") {
		out = strings.ReplaceAll(out, "[]", "slice_")
	}

	return ToPascalCase(out)
}
