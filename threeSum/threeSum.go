package threeSum

import "sort"

func threeSum(nums []int) [][]int {
	var result [][]int
	sort.Ints(nums)

	for i, v := range nums {
		if i > 0 && v == nums[i-1] {
			continue
		}

		temp := TwoSum(nums[i:], -v)
		if len(temp) == 2 {
			temp = append([]int{v}, temp...)
			result = append(result, temp)
		}
	}

	return result

}

func TwoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for i, v := range nums {
		_, n := m[v]

		if n {
			return []int{m[v], i}
		} else {
			m[target-v] = i
		}
	}

	return nil
}

func ThreeSum(nums []int) [][]int {
	const target int = 0
	var results [][]int
	sort.Ints(nums)

	for i, v := range nums {
		if i > 0 && v == nums[i-1] {
			continue
		}

		l := i + 1
		r := len(nums) - 1

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
