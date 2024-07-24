package array

/*
You are given an integer array height of length n. There are n vertical lines drawn such that the two endpoints of the ith line are (i, 0) and (i, height[i]).

Find two lines that together with the x-axis form a container, such that the container contains the most water.
Return the maximum amount of water a container can store.
*/

func maxArea(height []int) int {
	a := 0
	b := len(height) - 1

	maxArea := 0
	for a < b {
		y := min(height[a], height[b])
		x := b - a
		area := y * x
		maxArea = max(maxArea, area)

		if height[a] <= height[b] {
			a++
		} else {
			b--
		}
	}
	return maxArea
}
