package DP

func MaxSubArray(nums []int) int {
	res, temp := 0, 0

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for _, v := range nums {
		temp = max(v, temp+v)
		res = max(res, temp)
	}

	return res
}
