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

func SearchInsertV2(nums []int, target int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			return i
		}

		if target < nums[i] {
			if i == 0 {
				return 0
			}

			if target > nums[i-1] {
				return i
			}
		}

		if i+1 >= len(nums) {
			return i + 1
		}
	}

	return 0
}
