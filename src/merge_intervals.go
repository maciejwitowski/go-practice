package src

import (
	"sort"
)

// Merge
// Given an array of intervals where intervals[i] = [starti, endi], merge all overlapping intervals, and return an array of the non-overlapping intervals that cover all the intervals in the input.
func Merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return make([][]int, 0)
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	result := [][]int{intervals[0]}

	for i := 1; i < len(intervals); i++ {
		last := result[len(result)-1]
		current := intervals[i]

		if current[0] <= last[1] {
			if current[1] > last[1] {
				last[1] = current[1]
			}
		} else {
			result = append(result, current)
		}
	}
	return result
}
