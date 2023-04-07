package Matrix

func RotateImage(matrix [][]int) {
	rotateImage(matrix)
}

func rotateImage(matrix [][]int) {
	rotateMatrix(matrix)
}

func FindRotation(mat [][]int, target [][]int) bool {
	return findRotation(mat, target)
}

func findRotation(mat [][]int, target [][]int) bool {
	// 0
	if checkMatrixIsSame(mat, target) {
		return true
	}
	mat = rotateMatrix(mat)
	// 90
	if checkMatrixIsSame(mat, target) {
		return true
	}

	// 180
	mat = rotateMatrix(mat)
	if checkMatrixIsSame(mat, target) {
		return true
	}

	// 270
	mat = rotateMatrix(mat)
	if checkMatrixIsSame(mat, target) {
		return true
	}

	return false
}

// 90-degree clockwise
func rotateMatrix(mat [][]int) [][]int {
	l := 0
	r := len(mat) - 1

	for r > l {
		mat[l], mat[r] = mat[r], mat[l]
		l++
		r--
	}

	for i := 0; i < len(mat); i++ {
		for j := i + 1; j < len(mat); j++ {
			mat[i][j], mat[j][i] = mat[j][i], mat[i][j]
		}
	}

	return mat
}

func checkMatrixIsSame(a, b [][]int) bool {
	ma, na := len(a), len(a[0])
	mb, nb := len(b), len(b[0])

	if ma != mb || na != nb {
		return false
	}

	for i := 0; i < ma; i++ {
		for j := 0; j < na; j++ {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}

	return true

}
