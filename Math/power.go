package Math

func IsPowerOfTwo(n int) bool {
	if n <= 0 {
		return false
	}

	for n > 1 {
		if n%2 != 0 {
			return false
		}
		n = n >> 1
	}

	return true
}

func isPowerOfThree(n int) bool {
	if n <= 0 {
		return false
	}

	for n > 1 {
		if n%3 != 0 {
			return false
		}

		n = n / 3
	}

	/*
		a :=1
		for a < 3 { a*=3 }
		return a == n
	*/

	return true
}
