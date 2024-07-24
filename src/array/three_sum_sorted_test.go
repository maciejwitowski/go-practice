package array

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestThreeSumSorted(t *testing.T) {
	require.Equal(t, [][]int{{-1, -1, 2}, {-1, 0, 1}}, threeSum([]int{-1, 0, 1, 2, -1, -4}))
	require.Equal(t, [][]int{{-3, 1, 2}}, threeSum([]int{-3, 3, 4, -3, 1, 2}))
}
