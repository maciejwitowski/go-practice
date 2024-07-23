package array

/*
Given an integer array nums and an integer k, return the k most frequent elements within the array.

The test cases are generated such that the answer is always unique.
You may return the output in any order.
*/

func topFrequent(nums []int, k int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	frequencies := make(map[int]int)
	indexedFrequencies := make([][]int, len(nums)+1)

	for i := 0; i < len(nums); i++ {
		frequencies[nums[i]] = frequencies[nums[i]] + 1
	}

	for num, freq := range frequencies {
		indexedFrequencies[freq] = append(indexedFrequencies[freq], num)
	}

	var result []int
	for i := len(indexedFrequencies) - 1; i > 0; i-- {
		for _, num := range indexedFrequencies[i] {
			result = append(result, num)
			if len(result) == k {
				return result
			}
		}
	}
	return result
}
