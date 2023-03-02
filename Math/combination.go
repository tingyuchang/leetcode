package Math

import "sort"

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
