package DivideAndConquer

func Multiply(a, b [][]int) [][]int {
	// length check
	if len(a[0]) != len(b) {
		return nil
	}

	n := len(a)
	m := len(b[0])
	l := len(a[0])

	c := make([][]int, n)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			for k := 0; k < l; k++ {
				if len(c[i]) == 0 {
					c[i] = make([]int, m)
				}
				c[i][j] = c[i][j] + a[i][k]*b[k][j]
			}
		}
	}
	return c
}
