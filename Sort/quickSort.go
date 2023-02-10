package Sort

func QuickSort(nums []int, p, r int) {
	if p < r {
		q := partition(nums, p, r)
		QuickSort(nums, p, q-1)
		QuickSort(nums, q+1, r)
	}
}

func partition(nums []int, p, r int) int {
	x := nums[r]
	i := p - 1
	for j := p; j < r; j++ {
		if nums[j] <= x {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[i+1], nums[r] = nums[r], nums[i+1]
	return i + 1
}

// RandomizedSelect finds ith element in unsorted array
func RandomizedSelect(nums []int, p, r, i int) int {
	if p == r {
		return nums[p]
	}

	q := partition(nums, p, r)
	// k means in [p:r] q's position from p
	k := q - p + 1
	if i == k {
		return nums[q]
	} else if k > i {
		return RandomizedSelect(nums, p, q-1, i)
	} else {
		return RandomizedSelect(nums, q+1, r, i-k)
	}
}

func Select(nums []int, p, r, i int) int {
	// remove elements from array until n %5 == 0
	for (r-p+1)%5 != 0 {
		for j := p + 1; j < r; j++ {
			if nums[p] > nums[j] {
				nums[p], nums[j] = nums[j], nums[p]
			}
		}
		// if we want minimum. then done
		if i == 1 {
			return nums[p]
		}
		p++
		i--
	}

	g := (r - p + 1) / 5

	for j := p; j < p+g-1; j++ {
		// insertionSort
		for k := 1; k < 5; k++ {
			key := nums[k*g+j]
			for l := k - 1; l >= 0; l-- {
				if nums[l*g+j] > key {
					nums[(l+1)*g+j] = nums[l*g+j]
					if l == 0 {
						nums[l] = key
					}
				} else {
					nums[(l+1)*g+j] = key
					break
				}
			}
		}
	}

	x := Select(nums, p+2*g, p+3*g-1, g/2)
	q := partitionAround(nums, p, r, x)
	k := q - p + 1
	if i == k {
		return nums[q]
	} else if k > i {
		return Select(nums, p, q-1, i)
	} else {
		return Select(nums, q+1, r, i-k)
	}
}

// implementit later
func partitionAround(nums []int, p, r, x int) int {
	return 0
}
