package graph

import "math"

func minScore(n int, roads [][]int) int {
	uf := NewUnionFind(n)

	ans := math.MaxInt
	for i := 0; i < len(roads); i++ {
		n1 := roads[i][0] - 1
		n2 := roads[i][1] - 1
		uf.join(n1, n2)
	}

	for i := 0; i < len(roads); i++ {
		if uf.find(0) == uf.find(roads[i][0]-1) {
			if roads[i][2] < ans {
				ans = roads[i][2]
			}
		}

	}

	return ans
}
