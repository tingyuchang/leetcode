package Sort

func maxHeap(nums []int, n int) {
	l := ((n + 1) << 1) - 1
	r := (n + 1) << 1
	var largest int
	if l < len(nums) && nums[l] > nums[n] {
		largest = l
	} else {
		largest = n
	}

	if r < len(nums) && nums[r] > nums[largest] {
		largest = r
	}

	if largest != n {
		nums[n], nums[largest] = nums[largest], nums[n]
		maxHeap(nums, largest)
	}

}

func BuildHeap(nums []int) []int {
	for i := len(nums) / 2; i >= 0; i-- {
		maxHeap(nums, i)
	}

	return nums
}

func HeapSort(nums []int) []int {
	nums = BuildHeap(nums)
	n := len(nums)
	for i := n - 1; i > 0; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		nums = nums[:i]
		maxHeap(nums, 0)
	}
	return nums[:n]
}
