package __Daily_Prac

var MOD int = 1e9 + 7

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a > b {
		return b
	}

	return a
}

func Reverse(nums []int, n int) {
	l := n
	r := len(nums) - 1

	for r > l {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}

}
