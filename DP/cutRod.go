package DP

import "math"

func MaxCutRodButtonUp(prices []int, n int) int {
	return maxCutRodMemoButtonUp(prices, n)
}

func MaxCutRodMemo(prices []int, n int) int {
	memo := make([]int, len(prices))
	return maxCutRodMemo(prices, memo, n)
}

func MaxCutRodRecursive(prices []int, n int) int {
	return maxCutRodV1(prices, n)
}

func maxCutRodMemoButtonUp(prices []int, n int) int {
	dp := make([]int, n+1)

	for i := 1; i < len(dp); i++ {
		q := math.MinInt
		for j := 1; j <= i; j++ {
			q = max(q, prices[j]+dp[i-j])
		}

		dp[i] = q
	}

	return dp[n]
}

func maxCutRodMemo(prices, memo []int, n int) int {
	if n == 0 {
		return 0
	}

	if memo[n] != 0 {
		return memo[n]
	}

	q := math.MinInt
	for i := 1; i <= n; i++ {
		q = max(q, prices[i]+maxCutRodMemo(prices, memo, n-i))
	}
	memo[n] = q
	return q

}

// O(2^n)
func maxCutRodV1(prices []int, n int) int {
	if n == 0 {
		return 0
	}
	q := math.MinInt

	for i := 1; i <= n; i++ {
		q = max(q, prices[i]+maxCutRodV1(prices, n-i))
	}
	return q
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
