package findMedianSortedArrays

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)

	if  m > n {
		nums1, nums2 = nums2, nums1
		m, n = n, m
	}

	iMin, iMax, half := 0, m, (m+n+1)/2

	for iMin <= iMax {
		i := (iMax+iMin)/2
		j := half - i

		if i < m && nums2[j-1] > nums1[i] {
			iMin = i + 1
		} else if i > 0 && nums1[i-1] > nums2[j] {
			iMax = i - 1
		} else {
			// correct i
			var maxLeft int
			var minRight int
			if i == 0 {
				maxLeft = nums2[j-1]
			} else if j == 0 {
				maxLeft = nums1[i-1]
			} else  {
				maxLeft = max(nums1[i-1], nums2[j-1])
			}

			if (m+n)%2 == 1 {
				return float64(maxLeft)
			}

			if i == m {
				minRight = nums2[j]
			} else if j == n {
				minRight = nums1[i]
			} else {
				minRight = min(nums1[i], nums2[j])
			}
			return (float64(maxLeft)+float64(minRight))/2.0
		}
	}

	return 0.0
}

func max (i, j int) int {
	if i > j {
		return i
	}

	return j
}

func min (i, j int) int {
	if i > j {
		return j
	}

	return i
}
/*

  # i is perfect

            if i == 0: max_of_left = B[j-1]
            elif j == 0: max_of_left = A[i-1]
            else: max_of_left = max(A[i-1], B[j-1])

            if (m + n) % 2 == 1:
                return max_of_left

            if i == m: min_of_right = B[j]
            elif j == n: min_of_right = A[i]
            else: min_of_right = min(A[i], B[j])

            return (max_of_left + min_of_right) / 2.0
 */