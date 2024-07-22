package stack

import (
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestBasic(t *testing.T) {
	cases := map[string]int{
		"2 1 + 3 *":                     9,
		"4 13 5 / +":                    6,
		"10 6 9 3 + -11 * / * 17 + 5 +": 22,
	}

	for s, result := range cases {
		tokens := strings.Split(s, " ")
		require.Equal(t, result, evalRPN(tokens))
	}
}
