package Search

func FindMinInRotatedArray(nums []int) int {
	if nums[0] <= nums[len(nums)-1] {
		return nums[0]
	}
	l := 0
	r := len(nums) - 1
	pivot := 0
	for r >= l {
		mid := (l + r) / 2
		if nums[mid] > nums[mid+1] {
			pivot = mid + 1
			break
		}

		if nums[mid] > nums[l] {
			l = mid
		} else {
			r = mid
		}
	}

	return nums[pivot]
}

func SearchRotatedArray(nums []int, target int) int {
	if nums[0] < nums[len(nums)-1] || len(nums) == 1 {
		return BinarySearch(nums, target)
	}
	l := 0
	r := len(nums) - 1
	pivot := 0
	for r >= l {
		mid := (l + r) / 2
		if nums[mid] < nums[l] {
			r = r - 1
		} else if nums[mid] > nums[r] {
			l = l + 1
		} else {
			pivot = l
			break
		}
	}

	if target > nums[pivot] && target <= nums[len(nums)-1] {
		l = pivot
		r = len(nums) - 1
	} else if target <= nums[pivot-1] && target >= nums[0] {
		l = 0
		r = pivot
	} else {
		if nums[pivot] == target {
			return pivot
		}
		return -1
	}

	for r >= l {
		if r >= l {
			mid := (l + r) / 2
			if nums[mid] > target {
				r = mid - 1
			} else if nums[mid] < target {
				l = mid + 1
			} else {
				return mid
			}
		}
	}

	return -1
}

func BinarySearch(nums []int, target int) int {
	l := 0
	r := len(nums) - 1
	for r >= l {
		mid := (l + r) / 2
		if nums[mid] > target {
			r = mid - 1
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			return mid
		}
	}

	return -1
}
