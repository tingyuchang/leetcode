package twoSum

/*
Example 1:

Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Output: Because nums[0] + nums[1] == 9, we return [0, 1].
Example 2:

Input: nums = [3,2,4], target = 6
Output: [1,2]
Example 3:

Input: nums = [3,3], target = 6
Output: [0,1]
 */

func TwoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for i, v := range nums {
		_, n := m[v]

		if n {
			return []int{m[v], i}
		} else {
			m[target - v] = i
		}
	}

 	return nil
}