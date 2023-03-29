package Strings

func exist(board [][]byte, word string) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if backtracking(i, j, 0, word, board) {
				return true
			}
		}
	}

	return false
}

func backtracking(x, y, currentIndex int, word string, board [][]byte) bool {
	if currentIndex == len(word) {
		return true
	}

	if x >= len(board) || x < 0 ||
		y >= len(board[0]) || y < 0 ||
		word[currentIndex] != board[x][y] {
		return false
	}

	temp := board[x][y]
	board[x][y] = '#'

	res := backtracking(x+1, y, currentIndex+1, word, board) ||
		backtracking(x, y+1, currentIndex+1, word, board) ||
		backtracking(x-1, y, currentIndex+1, word, board) ||
		backtracking(x, y-1, currentIndex+1, word, board)

	board[x][y] = temp

	return res
}
