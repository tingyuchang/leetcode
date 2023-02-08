package jumpGame

func Jump(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	if nums[0] > len(nums)-1 {
		return 1
	}
	step := 1
	index := 0
	temp := nums[index+1 : index+1+nums[index]]

	maxIndex := 0
	for index+nums[index] < len(nums)-1 {
		max := 0
		for i, v := range temp {
			if i+v >= max {
				max = i + v
				maxIndex = i
			}
		}
		index = index + maxIndex + 1

		if index+nums[index] >= len(nums)-1 {
			step++
			break
		}
		temp = nums[index+1 : index+1+nums[index]]
		step++
	}
	return step
}
