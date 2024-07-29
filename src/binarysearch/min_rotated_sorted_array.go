package binarysearch

/*
You are given an array of length n which was originally sorted in ascending order. It has now been rotated between 1 and n times. For example, the array nums = [1,2,3,4,5,6] might become:

[3,4,5,6,1,2] if it was rotated 4 times.
[1,2,3,4,5,6] if it was rotated 6 times.
Notice that rotating the array 4 times moves the last four elements of the array to the beginning. Rotating the array 6 times produces the original array.

Assuming all elements in the rotated sorted array nums are unique, return the minimum element of this array.
*/
func minRotatedSortedArray(nums []int) int {
	l := 0
	r := len(nums) - 1
	minimum := nums[0]

	for l <= r {
		m := l + (r-l)/2
		if nums[m] < minimum {
			minimum = nums[m]
		}

		if nums[m] >= nums[r] {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return minimum
}
