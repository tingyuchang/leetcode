package DP

func StoneGameII(piles []int) int {
	return stoneGameII(piles)
}

func stoneGameII(piles []int) int {
	n := len(piles)
	dp := make([][][]int, 2)

	for i := range dp {
		dp[i] = make([][]int, n+1)
		for j := 0; j <= n; j++ {
			dp[i][j] = make([]int, n+1)
			for k := 0; k <= n; k++ {
				dp[i][j][k] = -1
			}
		}
	}

	return dpStoneGameII(piles, dp, 0, 0, 1)
}

func dpStoneGameII(piles []int, dp [][][]int, p, i, m int) int {
	if i == len(piles) {
		return 0
	}

	if dp[p][i][m] != -1 {
		return dp[p][i][m]
	}
	res := 1000000
	if p != 1 {
		res = -1
	}
	s := 0
	for j := 1; j <= min(2*m, len(piles)-i); j++ {
		s += piles[i+j-1]
		if p == 0 {
			res = max(res, s+dpStoneGameII(piles, dp, 1, i+j, max(m, j)))
		} else {
			res = min(res, dpStoneGameII(piles, dp, 0, i+j, max(m, j)))
		}
	}

	dp[p][i][m] = res
	return dp[p][i][m]
}

func StoneGame(piles []int) bool {
	return stoneGame(piles)
}

func stoneGame(piles []int) bool {
	/*
		Alice clearly always wins the 2 pile game. With some effort, we can see that she always wins the 4 pile game.
		If Alice takes the first pile initially, she can always take the third pile. If she takes the fourth pile initially, she can always take the second pile. At least one of first + third, second + fourth is larger, so she can always win.
		We can extend this idea to N piles. Say the first, third, fifth, seventh, etc. piles are white, and the second, fourth, sixth, eighth, etc. piles are black. Alice can always take either all white piles or all black piles, and one of the colors must have a sum number of stones larger than the other color.
		Hence, Alice always wins the game.
	*/
	return true
}
