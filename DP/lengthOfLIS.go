package DP

import "fmt"

func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	res := 1
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	dp[0] = 1
	for i := 1; i < len(nums); i++ {
		dp[i] = 1
		for j := i - 1; j >= 0; j-- {

			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
			if dp[i] > res {
				res = dp[i]
			}

		}
	}
	fmt.Println(dp)
	return res
}
