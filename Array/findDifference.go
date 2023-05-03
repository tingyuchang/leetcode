package Array

func FindDifference(nums1 []int, nums2 []int) [][]int {
	return findDifference(nums1, nums2)
}

func findDifference(nums1 []int, nums2 []int) [][]int {

	nums1Map := make(map[int]struct{})
	nums2Map := make(map[int]struct{})

	for i := 0; i < len(nums1); i++ {
		nums1Map[nums1[i]] = struct{}{}
	}

	for i := 0; i < len(nums2); i++ {
		nums2Map[nums2[i]] = struct{}{}
	}

	diff1 := make([]int, 0)
	for i := 0; i < len(nums1); i++ {
		if _, ok := nums2Map[nums1[i]]; !ok {
			diff1 = append(diff1, nums1[i])
			nums2Map[nums1[i]] = struct{}{}
		}
	}
	diff2 := make([]int, 0)
	for i := 0; i < len(nums2); i++ {
		if _, ok := nums1Map[nums2[i]]; !ok {
			diff2 = append(diff2, nums2[i])
			nums1Map[nums2[i]] = struct{}{}
		}
	}

	return [][]int{diff1, diff2}

}

/*

approach1 : using hashMap to store num1 & num2 's element and iterating them to find difference

TO(n)
SO(n)

approach2 : Union Find to create group1 for nums1, and group2 for nums2

TO(n)
SO(n)

*/