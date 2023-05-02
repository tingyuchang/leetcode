package Math

func myPow(x float64, n int) float64 {
	if x == 0.0 {
		return 0.0
	}

	if n == 0 || x == 1.0 {
		return 1.0
	}

	if x == -1.0 {
		if n%2 == 0 {
			return 1
		} else {
			return -1
		}
	}

	ans := 1.0
	isNegative := false
	if n < 0 {
		n = -n
		isNegative = true
	}

	// can improve by n/2
	for n > 0 {
		ans = ans * x
		n--
	}

	if isNegative {
		return 1 / ans
	}

	return ans
}

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

func IsPowerOfFour(n int) bool {
	//a := 1
	//
	//for a < n {
	//	a = a * 4
	//}
	//
	//return a == n

	if n == 0 {
		return false
	}
	if n == 1 {
		return true
	}

	if n%4 != 0 {
		return false
	}

	return IsPowerOfFour(n / 4)
}
