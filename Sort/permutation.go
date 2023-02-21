package Sort

func NextPermutation(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}

	if len(nums) == 1 {
		return nums
	}

	i := len(nums) - 1
	// find the largest index i such that nums[i-1] is less than nums[i]
	// if nums[i-1] smaller than nums[i], means we can find next permutation
	// If i is not the first index of the slice, then sub-slice nums[i…n) is stored in descending order
	// i.e. nums[i-1] < nums[i] >= nums[i+1] >= nums[i+2]
	for nums[i-1] >= nums[i] {
		i = i - 1
		if i == 0 {
			reverse(nums, 0)
			return nums
		}
	}

	j := len(nums) - 1
	// find highest index j such nums[j] is greater than nums[i-1]
	// i.e. 1,2,3,4,6,5 => 1,2,3,5,6,4
	for j > i && nums[j] <= nums[i-1] {
		j = j - 1
	}
	// swap nums[i-1] & nums[j]
	nums[i-1], nums[j] = nums[j], nums[i-1]
	// reverse nums[i:]
	reverse(nums, i)
	return nums
}

func reverse(nums []int, n int) {
	i, j := n, len(nums)-1
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}
