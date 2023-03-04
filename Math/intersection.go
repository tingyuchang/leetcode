package Math

func Intersection(nums1 []int, nums2 []int) []int {
	res := make([]int, 0)
	counter := make(map[int]bool)

	for i := 0; i < len(nums1); i++ {
		counter[nums1[i]] = true
	}

	for i := 0; i < len(nums2); i++ {
		if _, ok := counter[nums2[i]]; ok {
			res = append(res, nums2[i])
			delete(counter, nums2[i])
		}
	}

	return res
}