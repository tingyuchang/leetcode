package __InterviewBit

func majorityElement(A []int) int {

	cache := make(map[int]int)

	n := len(A) / 2
	if len(A)%2 != 0 {
		n++
	}

	for i := 0; i < len(A); i++ {
		cache[A[i]]++

		if cache[A[i]] == n {
			return A[i]
		}
	}

	return 0
}
