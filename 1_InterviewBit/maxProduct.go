package __InterviewBit

func maxProduct(A []int) int {
	minVal, maxVal, res := A[0], A[0], A[0]

	max := func(a, b int) int {
		if a > b {
			return a
		}

		return b
	}

	min := func(a, b int) int {
		if a < b {
			return a
		}

		return b
	}

	for i := 1; i < len(A); i++ {
		v := A[i]

		if v < 0 {
			minVal, maxVal = maxVal, minVal
		}

		minVal = min(v, minVal*v)
		maxVal = max(v, maxVal*v)

		if maxVal > res {
			res = maxVal
		}

	}

	return res

}
