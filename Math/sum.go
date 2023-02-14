package Math

import (
	"sort"
)

func ThreeSum(nums []int) [][]int {
	result := make([][]int, 0)
	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		l := i + 1
		r := len(nums) - 1

		for r > l {
			sum := nums[i] + nums[l] + nums[r]

			if sum > 0 {
				r--
			} else if sum < 0 {
				l++
			} else {
				// get it
				result = append(result, []int{nums[i], nums[l], nums[r]})
				l++
				// skip same value in l
				for l < r && nums[l] == nums[l-1] {
					l++
				}
			}
		}

	}

	return result
}

func TwoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		_, ok := m[v]

		if ok {
			return []int{m[v], i}
		} else {
			m[target-v] = i
		}
	}
	return nil
}
