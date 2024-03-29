package Math

func fib(n int) int {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	temp1 := 0
	temp2 := 1
	for i := 2; i <= n; i++ {
		temp3 := temp1 + temp2
		temp1, temp2 = temp2, temp3
	}

	return temp2
}

func climbStairs(n int) int {
	if n <= 1 {
		return 1
	}

	temp1 := 1

	temp2 := 1

	for i := 2; i <= n; i++ {
		temp3 := temp1 + temp2
		temp1, temp2 = temp2, temp3
	}

	return temp2
}

/*
1(1): 1
2(2): 1+1, 2
3(3): 1+1+1, 2+1, 1+2
4(5): 1+1+1+1, 1+1+2, 1+2+1, 2+1+1, 2+2
5(8): 1+1+1+1+1, 1+1+1+2, 1+1+2+1, 1+2+1+1, 2+1+1+1, 1+2+2, 2+1+2, 2+2+1
6(): 1+1+1+1+1+1, 1+1+1+1+2, 1+1+1+2+1, 1+1+2+1+1, 1+2+1+1+1, 2+1+1+1+1, 1+1+2+2, 1+2+2+1, 1+2+1+2, 2+1+2+1, 2+2+1+1, 2+2+2

*/
