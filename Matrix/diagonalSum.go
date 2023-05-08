package Matrix

func diagonalSum(mat [][]int) int {
	sum := 0
	n := len(mat) - 1
	for i := 0; i < len(mat); i++ {
		j := n - i
		sum += mat[i][i] + mat[i][j]
	}

	if n%2 == 0 {
		sum -= mat[n/2][n/2]
	}

	return sum

}
