package mergeSortedArray

func Merge(nums1 []int, m int, nums2 []int, n int) []int {
	// like merge sort, set up initial index
	i, j := 0, 0

	// if m <= 0, skip this loop
	for i < len(nums1) && j < len(nums2) && m > 0 {
		if nums1[i] > nums2[j] {
			// re-sort nums1 from latest element to current
			for k := len(nums1) - 1; k > i; k-- {
				nums1[k], nums1[k-1] = nums1[k-1], nums1[k]
			}
			nums1[i] = nums2[j]
			j++
			// why m++, is because nums1 add new element from nums2
			m++
		} else {
			// why m--, we use m to check that we need run this loop or not
			// m == 0 means activate element in nums1 are all used
			i++
			m--
		}
	}

	if j < len(nums2) {
		for k := i; k < len(nums1); k++ {
			nums1[k] = nums2[j]
			j++
		}
	}

	return nums1
}

/*
1,2,3,0,0,0
4,5,6

*/
