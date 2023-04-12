package __InterviewBit

func maxProfit(A []int) int {
	if len(A) <= 1 {
		return 0
	}

	minVal := A[0]
	maxProfit := 0

	for i := 1; i < len(A); i++ {
		if A[i] < minVal {
			minVal = A[i]
		} else if A[i]-minVal > maxProfit {
			maxProfit = A[i] - minVal
		}
	}

	return maxProfit

}

func maxProfitII(A []int) int {
	if len(A) == 0 {
		return 0
	}

	minVal := A[0]
	profit := 0

	for i := 0; i < len(A); i++ {
		if A[i] < minVal {
			minVal = A[i]
		} else if A[i] > minVal {
			profit += A[i] - minVal
			minVal = A[i]
		}
	}

	return profit

}
