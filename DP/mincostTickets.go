package DP

import (
	"math"
)

func MincostTickets(days []int, costs []int) int {
	dp := make([]int, len(days)+1)
	week := 1
	month := 1

	for i := 0; i < len(days); i++ {

		// if we decide buy dayPass
		dayPass := costs[0] + dp[i]

		// find week window, if we decide buy week pass
		for days[i]-days[week-1]+1 > 7 {
			week++
		}

		/// find month window
		weekPass := costs[1] + dp[week-1]

		for days[i]-days[month-1]+1 > 30 {
			month++
		}

		monthPass := costs[2] + dp[month-1]
		//fmt.Printf("DP: %v\t day: %d week(%d): %d, month(%d):%d \n", dp, dayPass, week, weekPass, month, monthPass)

		dp[i+1] = int(math.Min(float64(dayPass), math.Min(float64(monthPass), float64(weekPass))))
	}

	return dp[len(days)]
}
