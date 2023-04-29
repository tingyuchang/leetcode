package _0230429

import (
	__Daily_Prac "leetcode/0_Daily_Prac"
	"math"
	"sort"
	"strconv"
)

type Name struct {
}

// 1143. Longest Common Subsequence
/*
Input: text1 = "abcde", text2 = "ace"
Output: 3
Explanation: The longest common subsequence is "ace" and its length is 3.
*/
func LongestCommonSubsequence(text1 string, text2 string) int {
	dp := make([][]int, len(text1)+1)

	for i := range dp {
		dp[i] = make([]int, len(text2)+1)
	}

	for i := 1; i <= len(text1); i++ {
		for j := 1; j <= len(text2); j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = __Daily_Prac.Max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[len(text1)][len(text2)]
}

// 5. Longest Palindromic Substring
/*
Input: s = "babad"
Output: "bab"
Explanation: "aba" is also a valid answer.
*/
func LongestPalindrome(s string) string {

	if len(s) == 0 {
		return ""
	}
	start := 0
	end := 0
	for i := 0; i < len(s); i++ {
		ans1 := expand(s, i, i)
		ans2 := expand(s, i, i+1)

		ans := __Daily_Prac.Max(ans1, ans2)

		if ans > end-start+1 {
			start = i - (ans-1)/2
			end = i + ans/2
		}
	}

	return s[start : end+1]
}

func expand(s string, start, end int) int {
	for start >= 0 && end < len(s) && s[start] == s[end] {
		start--
		end++
	}

	return end - start + 1 - 2
}

// 516. Longest Palindromic Subsequence
/*
Input: s = "bbbab"
Output: 4
Explanation: One possible longest palindromic subsequence is "bbbb".
*/
func LongestPalindromeSubseq(s string) int {
	dp := make([][]int, len(s))

	for i := len(s) - 1; i >= 0; i-- {
		dp[i] = make([]int, len(s))
		dp[i][i] = 1
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = __Daily_Prac.Max(dp[i][j-1], dp[i+1][j])
			}
		}
	}

	return dp[0][len(s)-1]
}

// 322. Coin Change
/*
Input: coins = [1,2,5], amount = 11
Output: 3
Explanation: 11 = 5 + 5 + 1
*/
func CoinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	sort.Ints(coins)

	for i := 1; i <= amount; i++ {
		dp[i] = math.MaxInt
		for _, coin := range coins {
			if coin > i {
				break
			}

			if dp[i-coin] != math.MaxInt {
				dp[i] = __Daily_Prac.Min(dp[i], dp[i-coin]+1)
			}
		}
	}

	if dp[amount] != math.MaxInt {
		return dp[amount]
	}
	return -1
}

// 1416. Restore The Array
/*
Input: s = "1317", k = 2000
Output: 8
Explanation: Possible arrays are [1317],[131,7],[13,17],[1,317],[13,1,7],[1,31,7],[1,3,17],[1,3,1,7]
*/
func NumberOfArrays(s string, k int) int {
	// top down
	/*
		if memo[start] return memo[start]
		if start == len(s) return 1
		if s[start] == 0 return 0
		dfd(0) = dfs(1) + dfs(2) + ....
	*/

	/*
		dp:

		for i := 0

		if s[0:i] is validate

		dp[i+1] += dp[0]

	*/

	dp := make([]int, len(s)+1)
	dp[0] = 1

	for i := 0; i < len(s); i++ {
		if s[i] == '0' {
			continue
		}
		for j := i; j < len(s); j++ {
			x, _ := strconv.Atoi(s[i : j+1])
			if x > k {
				break
			}
			dp[j+1] = (dp[j+1] + dp[i]) % __Daily_Prac.MOD
		}
	}
	return dp[len(s)]
}

// 72. Edit Distance
/*
Input: word1 = "horse", word2 = "ros"
Output: 3
Explanation:
horse -> rorse (replace 'h' with 'r')
rorse -> rose (remove 'r')
rose -> ros (remove 'e')
*/
func MinDistance(word1 string, word2 string) int {
	dp := make([][]int, len(word1)+1)

	for i := range dp {
		dp[i] = make([]int, len(word2)+1)

		if i == 0 {
			for j := 0; j < len(dp[i]); j++ {
				dp[i][j] = j
			}
		} else {
			dp[i][0] = i
		}
	}

	for i := 1; i <= len(word1); i++ {
		for j := 1; j <= len(word2); j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = __Daily_Prac.Min(__Daily_Prac.Min(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1]) + 1
			}
		}
	}

	return dp[len(word1)][len(word2)]
}

// 139. Word Break
/*
Input: s = "applepenapple", wordDict = ["apple","pen"]
Output: true
Explanation: Return true because "applepenapple" can be segmented as "apple pen apple".
Note that you are allowed to reuse a dictionary word.
*/

func WordBreak(s string, wordDict []string) bool {
	wordsMap := make(map[string]struct{})

	for _, v := range wordDict {
		wordsMap[v] = struct{}{}
	}

	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if _, ok := wordsMap[s[j:i]]; ok && dp[j] {
				dp[i] = true
				break
			}
		}
	}

	return dp[len(s)]
}

// 127. Word Ladder
/*
Input: beginWord = "hit", endWord = "cog", wordList = ["hot","dot","dog","lot","log","cog"]
Output: 5
Explanation: One shortest transformation sequence is "hit" -> "hot" -> "dot" -> "dog" -> cog", which is 5 words long.
*/
func LadderLength(beginWord string, endWord string, wordList []string) int {
	queue := []string{beginWord}
	backQueue := make([]string, 0)
	vivisted := make(map[string]bool)

	for i := range wordList {
		vivisted[wordList[i]] = false
	}

	ans := 1
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		vivisted[current] = true

		for word, visit := range vivisted {
			if !visit {
				count := 0

				for i := 0; i < len(word); i++ {
					if current[i] != word[i] {
						count++
					}

					if count > 1 {
						break
					}
				}

				if count == 1 {
					if word == endWord {
						return ans + 1
					}

					backQueue = append(backQueue, word)
				}
			}
		}

		if len(queue) == 0 {
			ans++
			queue, backQueue = backQueue, make([]string, 0)
		}

	}

	return 0
}

// 221. Maximal Square
/*
Input: matrix = [["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]
Output: 4
*/
func MaximalSquare(matrix [][]byte) int {
	m, n := len(matrix), len(matrix[0])
	dp := make([][]int, m+1)

	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	ans := 0

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if matrix[i-1][j-1] == '1' {
				dp[i][j] = __Daily_Prac.Min(__Daily_Prac.Min(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1]) + 1

				if dp[i][j] > ans {
					ans = dp[i][j]
				}
			}
		}
	}

	return ans * ans
}
