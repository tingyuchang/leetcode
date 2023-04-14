package Strings

import "fmt"

func LongestPalindromeSubseq(s string) int {
	return longestPalindromeSubseq(s)
}

func longestPalindromeSubseq(s string) int {
	dp := make([][]int, len(s))
	fmt.Println("s len: ", len(s))
	for i, _ := range dp {
		dp[i] = make([]int, len(s))
	}
	return dpLongestPalindromeSubseq(dp, s, 0, len(s)-1)
}

func dpLongestPalindromeSubseq(dp [][]int, s string, x, y int) int {
	if x > y {
		return 0
	}
	if x == y {
		return 1
	}

	if dp[x][y] != 0 {
		return dp[x][y]
	}

	max := func(a, b int) int {
		if a > b {
			return a
		}

		return b
	}

	if s[x] == s[y] {
		dp[x][y] = dpLongestPalindromeSubseq(dp, s, x+1, y-1) + 2
	} else {
		dp[x][y] = max(dpLongestPalindromeSubseq(dp, s, x+1, y), dpLongestPalindromeSubseq(dp, s, x, y-1))
	}

	return dp[x][y]
}

func combinationOfSubSeq(s string, currentIndex int, res *int, current []byte) {
	if currentIndex > len(s) {
		return
	}

	isPalind := func(a []byte) bool {
		l := 0
		r := len(a) - 1

		for r > l {
			if a[l] != a[r] {
				return false
			}
			l++
			r--
		}

		return true
	}

	if len(current) > 0 {

		if isPalind(current) && len(current) > *res {
			*res = len(current)
		}

	}

	for i := currentIndex; i < len(s); i++ {
		current = append(current, s[i])
		combinationOfSubSeq(s, i+1, res, current)
		current = current[:len(current)-1]
	}
}

func LongestPalindrome(s string) int {
	return longestPalindrome(s)
}

func longestPalindrome(s string) int {
	res := 0
	cache := make(map[byte]int)
	for i, _ := range s {
		cache[s[i]]++

		if cache[s[i]] == 2 {
			res += 2
			cache[s[i]] = 0
		}
	}

	if res == len(s) {
		return res
	}

	res++

	return res
}

func IsPalindrome(x int) bool {

	if x < 0 {
		return false
	}
	y := reverse(x)

	if x == y {
		return true
	}

	return false
}

func reverse(x int) int {
	var isNegative bool
	if x < 0 {
		isNegative = true
		x = x * -1
	}

	n := []int{}
	count, result := 0, 0
	for float64(x)/10.0 > 0 {
		n = append(n, x%10)
		x = x / 10
		count++
	}

	for i := 0; i < count; i++ {
		result = result*10 + n[i]
	}

	if result > 2147483647 {
		return 0
	}

	if isNegative {
		return result * -1
	}
	return result
}
