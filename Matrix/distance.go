package Matrix

func UpdateMatrix(mat [][]int) [][]int {
	return updateMatrix(mat)
}

func updateMatrix(mat [][]int) [][]int {
	queue := make([][]int, 0)

	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			if mat[i][j] == 0 {
				queue = append(queue, []int{i, j})
			} else {
				mat[i][j] = -1
			}
		}
	}

	directions := [][]int{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
	}

	// bfs

	for len(queue) > 0 {
		i, j := queue[0][0], queue[0][1]
		queue = queue[1:]

		for _, direction := range directions {
			x, y := i+direction[0], j+direction[1]

			if x >= 0 && x < len(mat) && y >= 0 && y < len(mat[0]) {
				if mat[x][y] == -1 {
					mat[x][y] = mat[i][j] + 1
					queue = append(queue, []int{x, y})
				}
			}
		}
	}

	return mat
}

