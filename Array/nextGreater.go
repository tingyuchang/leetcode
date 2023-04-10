package Array

func NextGreaterElementsII(nums []int) []int {
	return nextGreaterElementsII(nums)
}

func nextGreaterElementsII(nums []int) []int {
	stack := make([]int, 0)
	//cache := make(map[int]int)
	res := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		res[i] = -1
		for len(stack) > 0 {
			if nums[i] <= nums[stack[len(stack)-1]] {
				break
			} else {
				res[stack[len(stack)-1]] = nums[i]
				stack = stack[:len(stack)-1]
			}
		}
		stack = append(stack, i)
	}

	for i := 0; i < len(nums); i++ {
		for len(stack) > 0 {

			if nums[i] <= nums[stack[len(stack)-1]] {
				break
			} else {
				res[stack[len(stack)-1]] = nums[i]
				stack = stack[:len(stack)-1]
			}
		}
	}

	return res
}

func NextGreaterElement(nums1 []int, nums2 []int) []int {
	return nextGreaterElement(nums1, nums2)
}

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	res := make([]int, len(nums1))
	stack := make([]int, 0)
	cache := make(map[int]int)
	for i := 0; i < len(nums2); i++ {

		for len(stack) != 0 {
			if nums2[i] < stack[len(stack)-1] {
				break
			} else {
				cache[stack[len(stack)-1]] = nums2[i]
				stack = stack[:len(stack)-1]
			}
		}

		stack = append(stack, nums2[i])
	}

	for i := 0; i < len(nums1); i++ {
		if v, ok := cache[nums1[i]]; ok {
			res[i] = v
		} else {
			res[i] = -1
		}
	}

	return res

}
