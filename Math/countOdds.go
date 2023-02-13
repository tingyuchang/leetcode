package Math

func CountOdds(low int, high int) int {
	count := high - low + 1
	if low%2 != 0 {
		if count%2 != 0 {
			return count/2 + 1
		} else {
			return count / 2
		}
	} else {
		return count / 2
	}
	return 0
}
