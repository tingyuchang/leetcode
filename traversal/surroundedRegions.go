package traversal

func SolveSurroundedRegions(board [][]byte) {
	for i := 0; i < len(board); i++ {
		dfsSurroundedRegions(i, 0, board)
		dfsSurroundedRegions(i, len(board[i])-1, board)

	}
	for j := 0; j < len(board[0]); j++ {
		dfsSurroundedRegions(0, j, board)
		dfsSurroundedRegions(len(board)-1, j, board)
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			} else if board[i][j] == '-' {
				board[i][j] = 'O'
			}
		}
	}

}

func dfsSurroundedRegions(x, y int, board [][]byte) {
	if x < 0 || x >= len(board) || y < 0 || y >= len(board[0]) || board[x][y] != 'O' {
		return
	}
	board[x][y] = '-'

	dfsSurroundedRegions(x, y+1, board)
	dfsSurroundedRegions(x, y-1, board)
	dfsSurroundedRegions(x+1, y, board)
	dfsSurroundedRegions(x-1, y, board)
}
