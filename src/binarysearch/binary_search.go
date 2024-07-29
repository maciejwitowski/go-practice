package binarysearch

func binarySearch(nums []int, target int) int {
	l := 0
	r := len(nums) - 1

	for l <= r {
		m := l + (r-l)/2

		if nums[m] > target {
			r = m - 1
		} else if nums[m] < target {
			l = m + 1
		} else {
			return m
		}
	}
	return -1
}