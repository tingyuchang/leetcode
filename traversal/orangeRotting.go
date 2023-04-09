package traversal

func OrangesRotting(grid [][]int) int {
	return orangesRotting(grid)
}

func orangesRotting(grid [][]int) int {
	counter := 0
	queue := make([][2]int, 0)

	adjacents := [][2]int{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
	}

	noFresh := true

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 2 {
				queue = append(queue, [2]int{i, j})
			}

			if grid[i][j] == 1 {
				// if adjacent all 0, then return -1
				noFresh = false
				empty := true
				for _, adjacent := range adjacents {
					x, y := i+adjacent[0], j+adjacent[1]

					if x >= 0 && x < len(grid) && y >= 0 && y < len(grid[0]) && grid[x][y] != 0 {
						empty = false
						break
					}
				}
				if empty {
					return -1
				}
			}
		}
	}

	if noFresh {
		return 0
	}

	// rotten
	var next [][2]int

	for len(queue) > 0 {
		rottingOrange := queue[0]
		queue = queue[1:]

		for _, adjacent := range adjacents {
			x, y := rottingOrange[0]+adjacent[0], rottingOrange[1]+adjacent[1]
			if x >= 0 && x < len(grid) && y >= 0 && y < len(grid[0]) && grid[x][y] == 1 {
				grid[x][y] = 2
				next = append(next, [2]int{x, y})
			}

		}

		if len(queue) == 0 {
			counter++
			queue, next = next, [][2]int{}
		}
	}

	counter--

	// final check

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				return -1
			}
		}
	}

	return counter
}
