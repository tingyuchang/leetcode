package traversal

func NumberOfPaths(grid [][]int, k int) int {
	return numberOfPaths(grid, k)
}

func numberOfPaths(grid [][]int, k int) int {
	m, n := len(grid), len(grid[0])
	dp := make([][][]int, m)
	for i := range dp {
		dp[i] = make([][]int, n)

	}
	return 0
}
