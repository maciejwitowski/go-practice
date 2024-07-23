package array

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestLongestSequence(t *testing.T) {
	require.Equal(t, 3, longestSequence([]int{1, 5, 3, 2}))
	require.Equal(t, 9, longestSequence([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}))
}
