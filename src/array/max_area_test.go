package array

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMaxArea(t *testing.T) {
	require.Equal(t, 49, maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}
