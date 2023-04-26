package DP

import "fmt"

func NumDistinct(s string, t string) int {
	return numDistinct(s, t)
}

func numDistinct(s string, t string) int {
	dp := make([][]int, len(t)+1)

	for i := range dp {
		dp[i] = make([]int, len(s)+1)

		if i == 0 {
			for j := range dp[i] {
				// target string is "", so no matter input string is, we only have 1 result
				dp[i][j] = 1
			}
		} else {
			// input string is "",  so 0 result
			dp[i][0] = 0
		}
	}
	fmt.Println(dp)
	for i := 1; i <= len(t); i++ {
		for j := 1; j <= len(s); j++ {
			if t[i-1] == s[j-1] {
				dp[i][j] = dp[i-1][j-1] + dp[i][j-1]
			} else {
				dp[i][j] = dp[i][j-1]
			}
		}
	}
	fmt.Println(dp)
	return dp[len(t)][len(s)]
}
