package DP

func MaxUncrossedLines(nums1 []int, nums2 []int) int {
	return maxUncrossedLines(nums1, nums2)
}

func maxUncrossedLines(nums1 []int, nums2 []int) int {
	/*
		memo := make([][]int, len(nums1))
		for i := range memo {
			memo[i] = make([]int, len(nums2))

			for j := range memo[i] {
				memo[i][j] = -1
			}
		}

		return maxUncrossedLinesTopDown(len(nums1)-1, len(nums2)-1, nums1, nums2, memo)

	*/

	dp := make([]int, len(nums2)+1)
	dpPre := make([]int, len(nums2)+1)

	for i := 1; i <= len(nums1); i++ {
		for j := 1; j <= len(nums2); j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[j] = dpPre[j-1] + 1
			} else {
				dp[j] = max(dpPre[j], dp[j-1])
			}
		}

		copy(dpPre[:], dp[:])
	}

	return dp[len(nums2)]
}

func maxUncrossedLinesTopDown(x, y int, nums1, nums2 []int, memo [][]int) int {
	if x < 0 || y < 0 {
		return 0
	}

	if memo[x][y] != -1 {
		return memo[x][y]
	}

	max := func(a, b int) int {
		if a > b {
			return a
		}

		return b
	}
	count := 0
	if nums1[x] == nums2[y] {
		count = maxUncrossedLinesTopDown(x-1, y-1, nums1, nums2, memo) + 1
	} else {
		count = max(maxUncrossedLinesTopDown(x-1, y, nums1, nums2, memo), maxUncrossedLinesTopDown(x, y-1, nums1, nums2, memo))
	}

	memo[x][y] = count

	return memo[x][y]
}

/*

Think dynamic programming. Given an oracle dp(i,j) that tells us how many lines A[i:], B[j:]

[the sequence A[i], A[i+1], ... and B[j], B[j+1], ...] are uncrossed, can we write this as a recursion?



*/
