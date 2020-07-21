package removeDuplicates

func RemoveDuplicates(nums []int) int {
	count := 0
	var current int

	for i,v := range nums {
		if i == 0 {
			current = v
			count++
		} else {
			if current < v {
				current = v
				count++
				nums[count-1] = v
			}
		}

		if i == len(nums) -1 {
			nums = nums[:count]
		}
	}

	return count
}