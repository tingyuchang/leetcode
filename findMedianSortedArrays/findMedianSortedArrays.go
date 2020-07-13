package findMedianSortedArrays

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	medium := len(nums1) + len(nums2)

	index1 := 0
	index2 := 0
	var temp []int


	for i := 0; i <= medium/2; i ++ {
		if index1+1 > len(nums1) {
			temp = append(temp, nums2[index2:index2+ medium/2+1-i]...)
			break
		}

		if index2+1 > len(nums2) {
			temp = append(temp, nums1[index1:index1+ medium/2+1-i]...)
			break
		}

		if nums1[index1] < nums2[index2] {
			temp = append(temp, nums1[index1])
			index1++
		} else {
			temp = append(temp, nums2[index2])
			index2++
		}

	}

	if medium%2 == 0 {
		return (float64(temp[len(temp)-1])+float64(temp[len(temp)-2]))/2.0
	}

	return float64(temp[len(temp)-1])
}