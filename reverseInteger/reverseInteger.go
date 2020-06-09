package reverseInteger

func Reverse(x int) int {
	var isNegative bool
	if x < 0 {
		isNegative = true
		x = x * -1
	}

	n := []int{}
	count, result := 0, 0
	for float64(x)/10.0 > 0 {
		n = append(n, x%10)
		x = x/10
		count++
	}

	for i := 0; i < count; i++ {
		result = result*10 + n[i]
	}

	if result > 2147483647 {
		return 0
	}

	if isNegative {
		return result * -1
	}
	return result
}