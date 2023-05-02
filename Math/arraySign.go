package Math

func arraySign(nums []int) int {
	ans := 1
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			return 0
		} else if nums[i] < 0 {
			ans = ans * -1
		}
	}
	return ans

}
