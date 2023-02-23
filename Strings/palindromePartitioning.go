package Strings

import "fmt"

/*
1. backtracking + dfs
2. backtracking with dynamic programming
*/
func PalindromePartition(s string) [][]string {
	result := make([][]string, 0)
	dp := make([][]bool, len(s))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, len(s))
	}
	dfsSubstringDP(0, &result, &[]string{}, s, &dp)
	return result
}

func dfsSubstringDP(start int, result *[][]string, current *[]string, s string, dp *[][]bool) {
	if start >= len(s) {
		temp := append([]string{}, *current...)
		*result = append(*result, temp)
		return
	}

	for end := start; end < len(s); end++ {
		if (s[start] == s[end]) && (end-start <= 2 || (*dp)[start+1][end-1]) {
			(*dp)[start][end] = true
			*current = append(*current, s[start:end+1])
			dfsSubstringDP(end+1, result, current, s, dp)
			*current = (*current)[:len(*current)-1]
		}
	}
}

func dfsSubstring(start int, result *[][]string, current *[]string, s string) {
	if start >= len(s) {
		temp := append([]string{}, *current...)
		//fmt.Println("ADD: ", temp)
		*result = append(*result, temp)
		return
	}

	for end := start; end < len(s); end++ {
		fmt.Printf("(%d, %d) %v\n", start, end, s[start:end+1])
		if isPalindrome(s, start, end) {
			*current = append(*current, s[start:end+1])
			//fmt.Println(current)
			dfsSubstring(end+1, result, current, s)
			*current = (*current)[:len(*current)-1]
			fmt.Println("remove: ", current)
		}
	}
}

func isPalindrome(s string, start, end int) bool {

	for start < end {
		if s[start] != s[end] {
			return false
		}

		start++
		end--
	}
	return true
}
