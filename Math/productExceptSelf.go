package Math

func ProductExceptSelf(nums []int) []int {
	result := make([]int, len(nums))
	sum := 1
	m := make(map[int]bool)
	for i, v := range nums {
		if v != 0 {
			sum = sum * v
		} else {
			m[i] = true
		}
	}

	if len(nums) == len(m) {
		sum = 0
	}

	for i := 0; i < len(nums); i++ {

		if len(m) == 0 {
			result[i] = sum / nums[i]
		} else {
			_, ok := m[i]

			if ok && len(m) == 1 {
				result[i] = sum
			} else {
				result[i] = 0
			}
		}

	}

	return result
}
