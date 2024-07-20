package stack

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSimplifyPath(t *testing.T) {
	cases := map[string]string{
		"/home/":                           "/home",
		"/home//foo":                       "/home/foo",
		"/home/user/Documents/../Pictures": "/home/user/Pictures",
		"/../":                             "/",
		"/.../a/../b/c/../d/./":            "/.../b/d",
	}

	for input, expected := range cases {
		require.Equal(t, expected, simplifyPath(input))
	}
}
