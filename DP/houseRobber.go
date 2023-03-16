package DP

import "math"

func Rob(nums []int) int {
	// return rob(nums, len(nums)-1)
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	//dp := make([]int, len(nums)+1)
	//dp[0] = 0
	//dp[1] = nums[0]
	//
	//for i := 1; i < len(nums); i++ {
	//	v := nums[i]
	//	dp[i+1] = max(dp[i], dp[i-1]+v)
	//}
	//
	//return dp[len(nums)]
	if len(nums) == 0 {
		return 0
	}
	temp1 := 0
	temp2 := 0

	for i := 0; i < len(nums); i++ {
		temp := temp1
		temp1 = max(temp1, temp2+nums[i])
		temp2 = temp
	}

	return temp1
}

func rob(nums []int, n int) int {
	if n < 0 {
		return 0
	}
	return int(math.Max(float64(nums[n]+rob(nums, n-2)), float64(rob(nums, n-1))))
}
