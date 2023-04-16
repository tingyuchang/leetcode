package DP

func NumWays(words []string, target string) int {
	return numWays(words, target)
}
func numWays(words []string, target string) int {
	alphabet := 26
	mod := 1000000007
	m := len(target)
	k := len(words[0])
	cnt := make([][]int, alphabet)
	for i := range cnt {
		cnt[i] = make([]int, k)
	}
	for j := 0; j < k; j++ {
		for _, word := range words {
			cnt[word[j]-'a'][j]++
		}
	}
	dp := make([][]int64, m+1)
	for i := range dp {
		dp[i] = make([]int64, k+1)
	}
	dp[0][0] = 1
	for i := 0; i <= m; i++ {
		for j := 0; j < k; j++ {
			if i < m {
				dp[i+1][j+1] += int64(cnt[target[i]-'a'][j]) * dp[i][j] % int64(mod)
				dp[i+1][j+1] %= int64(mod)
			}
			dp[i][j+1] += dp[i][j]
			dp[i][j+1] %= int64(mod)
		}
	}
	return int(dp[m][k])
}
