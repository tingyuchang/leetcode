package Math

import (
	"math"
	"sort"
)

// O(nlgn)
func MinimumTime(time []int, totalTrips int) int64 {
	min, max := math.MaxInt, 0

	for _, v := range time {
		if v > max {
			max = v
		}

		if v < min {
			min = v
		}
	}

	max = totalTrips * max

	for max > min {
		mid := int(uint(min+max) >> 1)

		if check(time, totalTrips, mid) {
			max = mid
		} else {
			min = mid + 1
		}
	}

	return int64(min)

}

func check(time []int, totalTrips, estimated int) bool {
	trips := 0
	for _, v := range time {
		trips += estimated / v
	}

	if trips >= totalTrips {
		return true
	}

	return false
}

// O(n^2)
func MinimumTimeSlow(time []int, totalTrips int) int64 {
	timeDuration := 0
	sort.Ints(time)
	counter := make(map[int]int)

	for i, v := range time {
		counter[i] = v
	}

	for totalTrips > 0 {
		// find smallest number in counter
		min := math.MaxInt
		for i := 0; i < len(counter); i++ {
			if counter[i] < min {
				min = counter[i]
			}
		}

		// reset time in counter
		finished := 0
		for i := 0; i < len(counter); i++ {
			if counter[i] == min {
				counter[i] = time[i]
				finished++
			} else {
				counter[i] -= min
			}
		}

		totalTrips -= finished
		timeDuration += min
	}

	return int64(timeDuration)
}
