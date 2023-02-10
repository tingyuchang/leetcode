package Math

func IsHappy(n int) bool {
	m := make(map[int]bool)
	m[n] = true
	for n != 1 {
		sum := 0
		for n != 0 {
			sum += (n % 10) * (n % 10)
			n = n / 10
		}

		n = sum
		_, ok := m[n]

		if ok {
			return false
		} else {
			m[n] = true
		}
		if n == 1 {
			return true
		}
	}
	return true
}
