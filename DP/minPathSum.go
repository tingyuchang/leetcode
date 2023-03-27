package DP

func MinPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([][]int, m+1)

	for i, _ := range dp {
		dp[i] = make([]int, n+1)
	}

	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if i == 1 {
				dp[i][j] = grid[i-1][j-1] + dp[i][j-1]
			} else if j == 1 {
				dp[i][j] = grid[i-1][j-1] + dp[i-1][j]
			} else {
				dp[i][j] = grid[i-1][j-1] + min(dp[i-1][j], dp[i][j-1])
			}

		}
	}

	return dp[m][n]
}
