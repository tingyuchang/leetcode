package traversal

func MinReorder(n int, connections [][]int) int {
	res := 0
	cache := make(map[int][]int)
	reverseCache := make(map[int][]int)

	for _, v := range connections {
		cache[v[0]] = append(cache[v[0]], v[1])
		reverseCache[v[1]] = append(reverseCache[v[1]], v[0])
	}

	queue := make([]int, 1)
	queue[0] = 0

	visited := make([]bool, n)

	for len(queue) > 0 {
		city := queue[0]
		queue = queue[1:]

		visited[city] = true
		// add cities need reorder
		for _, v := range cache[city] {
			if visited[v] {
				continue
			}
			queue = append(queue, v)
			res++
		}
		// add cities don't need reorder

		for _, v := range reverseCache[city] {
			if visited[v] {
				continue
			}
			queue = append(queue, v)
		}

	}

	return res
}
