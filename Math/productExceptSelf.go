package Math

func ProductExceptSelf(nums []int) []int {
	reslut := make([]int, len(nums))
	prefix := make([]int, len(nums))
	suffix := make([]int, len(nums))
	prefix[0] = 1
	suffix[len(suffix)-1] = 1

	for i := 1; i < len(nums); i++ {
		prefix[i] = prefix[i-1] * nums[i-1]
	}

	for i := len(nums) - 2; i >= 0; i-- {
		suffix[i] = suffix[i+1] * nums[i+1]
	}

	for i := 0; i < len(nums); i++ {
		reslut[i] = prefix[i] * suffix[i]
	}

	return reslut
}

func ProductExceptSelfSlow(nums []int) []int {
	result := make([]int, len(nums))
	sum := 1
	isZero := false
	isMultipleZero := false
	indexZero := 0
	for i, v := range nums {
		if v != 0 {
			sum = sum * v
		} else if isZero {
			isMultipleZero = true
		} else {
			isZero = true
			indexZero = i
		}
	}

	for i := 0; i < len(nums); i++ {

		if !isZero {
			result[i] = sum / nums[i]
		} else {
			if isMultipleZero {
				result[i] = 0
			} else {
				if i == indexZero {
					result[i] = sum
				} else {
					result[i] = 0
				}
			}
		}
	}

	return result
}
