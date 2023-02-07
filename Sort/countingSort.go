package Sort

func CountingSort(a []int, n, k int) []int {
	b := make([]int, n)
	c := make([]int, k+1)

	for i := 0; i < n; i++ {
		c[a[i]] = c[a[i]] + 1
	}
	for i := 1; i < len(c); i++ {
		c[i] = c[i] + c[i-1]
		// c[i] now contains the number of elements less than or equal to i
	}
	// copy a to b, starting from the end of a
	// {3, 2, 5, 5, 1, 8}
	for i := len(a) - 1; i >= 0; i-- {
		b[c[a[i]]-1] = a[i]
		c[a[i]] = c[a[i]] - 1 // ro handle the duplicate value
	}

	return b
}
