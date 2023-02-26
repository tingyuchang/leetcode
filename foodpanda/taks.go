package foodpanda

// "BALLOON"
func Solution(S string) int {
	// Implement your solution here
	count := make(map[int32]int)

	for _, v := range S {
		val, ok := count[v]
		if !ok {
			count[v] = 1
		} else {
			count[v] = val + 1
		}
	}

	sum := 0
	b := count['B']
	a := count['A']
	l := count['L']
	o := count['O']
	n := count['N']
	for {
		if b < 1 || a < 1 || l < 2 || o < 2 || n < 1 {
			break
		}
		b--
		a--
		l -= 2
		o -= 2
		n--
		sum++
		if b <= 0 || a <= 0 || l <= 1 || o <= 1 || n <= 0 {
			break
		}
	}

	return sum

}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
