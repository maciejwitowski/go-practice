package binarysearch

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	require.Equal(t, 3, binarySearch([]int{-1, 0, 2, 4, 6, 8}, 4))
	require.Equal(t, -1, binarySearch([]int{-1, 0, 2, 4, 6, 8}, 3))
	require.Equal(t, 4, binarySearch([]int{-1, 0, 3, 5, 9, 12}, 9))
}
