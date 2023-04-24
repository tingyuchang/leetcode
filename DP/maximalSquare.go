package DP

func MaximalSquare(matrix [][]byte) int {
	return maximalSquare(matrix)
}

func maximalSquare(matrix [][]byte) int {
	m, n := len(matrix), len(matrix[0])
	dp := make([][]int, m+1)
	for i, _ := range dp {
		dp[i] = make([]int, n+1)
	}

	ans := 0
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if matrix[i-1][j-1] == 1 {
				dp[i][j] = minFor3(dp[i-1][j-1], dp[i-1][j], dp[i][j-1]) + 1

				ans = max(ans, dp[i][j])
			}
		}
	}

	return ans * ans
}

func minFor3(a, b, c int) int {
	var smallest int

	if a > b {
		smallest = b
	} else {
		smallest = a
	}

	if c < smallest {
		smallest = c
	}

	return smallest
}
