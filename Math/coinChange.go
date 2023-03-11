package Math

import (
	"math"
	"sort"
)

func CoinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	dp[0] = 0

	sort.Ints(coins)

	for i := 1; i < len(dp); i++ {
		dp[i] = math.MaxInt
		for _, c := range coins {
			if c > i {
				break
			}

			if dp[i-c] != math.MaxInt {
				dp[i] = min(dp[i], dp[i-c]+1)
			}
		}
	}

	if dp[len(dp)-1] == math.MaxInt {
		return -1
	}
	return dp[len(dp)-1]
}
