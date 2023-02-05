package sqrtx

func MySqrt(x int) int {
	// WTF
	// return int(math.Sqrt(float64(x)))

	l := 0
	r := x

	for r >= l {
		mid := (l + r) / 2
		if mid*mid > x {
			r = mid - 1
		} else if mid*mid < x {
			l = mid + 1
		} else {
			return mid
		}
	}

	return l - 1
}
