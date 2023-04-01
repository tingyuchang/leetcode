package Math

func PivotIndex(nums []int) int {
	return pivotIndex(nums)
}

func pivotIndex(nums []int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}

	cache := make([]int, len(nums))
	cache[0] = nums[0]

	if sum-cache[0] == 0 {
		return 0
	}

	for i := 1; i < len(cache); i++ {
		cache[i] = cache[i-1] + nums[i]
		if cache[i-1] == sum-cache[i] {
			return i
		}
	}

	return -1
}
