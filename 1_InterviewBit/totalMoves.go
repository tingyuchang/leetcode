package __InterviewBit

func totalMoves(A int, B int) int {
	row, col := 8, 8
	diresctins := [][2]int{
		{-1, -1},
		{-1, 1},
		{1, -1},
		{1, 1},
	}

	move := 0

	for _, direction := range diresctins {
		x, y := A, B
		for {
			x, y = x+direction[0], y+direction[1]

			if x >= row+1 || x <= 0 || y <= 0 || y >= col+1 {
				break
			}

			move++
		}
	}

	return move
}
