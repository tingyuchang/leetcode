package Matrix

func SearchMatrix(matrix [][]int, target int) bool {
	n := len(matrix[0]) - 1
	l := 0
	r := len(matrix) - 1

	slot := 0
	for r >= l {
		mid := int(uint(l+r) >> 1)
		if target > matrix[mid][n] {
			l = mid + 1
		} else if target < matrix[mid][0] {
			r = mid - 1
		} else {
			slot = mid
			break
		}
	}

	l = 0

	r = n

	for r >= l {
		mid := int(uint(l+r) >> 1)

		if matrix[slot][mid] > target {
			r = mid - 1
		} else if matrix[slot][mid] < target {
			l = mid + 1
		} else {
			return true
		}

	}

	return false
}
