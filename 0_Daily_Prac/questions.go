package __Daily_Prac

func _LengthOfLIS(nums []int) int {
	return _lengthOfLIS(nums)
}

func _lengthOfLIS(nums []int) int {
	/*
		dp := make([]int, len(nums))
			ans := 0
			for i := 0; i < len(nums); i++ {
				for j := 0; j < i; j++ {
					if nums[i] > nums[j] && dp[i] <= dp[j] {
						dp[i] = dp[j] + 1

						if dp[i] > ans {
							ans = dp[i]
						}
					}
				}
			}

			return ans + 1
	*/

	subs := make([]int, 0)

	for i := 0; i < len(nums); i++ {
		if len(subs) == 0 || nums[i] > subs[len(subs)-1] {
			subs = append(subs, nums[i])
		} else {
			l := 0
			r := len(subs) - 1

			replaceIndex := 0
			for r >= l {
				mid := int(uint(l+r) >> 1)
				// fmt.Println(subs, nums[i], mid)
				if mid > 0 && (subs[mid] == nums[i] || subs[mid] > nums[i] && subs[mid-1] < nums[i]) {
					replaceIndex = mid
					break
				}

				if subs[mid] > nums[i] {
					r = mid - 1
				} else {
					l = mid + 1

				}
			}

			subs[replaceIndex] = nums[i]

		}

	}

	return len(subs)
}

/*

Approach 1: DP
TO(n^2)
SO(n)
nums = [10,9,2,5,3,7,101,18]

     0 1 2 3 4 5 6 7
10
9    x
2    x x
5    x x 1
3    x x 1 x
7    x x 1 2 x
101  1 x x 2 x 3
18   1 x x 2 x 3

Approach 2: Greedy & Binary Search

holds a sub array

if nums[i] > subArray[last] -> append it

else find element which is the smallest element in elements which greater than nums[i]

*/

func findNumberOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	count := make([]int, len(nums))
	ans := 0
	longest := 0
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		count[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				if dp[i] <= dp[j] {
					dp[i] = dp[j] + 1
					count[i] = count[j]
				} else if dp[i]-1 == dp[j] {
					count[i] += count[j]
				}
			}
		}

		if longest == dp[i] {
			ans += count[i]
		}
		if dp[i] > longest {
			longest = dp[i]
			ans = count[i]
		}

	}
	return ans
}
