package binarysearch

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMinRotatedSortedArray(t *testing.T) {
	require.Equal(t, 1, minRotatedSortedArray([]int{3, 4, 5, 6, 1, 2}))
	require.Equal(t, 0, minRotatedSortedArray([]int{4, 5, 6, 7, 0, 1, 2}))
	require.Equal(t, 1, minRotatedSortedArray([]int{1}))
}
