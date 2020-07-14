package findMedianSortedArrays

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	l :=  m+ n

	index1 := 0
	index2 := 0
	start := 0
	var temp [2]int

	for start <= l/2 && index1 < m && index2 < n {
		if nums1[index1] < nums2[index2] {
			temp[0], temp[1] = temp[1], nums1[index1]
			index1++
		} else {
			temp[0], temp[1] = temp[1], nums2[index2]
			index2++
		}
		start++
	}

	if start <= l/2 && index1 < m {
		if l/2 - start >=1 {
			temp[0], temp[1] = nums1[index1+l/2-start-1], nums1[index1+l/2-start]
		} else {
			temp[0], temp[1] = temp[1], nums1[index1]
		}
	}

	if start <= l/2 && index2 < n  {
		if l/2 - start >=1 {
			temp[0], temp[1] = nums2[index2+l/2-start-1], nums2[index2+l/2-start]
		} else {
			temp[0], temp[1] = temp[1], nums2[index2]
		}
	}

	if l%2 == 0 {
		return float64(temp[len(temp)-1] + temp[len(temp)-2])/2.0
	}

	return float64(temp[len(temp)-1])
}
