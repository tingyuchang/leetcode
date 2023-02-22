package BinarySearch

func SearchRange(nums []int, target int) []int {

	//start, end := -1, -1
	//
	//for i := 0; i < len(nums); i++ {
	//	v := nums[i]
	//	if v == target {
	//		if start == -1 {
	//			start = i
	//		}
	//		end = i
	//	}
	//
	//	if v > target {
	//		break
	//	}
	//}
	//return []int{start, end}
	l := 0
	r := len(nums) - 1

	start, end := -1, -1

	for r >= l {
		mid := (l + r) / 2

		if nums[mid] > target {
			r = mid - 1
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			start, end = mid, mid
			for {
				if start != 0 && nums[start-1] == target {
					start--
				}

				if end < len(nums)-1 && nums[end+1] == target {
					end++
				}
				if (start == 0 || nums[start-1] != target) && (end == len(nums)-1 || nums[end+1] != target) {
					break
				}
			}

			break
		}
	}

	return []int{start, end}

}
