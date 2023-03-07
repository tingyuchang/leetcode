package greedy

import (
	"math"
	"sort"
)

func EraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	count := 0
	end := math.MinInt
	for _, interval := range intervals {
		if interval[0] >= end {
			end = interval[1]
		} else {
			count++
		}
	}

	return count
}
