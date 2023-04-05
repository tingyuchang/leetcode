package greedy

func MinimizeArrayValue(nums []int) int {
	return minimizeArrayValue(nums)
}

func minimizeArrayValue(nums []int) int {
	res, sum := 0, 0

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		// sum +i: is because mod
		res = max(res, (sum+i)/(i+1))
	}

	return res
}
