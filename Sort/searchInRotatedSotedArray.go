package Sort

func SearchInRotatedII(nums []int, target int) bool {
	l := 0
	r := len(nums) - 1

	for r >= l {
		mid := int(uint(l+r) >> 1)

		if nums[mid] == target {
			return true
		}

		if nums[l] == nums[mid] && nums[r] == nums[mid] {
			l++
			r--
		} else if nums[l] <= nums[mid] {
			if nums[l] <= target && nums[mid] > target {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			if nums[mid] < target && nums[r] >= target {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}

	return false

}
