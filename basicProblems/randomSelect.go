package basicProblems

// RandomizedSelect return s target-th element in nums
// using quick sort to implement partition and find target
func RandomizedSelect(nums []int, target int) int {
	return randomizedSelect(nums, 0, len(nums)-1, target)
}

func randomizedSelect(nums []int, p, r, target int) int {
	if p == r {
		return nums[p]
	}

	q := partitionRandomized(nums, p, r)
	// q means ith in partition, we need convert it to the position in entire array
	k := q - p + 1

	if target == k {
		return nums[q]
	}

	if k > target {
		return randomizedSelect(nums, p, q-1, target)
	} else {
		return randomizedSelect(nums, q+1, r, target-k)
	}
}

// partitionRandomized is the same in quick sort
func partitionRandomized(nums []int, p, r int) int {
	x := nums[r]
	i := p - 1

	for j := p; j < r; j++ {
		if nums[j] < x {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}

	i++
	nums[i], nums[r] = nums[r], nums[i]
	return i
}
