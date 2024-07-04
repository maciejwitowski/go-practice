package src

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMergeIntervals(t *testing.T) {
	intervals := [][]int{
		{1, 3},
		{2, 6},
		{8, 10},
		{15, 18},
	}

	merged := Merge(intervals)
	expected := [][]int{
		{1, 6},
		{8, 10},
		{15, 18},
	}

	assert.Equal(t, expected, merged)
}
