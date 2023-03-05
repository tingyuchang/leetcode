package Math

func getNoZeroIntegers(n int) []int {
	check := func(a int) bool {
		for a > 0 {
			if a%10 == 0 {
				return false
			}
			a = a / 10
		}
		return true
	}

	for i := 1; i < n; i++ {
		if check(i) && check(n-i) {
			return []int{i, n - i}
		}
	}
	return nil
}
