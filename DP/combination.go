package DP

import "sort"

func CombinationSum4(nums []int, target int) int {
	sort.Ints(nums)
	dp := make([]int, target+1)
	dp[0] = 1

	for i := 1; i < len(dp); i++ {
		for _, v := range nums {
			if v > i {
				break
			}
			dp[i] = dp[i] + dp[i-v]
		}
	}

	return dp[target]
}
