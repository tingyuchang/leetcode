package BinarySearch

import "math"

func minEatingSpeed(piles []int, h int) int {
	l, r := 1, 0

	for _, p := range piles {
		r += p
	}

	for r > l {
		mid := (l + r) / 2

		estimate := 0

		for _, pile := range piles {
			estimate += int(math.Ceil(float64(pile) / float64(mid)))
		}

		if estimate <= h {
			r = mid
		} else {
			l = mid + 1
		}

	}

	return l
}

/*
using Binary search to find best answer

[maxWight -> totalWeight] must exist 1 capacity
*/
func shipWithinDays(weights []int, days int) int {

	l, r := 0, 0

	for _, w := range weights {
		r += w
		if w > l {
			l = w
		}
	}

	for r > l {
		mid := (l + r) / 2

		sum := 0
		d := 1
		for _, w := range weights {
			sum += w
			if sum > mid {
				d++
				sum = w
			}
		}

		if d <= days {
			r = mid
		} else {
			l = mid + 1
		}
	}

	return l
}
