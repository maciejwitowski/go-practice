package array

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestBasic(t *testing.T) {
	require.Equal(t, []int{3, 2}, topFrequent([]int{1, 2, 2, 2, 3, 3, 3, 3}, 2))
}
