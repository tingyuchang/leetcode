package Array

import "math"

func minSubArrayLen(target int, nums []int) int {
	l, r := 0, 0
	sum := 0
	ans := math.MaxInt
	for r < len(nums) {
		sum += nums[r]
		for sum >= target {
			if r-l+1 < ans {
				ans = r - l + 1
			}
			sum -= nums[l]
			l++
		}
		r++
	}

	if ans == math.MaxInt {
		return 0
	}

	return ans
}
