package DP

func New21Game(n int, k int, maxPts int) float64 {
	return new21Game(n, k, maxPts)
}

func new21Game(n int, k int, maxPts int) float64 {
	if k == 0 || k+maxPts <= n {
		return 1
	}
	/*
		Input: n = 21, k = 17, maxPts = 10
		bottom up
		base case:
		form k-1 to calculate base percentage
		ex: k = 17, then at point 16, there are [1,2,3,4,5] can pass
		so pass percentage = [1,2,3,4,5] / maxPts

		in point 15. there 3 cases
		1. get 1 pt => go 16
		2. get [2,3,4,5,6] = pass
		3. get [7,8,9,10] => fail

		therefore pt 15's pass percentage = 1 + 2 :
		1/maxPts * pass_percentage(16) + [2,3,4,5,6]/maxPts
		we can make it shorter: pass_percentage(15) = pass_percentage(16) * (1 + 1/maxPts)
		pass_percentage(14) = pass_percentage(15) * (1 + 1/maxPts)
		and so on...
		but in point 10, only has [7,8,9,10] can pass
		so pass_percentage(10) = pass_percentage(11) * (1 + 1/maxPts) - 1/maxPts

		final in point 6, things will change again, even we get maxPts, we still can't pass
		now we need to combine each case
		pass_percentage(6) = 1/maxPts * (pass_percentage(7) + ... pass_percentage(16))
							or =  pass_percentage(7) * (1 + 1/maxPts) - pass_percentage(16) -> sliding window

		in pass_percentage(0) we will get the answer.
	*/
	start := k - 1
	maxPtsFloat := float64(maxPts)
	basePercentage := 1.0 / maxPtsFloat
	dp := make([]float64, start+1)
	ans := float64(n-k+1) / maxPtsFloat
	dp[start] = ans
	j := start
	//fmt.Println("base: ", basePercentage, "ans: ", ans)
	for i := start - 1; i >= 0; i-- {
		if i >= n-maxPts {
			ans = ans + ans*basePercentage
		} else if i >= k-maxPts-1 {
			ans = ans + ans*basePercentage - basePercentage
		} else {
			ans = ans + ans*basePercentage - dp[j]*basePercentage
			j--
		}
		dp[i] = ans

		//fmt.Printf("index: %d ans: %.5f\n", i, ans)
	}
	return float64(int(ans*100000)) / 100000.0
}
