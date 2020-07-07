package fourSum

import (
	"fmt"
	"sort"
)

func FourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	return kSum(nums, target, 4)
}

func kSum(nums []int, target, k int) [][]int {
	var results [][]int

	if len(nums) == 0 || nums[0] *k > target || target > nums[len(nums)-1]*k {
		return results
	}

	if k == 2 {
		return twoSum(nums, target)
	}

	for i,v := range nums {
		fmt.Println(v)
		if i == 0 || nums[i] != nums[i-1] {
			for _,v2 := range kSum(nums[i+1:], target-nums[i], k-1) {
				v2 = append(v2, v)
				results = append(results,v2)
			}
		}
	}

	return results
}


func twoSum(nums []int, target int) [][]int {
	var results [][]int
	l := 0
	r := len(nums)-1

	for l < r {
		sum := nums[l] + nums[r]
		if sum > target || (r < len(nums)-1 && nums[r] == nums[r+1]){
			r--
		} else if sum < target || (l > 0 && nums[l] == nums[l-1]){
			l++
		} else {
			results = append(results, []int{nums[l], nums[r]})
			l++
			r--
		}
	}

	return results
}
