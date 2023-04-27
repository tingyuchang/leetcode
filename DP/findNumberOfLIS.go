package DP

func FindNumberOfLIS(nums []int) int {
	return findNumberOfLIS(nums)
}

func findNumberOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	count := make([]int, len(nums))
	ans := 0
	longest := 0
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		count[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				if dp[i] <= dp[j] {
					dp[i] = dp[j] + 1
					count[i] = count[j]
				} else if dp[i]-1 == dp[j] {
					count[i] += count[j]
				}
			}
		}

		if longest == dp[i] {
			ans += count[i]
		}
		if dp[i] > longest {
			longest = dp[i]
			ans = count[i]
		}

	}
	return ans
}
