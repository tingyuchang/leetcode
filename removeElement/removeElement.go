package removeElement

import "sort"

func RemoveElement(nums []int, val int) int {
	sort.Ints(nums)
	start := 0
	end := -1

	for i,v := range nums {
		if v == val {
			if start == 0 && end == -1 {
				start = i
			}
			end = i
		}
	}

	nums = append(nums[:start], nums[end+1:]...)

	return len(nums)
}