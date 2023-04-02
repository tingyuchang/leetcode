package Sort

func SortedSquares(nums []int) []int {
	return sortedSquares(nums)
}

func sortedSquares(nums []int) []int {
	n := len(nums)
	l := 0
	r := n - 1
	res := make([]int, len(nums))
	for i := n - 1; i >= 0; i-- {
		if square(nums[r]) > square(nums[l]) {
			res[i] = square(nums[r])
			r--
		} else {
			res[i] = square(nums[l])
			l++
		}
	}

	return res
}

func square(n int) int {
	return n * n
}
