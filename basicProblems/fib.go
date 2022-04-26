package basicProblems



func fib(n int) int {
	if n < 2 {
		return n
	}

	return  fib(n-1) + fib(n-2)
}

func fib2(n int) int {
	previous := 0
	current := 1
	sum := 0

	if n == 0 {
		return previous
	}

	if n == 1 {
		return current
	}

	for i:=1; i<n; i++ {
		sum = previous+current
		previous = current
		current = sum
	}

	return sum
}