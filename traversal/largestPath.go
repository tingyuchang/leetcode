package traversal

func LargestPathValue(colors string, edges [][]int) int {
	return largestPathValue(colors, edges)
}
func largestPathValue(colors string, edges [][]int) int {
	res := 0

	// source -> []destinations
	cache := make(map[int][]int, 0)
	dp := make([]map[byte]int, len(colors))
	visited := make([]bool, len(colors))

	for _, edge := range edges {
		cache[edge[0]] = append(cache[edge[0]], edge[1])
	}

	for i := 0; i < len(colors); i++ {
		if isCycle, temp := dfsLargestPathValue(i, cache, colors, visited, dp); isCycle {
			return -1
		} else {
			for _, v := range temp {
				if v > res {
					res = v
				}
			}
		}
	}

	return res
}

func dfsLargestPathValue(start int, cache map[int][]int, colors string, visited []bool, dp []map[byte]int) (bool, map[byte]int) {
	if len(cache[start]) == 0 {
		return false, map[byte]int{colors[start]: 1}
	}

	// dp?

	if len(dp[start]) > 0 {
		return false, dp[start]
	}

	visited[start] = true

	res := make(map[byte]int)

	for _, next := range cache[start] {
		if visited[next] {
			// cycle detected
			return true, nil
		}

		isCycle, temp := dfsLargestPathValue(next, cache, colors, visited, dp)

		if isCycle {
			return true, nil
		}
		for k, v := range temp {
			if v > res[k] {
				res[k] = v
			}
		}
	}

	res[colors[start]]++
	visited[start] = false
	dp[start] = res
	return false, res
}
