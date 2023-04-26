package Array

func RemoveElement(nums []int, val int) int {

	notValCount := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[i], nums[notValCount] = nums[notValCount], nums[i]
			notValCount++
		}
	}

	return notValCount
}
