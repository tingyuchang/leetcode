package Array

import (
	"math"
)

func NextGreaterElementIII(n int) int {
	return nextGreaterElementIII(n)
}

func nextGreaterElementIII(n int) int {
	nums := make([]int, 0)

	for n > 0 {
		nums = append([]int{n % 10}, nums...)
		n = n / 10
	}

	if len(nums) <= 1 {
		return -1
	}

	i := len(nums) - 1

	for nums[i-1] >= nums[i] {
		i--
		if i == 0 {
			return -1
		}
	}

	j := len(nums) - 1

	for j > i && nums[j] <= nums[i-1] {
		j--
	}

	nums[j], nums[i-1] = nums[i-1], nums[j]

	l := i
	r := len(nums) - 1

	for r > l {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}

	res := 0
	decimal := 0.0
	for i := len(nums) - 1; i >= 0; i-- {
		res += nums[i] * int(math.Pow(10.0, decimal))
		decimal++
	}

	if res > int(math.Pow(2.0, 31.0))-1 {
		return -1
	}

	return res
}

/*
	next := make([]int, len(nums))
	stack := make([]int, 0)

	for i := len(nums) - 1; i >= 0; i-- {
		next[i] = -1
		for len(stack) > 0 {

			if nums[i] <= nums[stack[len(stack)-1]] {
				break
			} else {
				next[stack[len(stack)-1]] = i
				stack = stack[:len(stack)-1]
			}
		}

		stack = append(stack, i)
	}

	fmt.Println(next)

	ok := false
	for i := len(next) - 1; i >= 0; i-- {
		if next[i] != -1 {
			ok = true
			nums[i], nums[next[i]] = nums[next[i]], nums[i]
			break
		}
	}

	if !ok {
		return -1
	}

*/

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
