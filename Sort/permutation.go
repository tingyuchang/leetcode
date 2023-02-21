package Sort

import "sort"

func permuteUnique(nums []int) [][]int {
	result := make([][]int, 0)
	if len(nums) == 0 {
		return result
	}
	sort.Ints(nums)
	result = append(result, append([]int{}, nums...))
	if len(nums) == 1 {
		return result
	}

	// n!/(a!*b!*c!), a, b, c means numbers of duplicated elements
	n := 1
	count := make(map[int]int, 0)
	for i := 1; i <= len(nums); i++ {
		n = n * i
		val, ok := count[nums[i-1]]
		if !ok {
			count[nums[i-1]] = 1
		} else {
			count[nums[i-1]] = val * (val + 1)
		}
	}

	for _, v := range count {
		if v != 1 {
			n = n / v
		}
	}

	for n > 0 {
		n--
		i := len(nums) - 1

		for nums[i-1] >= nums[i] {
			i--
			if i == 0 {
				return result
			}
		}

		j := len(nums) - 1
		for j > i && nums[j] <= nums[i-1] {
			j--
		}
		nums[j], nums[i-1] = nums[i-1], nums[j]
		reverse(nums, i)
		temp := append([]int{}, nums...)
		result = append(result, temp)
	}

	return result
}

func Permutation(nums []int) [][]int {
	result := make([][]int, 0)
	if len(nums) == 0 {
		return result
	}
	sort.Ints(nums)
	result = append(result, append([]int{}, nums...))
	if len(nums) == 1 {
		return result
	}
	n := 1
	for i := 1; i <= len(nums); i++ {
		n = n * i
	}

	for n > 0 {
		n--
		if n == 0 {
			return result
		}
		i := len(nums) - 1

		for nums[i-1] >= nums[i] {
			i--
		}

		j := len(nums) - 1
		for j > i && nums[j] <= nums[i-1] {
			j--
		}
		nums[j], nums[i-1] = nums[i-1], nums[j]
		reverse(nums, i)
		temp := append([]int{}, nums...)
		result = append(result, temp)
	}

	return result
}

func NextPermutation(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}

	if len(nums) == 1 {
		return nums
	}

	i := len(nums) - 1
	// find the largest index i such that nums[i-1] is less than nums[i]
	// if nums[i-1] smaller than nums[i], means we can find next permutation
	// If i is not the first index of the slice, then sub-slice nums[i…n) is stored in descending order
	// i.e. nums[i-1] < nums[i] >= nums[i+1] >= nums[i+2]
	for nums[i-1] >= nums[i] {
		i = i - 1
		if i == 0 {
			reverse(nums, 0)
			return nums
		}
	}

	j := len(nums) - 1
	// find highest index j such nums[j] is greater than nums[i-1]
	// i.e. 1,2,3,4,6,5 => 1,2,3,5,6,4
	for j > i && nums[j] <= nums[i-1] {
		j = j - 1
	}
	// swap nums[i-1] & nums[j]
	nums[i-1], nums[j] = nums[j], nums[i-1]
	// reverse nums[i:]
	reverse(nums, i)
	return nums
}

func reverse(nums []int, n int) {
	i, j := n, len(nums)-1
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}
