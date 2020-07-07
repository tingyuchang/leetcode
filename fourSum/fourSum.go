package fourSum

import (
	"sort"
)

func FourSum(nums []int, target int) [][]int {
	var result [][]int
	sort.Ints(nums)

	for i,v := range nums {
		if i > 0 && v == nums[i-1] {
			continue
		}

		temp := threeSum(nums[i+1:], target-v)
		if len(temp) > 0 {
			for _,s := range temp {
				s = append(s, v)
				result = append(result, s)
			}
		}
	}

	return  result
}

func threeSum(nums []int, target int) [][]int {
	var results [][]int


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