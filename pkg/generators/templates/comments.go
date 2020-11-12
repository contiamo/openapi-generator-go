package templates

import "strings"

// CommentBlock outputs the string as a comment block, ie prefixing each line with //
func CommentBlock(str string) string {
	return "// " + strings.Replace(strings.TrimSpace(str), "\n", "\n// ", -1)
}
