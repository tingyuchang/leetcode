package Math

import "sort"

func CheckArithmeticSubarrays(nums []int, l []int, r []int) []bool {
	return checkArithmeticSubarrays(nums, l, r)
}

func checkArithmeticSubarrays(nums []int, l []int, r []int) []bool {
	// [4,6,5,9,3,7], l = [0,0,2], r = [2,3,5]
	// [4,6,5]
	// [4,6,5,9]
	// [5,9,3,7]

	res := make([]bool, 0)
	for i := 0; i < len(l); i++ {
		res = append(res, checkArithmetic(nums[l[i]:r[i]+1]))
	}
	return res

}

func checkArithmetic(nums []int) bool {
	if len(nums) <= 1 {
		return true
	}
	temp := make([]int, len(nums))
	for i, v := range nums {
		temp[i] = v
	}

	sort.Ints(temp)
	a := temp[1] - temp[0]

	for i := 2; i < len(temp); i++ {
		if temp[i]-a != temp[i-1] {
			return false
		}
	}

	return true

}
