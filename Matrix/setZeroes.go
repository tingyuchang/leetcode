package Matrix

func SetZeroes(matrix [][]int) {
	queueRow := make(map[int]struct{})
	queueCol := make(map[int]struct{})
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				queueRow[i] = struct{}{}
				queueCol[j] = struct{}{}
			}
		}
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			_, ok := queueRow[i]
			_, ok2 := queueCol[j]
			if ok || ok2 {
				matrix[i][j] = 0
			}
		}
	}

}
