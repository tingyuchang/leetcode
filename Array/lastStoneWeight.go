package Array

func LastStoneWeight(stones []int) int {
	return lastStoneWeight(stones)
}

func lastStoneWeight(stones []int) int {
	for i := len(stones) / 2; i >= 0; i-- {
		maxHeap(stones, i)
	}
	for len(stones) > 1 {
		s1 := stones[0]
		stones[0], stones[len(stones)-1] = stones[len(stones)-1], stones[0]
		stones = stones[:len(stones)-1]
		maxHeap(stones, 0)
		s2 := stones[0]

		if s1 != s2 {
			// replace 0 th
			stones[0] = abs(s1 - s2)
		} else {
			// remove 0-th
			stones[0], stones[len(stones)-1] = stones[len(stones)-1], stones[0]
			stones = stones[:len(stones)-1]
		}

		maxHeap(stones, 0)
	}

	if len(stones) == 1 {
		return stones[0]
	}

	return 0
}

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

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
