package Matrix

import "fmt"

func isValidSudoku(board [][]byte) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == '.' {
				continue
			}

			if !checkSudoku(i, j, board) {
				return false
			}
		}
	}

	return true
}

func checkSudoku(x, y int, board [][]byte) bool {

	for i := 0; i < 9; i++ {
		if i == y {
			continue
		}

		if board[x][i] == board[x][y] {
			return false
		}
	}

	for i := 0; i < 9; i++ {
		if i == x {
			continue
		}

		if board[i][y] == board[x][y] {
			return false
		}
	}

	positionRow := (x / 3) * 3
	positionCol := (y / 3) * 3

	for i := 0; i < 9; i++ {
		r, c := i/3, i%3

		if r == x%3 && c == y%3 {
			continue
		}

		if board[positionRow+r][positionCol+c] == board[x][y] {
			fmt.Println(positionRow+r, c, x, y)
			return false
		}

	}

	return true
}
