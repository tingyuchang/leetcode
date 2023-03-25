package traversal

func CountPairs(n int, edges [][]int) int64 {
	res := 0
	cache := make(map[int][]int)

	for _, edge := range edges {
		cache[edge[0]] = append(cache[edge[0]], edge[1])
		cache[edge[1]] = append(cache[edge[1]], edge[0])
	}

	visited := make([]bool, n)
	sum := n
	for i := 0; i < n; i++ {
		if !visited[i] {
			count := 0
			dfsCountPairs(cache, visited, i, &count)
			sum -= count
			res += count * sum
		}

	}

	return int64(res)
}

func dfsCountPairs(cache map[int][]int, visited []bool, currentIndex int, count *int) {

	if visited[currentIndex] {
		return
	}

	visited[currentIndex] = true
	*count++
	for _, next := range cache[currentIndex] {
		dfsCountPairs(cache, visited, next, count)
	}
}
