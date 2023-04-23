package Array

import "strconv"

var MOD int = 1e9 + 7

func NumberOfArrays(s string, k int) int {
	return numberOfArraysButtonUp(s, k)
}

func numberOfArrays(s string, k int) int {
	dp := make([]int, len(s)+1)
	dp[len(s)] = 0

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != '0' {
			if dp[i+1] == 0 {
				dp[i] = 1
			} else {
				dp[i] = dp[i+1] * 2
			}

		} else {
			dp[i] = dp[i+1]
		}
	}

	return dp[0]
}

func numberOfArraysTopDown(s string, k int) int {
	dp := make([]int, len(s)+1)
	return dfsNumbersOfArrayTopDown(0, k, s, dp)
}

func dfsNumbersOfArrayTopDown(start, k int, s string, dp []int) int {
	if dp[start] != 0 {
		return dp[start]
	}

	if start == len(s) {
		return 1
	}

	if s[start] == '0' {
		return 0
	}

	count := 0

	for i := start; i < len(s); i++ {
		x, _ := strconv.Atoi(string(s[start : i+1]))

		if x > k {
			break
		}
		count = (count + dfsNumbersOfArrayTopDown(i+1, k, s, dp)) % MOD
	}

	dp[start] = count

	return dp[start]
}

func numberOfArraysButtonUp(s string, k int) int {
	dp := make([]int, len(s)+1)
	dp[0] = 1
	for i := 0; i < len(s); i++ {
		if s[i] == '0' {
			continue
		}

		for j := i; j < len(s); j++ {
			x, _ := strconv.Atoi(string(s[i : j+1]))

			if x > k {
				break
			}

			dp[j+1] = (dp[j+1] + dp[i]) % MOD
		}
	}

	return dp[len(s)]
}
