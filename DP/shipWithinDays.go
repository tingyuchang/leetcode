package DP

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
