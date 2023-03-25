package DP

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