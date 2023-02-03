package searchInsertPosition

func SearchInsert(nums []int, target int) int {
	n := len(nums)
	l := 0
	r := n - 1

	for r >= l {
		mid := (r + l) / 2

		if nums[mid] > target {
			r = mid - 1
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			return mid
		}
	}
	return l
}
