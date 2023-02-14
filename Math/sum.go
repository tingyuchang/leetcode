package Math

func ThreeSum(nums []int) [][]int {
	result := make([][]int, 0)

	for i := 0; i < len(nums); i++ {
		others := make([]int, 0)

		for j := 0; j < len(nums); j++ {
			if j == i {
				continue
			}
			others = append(others, nums[j])
		}
		temp := TwoSumV2(others, -nums[i])

		if len(temp) == 2 {
			result = append(result, []int{nums[i], temp[0], temp[1]})
		}
	}
	return result
}

func TwoSumV2(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		_, ok := m[v]

		if ok {
			return []int{nums[m[v]], nums[i]}
		} else {
			m[target-v] = i
		}
	}
	return nil
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
