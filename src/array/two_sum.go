package array

/*
Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.
You can return the answer in any order.
*/
func twoSum(nums []int, target int) []int {
	lookup := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		paired := target - nums[i]
		index, ok := lookup[paired]
		if ok {
			return []int{index, i}
		} else {
			lookup[nums[i]] = i
		}
	}
	return []int{}
}
