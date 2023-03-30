package Sort

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
	count := 0
	var current int

	for i, v := range nums {
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

		if i == len(nums)-1 {
			nums = nums[:count]
		}
	}

	return count
}
