package Math

import (
	"fmt"
	"sort"
)

func CountSubarrays(nums []int, minK int, maxK int) int64 {
	var res int64
	l := -1
	lastMin, lastMax := -1, -1

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	for i := 0; i < len(nums); i++ {

		if nums[i] >= minK && nums[i] <= maxK {
			if nums[i] == minK {
				lastMin = i
			}
			if nums[i] == maxK {
				lastMax = i
			}

			res += int64(max(0, min(lastMin, lastMax)-l))
			fmt.Println(res)
		} else {
			l = i
			lastMin, lastMax = -1, -1
		}
	}

	return res
}

func CombinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	res := make([][]int, 0)
	current := make([]int, 0)
	combination(candidates, target, 0, 0, &res, &current)
	return res
}

func combination(candidates []int, target, currentSum, currentIndex int, res *[][]int, current *[]int) {
	if currentSum > target {
		return
	}
	if currentSum == target {
		temp := make([]int, len(*current))
		copy(temp[:], (*current)[:])
		*res = append(*res, temp)
		return
	}

	for i := currentIndex; i < len(candidates); i++ {
		*current = append(*current, candidates[i])
		currentSum += candidates[i]

		combination(candidates, target, currentSum, i, res, current)
		*current = (*current)[:len(*current)-1]
		currentSum -= candidates[i]
	}

}
