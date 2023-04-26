package DP

import "fmt"

func MinimumDeleteSum(s1 string, s2 string) int {
	return minimumDeleteSum(s1, s2)
}

func minimumDeleteSum(s1 string, s2 string) int {
	dp := make([][]int, len(s1)+1)

	for i := range dp {
		dp[i] = make([]int, len(s2)+1)

		if i == 0 {
			for j := 1; j < len(dp[i]); j++ {
				dp[i][j] = dp[i][j-1] + int(s2[j-1])
			}
		} else {
			dp[i][0] = dp[i-1][0] + int(s1[i-1])
		}
	}

	fmt.Println(dp)

	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}

	for i := 1; i <= len(s1); i++ {
		for j := 1; j <= len(s2); j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j]+int(s1[i-1]), dp[i][j-1]+int(s2[j-1]))
			}
		}
	}

	return dp[len(s1)][len(s2)]
}
