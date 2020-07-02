package threeSum

import "sort"

func ThreeSum(nums []int) [][]int {
	const target int = 0
	var results [][]int
	sort.Ints(nums)

	for i,v := range nums {
		if i > 0 && v == nums[i-1] {
			continue
		}

		l := i+1
		r := len(nums)-1

		for l < r {
			sum := v + nums[l] + nums[r]
			if sum > target {
				r--
			} else if sum < target {
				l++
			} else {
				results = append(results, []int{v, nums[l], nums[r]})
				l++
				for nums[l] == nums[l-1] && l < r {
					l++
				}
			}
		}

	}

	return results
}