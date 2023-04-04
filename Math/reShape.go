package Math

func MatrixReshape(mat [][]int, r int, c int) [][]int {
	return matrixReshape(mat, r, c)
}

func matrixReshape(mat [][]int, r int, c int) [][]int {
	res := make([][]int, 0)

	m := len(mat)
	n := len(mat[0])

	if m*n != r*c {
		return mat
	}

	current := make([]int, c)
	index := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			current[index] = mat[i][j]
			index++

			if index == c {
				index = 0
				//temp := make([]int, c)
				//copy(temp[:], current)
				res = append(res, current)
				current = make([]int, c)
			}

		}
	}

	return res
}
