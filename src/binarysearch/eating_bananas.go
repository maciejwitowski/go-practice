package binarysearch

/*
You are given an integer array piles where piles[i] is the number of bananas in the ith pile. You are also given an integer h, which represents the number of hours you have to eat all the bananas.

You may decide your bananas-per-hour eating rate of k. Each hour, you may choose a pile of bananas and eats k bananas from that pile. If the pile has less than k bananas, you may finish eating the pile but you can not eat from another pile in the same hour.
Return the minimum integer k such that you can eat all the bananas within h hours.
*/

func eatingBananas(piles []int, h int) int {
	var highest int
	for _, p := range piles {
		if p > highest {
			highest = p
		}
	}

	i := 1
	j := highest
	var minM int

	for i <= j {
		m := i + (j-i)/2
		ateOnTime := managedToAtePiles(piles, m, h)
		if ateOnTime {
			minM = m
			j = m - 1
		} else {
			i = m + 1
		}
	}
	return minM
}

// true if managed to eat all piles within h having pace m, false otherwise
func managedToAtePiles(piles []int, m int, h int) bool {
	for _, pileSize := range piles {
		timeToEatPile := pileSize / m
		if pileSize%m != 0 {
			timeToEatPile += 1
		}

		h -= timeToEatPile
		if h < 0 {
			return false
		}
	}
	return true
}
