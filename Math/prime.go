package Math

func IsUgly(n int) bool {
	if n == 1 {
		return true
	}

	if n == 0 {
		return false
	}

	for n%2 == 0 {
		n = n / 2
	}

	for n%3 == 0 {
		n = n / 3
	}

	for n%5 == 0 {
		n = n / 5
	}

	return n == 1
}
