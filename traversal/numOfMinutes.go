package traversal

func NumOfMinutes(n int, headID int, manager []int, informTime []int) int {
	return numOfMinutes(n, headID, manager, informTime)
}

func numOfMinutes(n int, headID int, manager []int, informTime []int) int {
	cache := make(map[int][]int, n)

	for i, v := range manager {
		cache[v] = append(cache[v], i)
	}

	queue := []int{headID}

	for len(queue) != 0 {
		m := queue[0]
		queue = queue[1:]

		queue = append(queue, cache[m]...)
		for _, v := range cache[m] {
			informTime[v] = informTime[v] + informTime[m]
		}

	}

	res := 0
	for i := 0; i < len(informTime); i++ {
		if informTime[i] > res {
			res = informTime[i]
		}
	}

	return res
}
