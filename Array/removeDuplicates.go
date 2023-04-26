package Array

import "fmt"

func RemoveDuplicatesII(nums []int) int {
	return removeDuplicatesII(nums)
}

func removeDuplicatesII(nums []int) int {
	count := 0
	var current int
	buffer := 1
	for i, v := range nums {
		if i == 0 {
			current = v
			count++
		} else {
			if current != v {
				current = v
				count++
				nums[count-1] = v
				buffer = 1

			} else {
				buffer--
				if buffer == 0 {
					count++
					nums[count-1] = current
				}
			}
		}
		if i == len(nums)-1 {
			nums = nums[:count]
		}
	}

	fmt.Println(nums)

	return count
}

func RemoveDuplicates(nums []int) int {
	nonDuplicated := 0

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[nonDuplicated] {
			nonDuplicated++
			nums[i], nums[nonDuplicated] = nums[nonDuplicated], nums[i]

		}
	}
	// nums = nums[:nonDuplicated+1]
	// fmt.Println(nums,  nonDuplicated)
	return nonDuplicated + 1
}
