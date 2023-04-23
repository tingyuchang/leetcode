package DP

// 63 Unique PathII
func UniquePathsWithObstacles(obstacleGrid [][]int) int {
	return uniquePathsWithObstacles(obstacleGrid)
}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid)-1, len(obstacleGrid[0])-1
	dp := make([][]int, m+1)

	for i, _ := range dp {
		dp[i] = make([]int, n+1)

		if i == 0 {
			if obstacleGrid[0][0] == 0 {
				dp[0][0] = 1
			}

			for j := 1; j <= n; j++ {
				if obstacleGrid[i][j] == 0 && dp[i][j-1] == 1 {
					dp[i][j] = 1
				} else {
					break
				}
			}
		} else {
			if obstacleGrid[i][0] == 0 && dp[i-1][0] == 1 {
				dp[i][0] = 1
			}
		}
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if obstacleGrid[i][j] != 1 {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}

	return dp[m][n]
}

func UniquePaths(m int, n int) int {
	dp := make([][]int, m)

	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n)

		for j := 0; j < len(dp[i]); j++ {
			if i == 0 {
				dp[i][j] = 1
				continue
			}

			if j == 0 {
				dp[i][j] = 1
				continue
			}

			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	return dp[m-1][n-1]
}
