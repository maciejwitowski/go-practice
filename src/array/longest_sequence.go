package array

/*
Given an unsorted array of integers nums, return the length of the longest consecutive elements sequence.
*/

func longestSequence(nums []int) int {
	set := map[int]bool{}

	for _, n := range nums {
		set[n] = true
	}

	maxLength := 0
	for _, n := range nums {
		if set[n+1] {
			continue
		}

		// Sequence start. Go right to find length
		length := 0
		for curr := n; set[curr]; curr-- {
			length++
		}
		if length > maxLength {
			maxLength = length
		}
	}
	return maxLength
}
