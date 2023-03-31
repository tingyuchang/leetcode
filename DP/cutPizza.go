package DP

func CurPizza(pizza []string, k int) int {
	return ways(pizza, k)
}

func ways(pizza []string, k int) int {
	s := make([][]int, 51)
	for i := 0; i < len(s); i++ {
		s[i] = make([]int, 51)
	}
	count := func(r1, c1, r2, c2 int) int {
		var sum int
		for i := r1; i <= r2; i++ {
			sum += s[i][c2+1] - s[i][c1]
		}
		return sum
	}
	f := make([][][]int, 51)
	for i := 0; i < len(f); i++ {
		f[i] = make([][]int, 51)
		for j := 0; j < len(f[i]); j++ {
			f[i][j] = make([]int, 11)
		}
	}

	mod := int(1e9 + 7)
	rows, cols := len(pizza), len(pizza[0])
	for i := 0; i < rows; i++ {
		s[i][0] = 0
		for j := 0; j < cols; j++ {
			var tmp int
			if pizza[i][j] == 'A' {
				tmp = 1
			}
			s[i][j+1] = s[i][j] + tmp
		}
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if count(i, j, rows-1, cols-1) > 0 {
				f[i][j][0] = 1
			}
		}
	}

	for p := 1; p < k; p++ {
		for i := 0; i < rows; i++ {
			for j := 0; j < cols; j++ {
				f[i][j][p] = 0
				for q := j; q < cols; q++ {
					if count(i, j, rows-1, q) > 0 {
						f[i][j][p] = (f[i][j][p] + f[i][q+1][p-1]) % mod
					}
				}
				for q := i; q < rows; q++ {
					if count(i, j, q, cols-1) > 0 {
						f[i][j][p] = (f[i][j][p] + f[q+1][j][p-1]) % mod
					}
				}
			}
		}
	}
	return f[0][0][k-1]

}
