package DP

import (
	"fmt"
	"sort"
)

func DeleteAndEarn(nums []int) int {
	return deleteAndEarn(nums)
}

func deleteAndEarn(nums []int) int {
	dp := make(map[int]int)
	keys := make([]int, 0)
	for _, v := range nums {
		dp[v] += v
	}

	for k, _ := range dp {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	// [2,3,4]
	// 4,9,4
	// like robbing house?

	ans := dp[keys[0]]
	current := 0
	fmt.Println(dp)
	for i := 1; i < len(keys); i++ {
		fmt.Println(ans, current)
		if keys[i] == keys[i-1]+1 {
			temp := ans
			ans = max(ans, current+dp[keys[i]])
			current = temp
		} else {
			current = ans
			ans += dp[keys[i]]
		}

	}

	return ans
}
