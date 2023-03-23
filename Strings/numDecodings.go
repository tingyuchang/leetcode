package Strings

func NumDecodings(s string) int {
	n := len(s)

	dp := make([]int, n+1)
	dp[n] = 1
	for i := n - 1; i >= 0; i-- {
		if s[i] != '0' {
			dp[i] = dp[i+1]

			// to get dp[i] we have 2 possible ways, add 1 char or add 2 char
			if (i < n-1) && ((s[i]) == '1' ||
				s[i] == '2' && s[i+1] < '7') {
				dp[i] += dp[i+2]
			}

		}
	}

	return dp[0]
}
