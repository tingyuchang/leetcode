package Strings

/*
base case: when 1 word is empty, the minimum number is the len of another word

0 1 2 3 4
1 x x x x
2 x x x x
3 x x x x
4 x x x x
5 x x x x

for general case

compare current charaters, if match, then reference dp[i-1][j-1]

if not match find minimum val from dp[i-1][j-1], dp[i][j-1], dp[i-1][j] and plus 1

general case:
*/
func MinDistance(word1 string, word2 string) int {
	dp := make([][]int, len(word1)+1)

	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(word2)+1)
		dp[i][0] = i

		if i == 0 {
			for j, _ := range dp[0] {
				dp[0][j] = j
			}
		}
	}

	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[i]); j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = 1 + minInThree(dp[i][j-1], dp[i-1][j], dp[i-1][j-1])
			}
		}
	}

	return dp[len(word1)][len(word2)]
}

func minInThree(a, b, c int) int {
	var min int
	if a > b {
		min = b
	} else {
		min = a
	}

	if min > c {
		return c
	}

	return min
}
