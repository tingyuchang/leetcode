package DP

func MaxValueOfCoins(piles [][]int, k int) int {
	return maxValueOfCoins(piles, k)
}

func maxValueOfCoins(piles [][]int, k int) int {
	dp := make([][]int, len(piles)+1)
	for i, _ := range dp {
		dp[i] = make([]int, k+1)
	}
	//
	//for i, _ := range dp {
	//	dp[i] = make([]int, len(piles[i]))
	//
	//	if len(dp[i]) == 0 {
	//		continue
	//	}
	//	dp[i][0] = piles[i][0]
	//	for j := 1; j < len(dp[i]); j++ {
	//		dp[i][j] = dp[i][j-1] + piles[i][j]
	//	}
	//}

	return dpMaxValueOfCoins(dp, piles, 0, k)
}

func dpMaxValueOfCoins(dp, piles [][]int, n, k int) int {
	if k == 0 || n == len(piles) {
		return 0
	}

	if dp[n][k] != 0 {
		return dp[n][k]
	}

	res := dpMaxValueOfCoins(dp, piles, n+1, k)
	current := 0

	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}

	max := func(a, b int) int {
		if a > b {
			return a
		}

		return b
	}

	for i := 0; i < min(len(piles[n]), k); i++ {
		current += piles[n][i]
		res = max(res, current+dpMaxValueOfCoins(dp, piles, n+1, k-i-1))
	}

	dp[n][k] = res
	return res
}
