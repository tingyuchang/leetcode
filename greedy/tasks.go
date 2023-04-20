package greedy

func LeastInterval(tasks []byte, n int) int {
	return leastInterval(tasks, n)
}

func leastInterval(tasks []byte, n int) int {
	cache := make([]int, 26)
	maxVal := 0
	maxCount := 0

	for _, v := range tasks {
		cache[v-'A']++

		if maxVal == cache[v-'A'] {
			maxCount++
		} else if cache[v-'A'] > maxVal {
			maxVal = cache[v-'A']
			maxCount = 1
		}
	}

	max := func(a, b int) int {
		if a > b {
			return a
		}

		return b
	}

	intervals := maxVal - 1
	intervalLength := n - (maxCount - 1)
	slots := intervalLength * intervals
	availableTasks := len(tasks) - maxCount*maxVal
	idles := max(0, slots-availableTasks)

	return len(tasks) + idles

}
