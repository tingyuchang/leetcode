package Math

func CountBits(n int) []int {
	result := make([]int, n+1)
	result[0] = 0

	for i := 1; i <= n; i++ {
		if i%2 == 0 {
			result[i] = result[i/2]
		} else {
			result[i] = result[i/2] + 1
		}

	}
	return result
}
