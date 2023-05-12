package Math

func InsertInterval(intervals [][]int, newInterval []int) [][]int { // insert interval
	intervals = insertInterval(intervals, newInterval)
	ans := make([][]int, 0)
	for i := 0; i < len(intervals); i++ {
		currentInterval := intervals[i]
		for i < len(intervals) && hasOverLap(currentInterval, intervals[i]) {
			currentInterval = mergeIntervals(currentInterval, intervals[i])
			i++
		}

		i--
		ans = append(ans, currentInterval)
	}

	return ans

}

func insertInterval(intervals [][]int, newInterval []int) [][]int {
	l := 0
	r := len(intervals) - 1
	for r >= l {

		mid := int(uint(l+r) >> 1)

		if intervals[mid][0] < newInterval[0] {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	if l != len(intervals) {
		intervals = append(intervals[:l+1], intervals[l:]...)
		intervals[l] = newInterval
	} else {
		intervals = append(intervals, newInterval)
	}

	return intervals
}

func mergeIntervals(a, b []int) []int {
	return []int{min(a[0], b[0]), max(a[1], b[1])}
}

func hasOverLap(a, b []int) bool {
	return (min(a[1], b[1]) - max(a[0], b[0])) >= 0
}
