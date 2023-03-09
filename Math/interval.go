package Math

import "fmt"

func InsertInterval(intervals [][]int, newInterval []int) [][]int {
	start := newInterval[0]
	end := newInterval[1]

	insertIndex := -1
	endIndex := len(intervals)
	merge := false
	for i := 0; i < len(intervals); i++ {
		interval := intervals[i]

		if interval[0] <= start && interval[1] >= end {
			return intervals
		}
		if interval[0] > end {
			endIndex = i - 1
			if insertIndex == -1 {
				insertIndex = i - 1
			}
			break
		}

		if interval[1] < start {
			continue
		}

		if insertIndex == -1 {
			insertIndex = i
			merge = true
		}

	}

	max := func(a, b int) int {
		if a > b {
			return a
		}

		return b
	}

	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}
	fmt.Println(insertIndex, endIndex, merge)
	// merge or insert?
	if endIndex == len(intervals) {
		if merge {
			interval1 := intervals[insertIndex]
			interval2 := intervals[endIndex-1]
			start = min(start, interval1[0])
			end = max(end, interval2[1])
			intervals[insertIndex] = []int{start, end}
			return intervals[:insertIndex+1]
		}
		return append(intervals, newInterval)
	}

	if insertIndex == -1 {
		intervals = append([][]int{newInterval}, intervals...)
		return intervals
	}

	if merge {
		interval1 := intervals[insertIndex]
		interval2 := intervals[endIndex]

		start = min(start, interval1[0])
		end = max(end, interval2[1])
		intervals[insertIndex] = []int{start, end}
		for i := insertIndex + 1; i <= insertIndex+len(intervals)-1-endIndex; i++ {
			intervals[i] = intervals[endIndex-insertIndex+i]
		}

		return intervals[:insertIndex+len(intervals)-endIndex]
	} else {
		insertIndex++
		intervals = append(intervals, []int{})

		for i := len(intervals) - 1; i > insertIndex; i-- {
			intervals[i] = intervals[i-1]
		}
		intervals[insertIndex] = []int{start, end}
		return intervals
	}
}
