package Math

func IsMonotonic(nums []int) bool {
	return isMonotonic(nums)
}

func isMonotonic(nums []int) bool {
	if len(nums) <= 1 {
		return true
	}
	increasing := false

	start := 1

	for start < len(nums) {
		if nums[start-1] != nums[start] {
			increasing = nums[start-1] < nums[start]
			start++
			break
		}
		start++

	}

	for i := start; i < len(nums); i++ {
		if increasing && nums[i-1] > nums[i] {
			return false
		} else if !increasing && nums[i-1] < nums[i] {
			return false
		}
	}

	return true
}
