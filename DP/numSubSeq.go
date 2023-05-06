package DP

import (
	"sort"
)

func NumSubseq(nums []int, target int) int {
	return numSubseq(nums, target)
}

func numSubseq(nums []int, target int) int {
	n := len(nums)
	sort.Ints(nums)
	ans := 0

	power := make([]int, n)
	power[0] = 1

	for i := 1; i < n; i++ {
		power[i] = (power[i-1] * 2) % MOD
	}

	for minVal := 0; minVal < n; minVal++ {

		l := 0
		r := n - 1

		for r >= l {
			mid := int(uint(l+r) >> 1)
			if nums[mid] <= target-nums[minVal] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}

		maxVal := r

		if maxVal >= minVal {
			ans = (ans + power[maxVal-minVal]) % MOD
		}
	}

	return ans
}
