package Strings

import "fmt"

func MinInsertions(s string) int {
	return minInsertions(s)
}

func minInsertions(s string) int {
	n := longestPalindromeSubseq(s)

	return len(s) - n
}

func LongestPalindromeByConcatenating(words []string) int {
	return longestPalindromeByConcatenating(words)
}

func longestPalindromeByConcatenating(words []string) int {
	// find current word's palindrome if find , the there 2 can use to create palindrome string
	// notes:
	// 1. if find more than 2, we only use 2n candidates
	// 2. if word itself is palindrome, then it can be along
	// ex: aa, we can use below cases {"aa"} or {"aa", "aa"} or {"aa", "aa", "aa"}
	// T O(2n) n is number of keys
	// S O(n) n is number of keys

	cache := make(map[string]int)
	res := 0

	for _, v := range words {
		cache[v]++
	}

	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}

	only := true
	for k, v := range cache {

		kr := string(k[1]) + string(k[0])
		//fmt.Println(res, k, kr)
		if kr != k {
			v2, ok := cache[kr]
			if ok {
				res += min(v, v2) * 4
			}
			// reset k & kr
			cache[k] = 0
			cache[kr] = 0
		} else {

			res += (v / 2) * 2 * 2
			// only accept 1
			if only && v%2 != 0 {
				res += 2
				only = false
			}
		}
	}

	return res
}

func LongestPalindromeSubString(s string) string {
	if len(s) == 0 {
		return ""
	}
	start, end := 0, 0

	max := func(a, b int) int {
		if a > b {
			return a
		}

		return b
	}

	for i := 0; i < len(s); i++ {
		len1 := expandAroundCenter(s, i, i)
		len2 := expandAroundCenter(s, i, i+1) // if center is between 2 nodes
		ans := max(len1, len2)
		if ans > end-start+1 {
			start = i - (ans-1)/2
			end = i + ans/2
		}
	}

	return s[start : end+1]
}

func expandAroundCenter(s string, l, r int) int {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}

	return r - l - 1
}

func LongestPalindromeSubseq(s string) int {
	return longestPalindromeSubSequenceButtonUp(s)
}

func longestPalindromeSubSequenceButtonUp(s string) int {
	dp := make([][]int, len(s))
	for i, _ := range dp {
		dp[i] = make([]int, len(s))
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}

		return b
	}

	for i := len(s) - 1; i >= 0; i-- {
		dp[i][i] = 1
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}

	return dp[0][len(s)-1]

}

func longestPalindromeSubseq(s string) int {
	// dp[i][j] means s[i:j] is palindrome string length
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
