package DP

func MostPoints(questions [][]int) int64 {
	return mostPoints(questions)
}

func mostPoints(questions [][]int) int64 {
	/*
		memo := make([]int, len(questions))
		return int64(mostPointsTopDown(questions, 0, memo))
	*/

	dp := make([]int, len(questions))
	dp[len(dp)-1] = questions[len(questions)-1][0]
	for i := len(questions) - 2; i >= 0; i-- {
		dp[i] = questions[i][0]

		skip := questions[i][1]

		if i+skip < len(questions)-1 {
			dp[i] += dp[i+skip+1]
		}

		// dp[i] = max(solve it, skip it)

		dp[i] = max(dp[i], dp[i+1])
	}

	return int64(dp[0])
}

func mostPointsTopDown(questions [][]int, start int, memo []int) int {
	if start >= len(questions) {
		return 0
	}

	if memo[start] != 0 {
		return memo[start]
	}

	takeCurrent := questions[start][0] + mostPointsTopDown(questions, start+questions[start][1]+1, memo)
	notTakeCurrent := mostPointsTopDown(questions, start+1, memo)
	memo[start] = max(takeCurrent, notTakeCurrent)
	return memo[start]
}
