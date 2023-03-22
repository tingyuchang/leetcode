package traversal

import (
	"math"
)

func MinScore(n int, roads [][]int) int {
	res := math.MaxInt
	cache := make(map[int][]map[int]int)

	for _, r := range roads {
		cache[r[0]] = append(cache[r[0]], map[int]int{
			r[1]: r[2],
		})
		cache[r[1]] = append(cache[r[1]], map[int]int{
			r[0]: r[2],
		})
	}
	visited := make([]bool, n+1)
	dfsMinScore(cache, visited, 1, &res)
	return res
}

func dfsMinScore(cache map[int][]map[int]int, visited []bool, current int, res *int) int {
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}

	visited[current] = true
	for _, v := range cache[current] {
		for des, dis := range v {
			*res = min(*res, dis)
			if visited[des] == false {
				visited[des] = true
				*res = min(*res, dfsMinScore(cache, visited, des, res))
			}
		}
	}

	return *res
}
