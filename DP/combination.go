package DP

import (
	"sort"
)

func CombinationSum3(k int, n int) [][]int {
	res := make([][]int, 0)
	currentRes := make([]int, 0)
	combinationSum3(k, n, 0, 1, &res, &currentRes)
	return res
}

func combinationSum3(k, n, currentSum, currentIndex int, res *[][]int, currentRes *[]int) {
	if currentSum > n {
		return
	}

	if currentSum == n && k == 0 {
		temp := make([]int, len(*currentRes))
		copy(temp[:], (*currentRes)[:])
		*res = append(*res, temp)
	}

	for i := currentIndex; i <= 9; i++ {
		currentSum += i
		*currentRes = append(*currentRes, i)
		combinationSum3(k-1, n, currentSum, i+1, res, currentRes)
		currentSum -= i
		*currentRes = (*currentRes)[:len(*currentRes)-1]
	}
}

func CombinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	res := make([][]int, 0)
	currentRes := make([]int, 0)
	combinationSum2(candidates, target, 0, 0, &res, &currentRes)
	return res
}

func combinationSum2(candidates []int, target, currentIndex, currentSum int, res *[][]int, currentRes *[]int) {
	if currentSum > target {
		return
	}

	if currentSum == target {
		temp := make([]int, len(*currentRes))
		copy(temp[:], (*currentRes)[:])
		*res = append(*res, temp)
	}
	previous := 0
	for i := currentIndex; i < len(candidates); i++ {
		if candidates[i] == previous {
			continue
		}
		currentSum += candidates[i]
		*currentRes = append(*currentRes, candidates[i])
		combinationSum2(candidates, target, i+1, currentSum, res, currentRes)
		currentSum -= candidates[i]
		*currentRes = (*currentRes)[:len(*currentRes)-1]
		previous = candidates[i]
	}
}

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

func Combination(n, k int) [][]int {
	res := make([][]int, 0)
	currentRes := make([]int, 0)

	dfsCombination(n, k, 1, &res, &currentRes)

	return res
}

func dfsCombination(n, k, currentIndex int, res *[][]int, currentRes *[]int) {
	if len(*currentRes) > k {
		return
	}

	if len(*currentRes) == k {
		temp := make([]int, k)
		copy(temp[:], (*currentRes)[:])
		*res = append(*res, temp)
	}

	for i := currentIndex; i <= n; i++ {
		*currentRes = append(*currentRes, i)
		dfsCombination(n, k, i+1, res, currentRes)
		*currentRes = (*currentRes)[:len(*currentRes)-1]
	}

}
