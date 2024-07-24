package array

/*
Given a 1-indexed array of integers numbers that is already sorted in non-decreasing order,
find two numbers such that they add up to a specific target number.

Let these two numbers be numbers[index1] and numbers[index2] where 1 <= index1 < index2 <= numbers.length.

Return the indices of the two numbers, index1 and index2, added by one as an integer array [index1, index2] of length 2.
Your solution must use only constant extra space.
*/

func twoSumSorted(numbers []int, target int) []int {
	a := 0
	b := len(numbers) - 1
	result := make([]int, 2)

	for {
		if a >= b {
			break
		}

		sum := numbers[a] + numbers[b]
		if sum > target {
			b--
		} else if sum < target {
			a++
		} else {
			result[0] = a + 1
			result[1] = b + 1
			break
		}
	}

	return result
}
