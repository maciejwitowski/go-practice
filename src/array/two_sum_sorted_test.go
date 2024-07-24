package array

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestTwoSumSorted(t *testing.T) {
	require.Equal(t, []int{1, 2}, twoSumSorted([]int{2, 7, 11, 15}, 9))
	require.Equal(t, []int{1, 3}, twoSumSorted([]int{2, 3, 4}, 6))
	require.Equal(t, []int{1, 2}, twoSumSorted([]int{-1, 0}, -1))
}
