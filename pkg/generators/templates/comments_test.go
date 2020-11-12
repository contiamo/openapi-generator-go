package templates

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCommentBlock(t *testing.T) {
	require.Equal(t, "// some\n// multiline\n// comment", CommentBlock("some\nmultiline\ncomment\n"))
	require.Equal(t, "// single line comment", CommentBlock("single line comment"))
}
