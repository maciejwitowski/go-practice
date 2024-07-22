package array

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMinStack(t *testing.T) {
	require.Equal(t, []int{0, 1}, twoSum([]int{7, 2, 11, 15}, 9))
	require.Equal(t, []int{0, 1}, twoSum([]int{3, 3}, 6))
}
