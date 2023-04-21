package DP

import "fmt"

func Kanspack(prices, weights []int, maxWeight int) int {
	dp := make([]int, maxWeight+1)
	for i := 0; i < len(prices); i++ {
		p, w := prices[i], weights[i]
		for j := len(dp) - 1; j >= 0; j-- {
			if w <= j {
				dp[j] = max(dp[j], dp[max(0, j-w)]+p)
			}
			fmt.Println(dp)
		}
	}

	return dp[maxWeight]
}
