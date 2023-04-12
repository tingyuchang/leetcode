package __InterviewBit

func climbStairs(A int) int {

	preStep1 := 1

	preStep2 := 1

	for i := 2; i <= A; i++ {
		temp := preStep1
		preStep1 = preStep2
		preStep2 = temp + preStep2
	}

	return preStep2

}
