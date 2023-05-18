package graph

func findSmallestSetOfVertices(n int, edges [][]int) []int {
	edgeMaps := make(map[int]struct{})

	for _, edge := range edges {
		to := edge[1]

		if _, ok := edgeMaps[to]; !ok {
			edgeMaps[to] = struct{}{}
		}
	}

	ans := make([]int, 0)
	for i := 0; i < n; i++ {
		if _, ok := edgeMaps[i]; !ok {
			ans = append(ans, i)
		}
	}

	return ans
}
