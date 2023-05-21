package traversal

func ShortestBridge(grid [][]int) int {
	return shortestBridge(grid)
}

func shortestBridge(grid [][]int) int {
	ans := 0
	// using dfs to build groupMap

	group := 2
	queue := make([][]int, 0)
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				dfsShortestBridge(grid, i, j, group, &queue)
				group += 1
			}
		}
	}

	//fmt.Println(queue, grid)

	directions := [][]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}

	for len(queue) > 0 {
		currentSize := len(queue)
		for i := 0; i < currentSize; i++ {
			currentCell := queue[i]
			for _, direction := range directions {
				x := currentCell[0] + direction[0]
				y := currentCell[1] + direction[1]

				if x >= 0 && x < len(grid) && y >= 0 && y < len(grid[0]) {

					if grid[x][y] == 0 {
						queue = append(queue, []int{x, y})
						grid[x][y] = 2
					} else if grid[x][y] == 3 {
						return ans
					}
				}
			}
		}

		queue = queue[currentSize:]
		ans += 1
	}

	return ans
}

func dfsShortestBridge(grid [][]int, x, y int, group int, queue *[][]int) {
	if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) || grid[x][y] == 0 || grid[x][y] == group {
		return
	}

	grid[x][y] = group
	if group == 2 {
		*queue = append(*queue, []int{x, y})
	}
	dfsShortestBridge(grid, x, y+1, group, queue)
	dfsShortestBridge(grid, x, y-1, group, queue)
	dfsShortestBridge(grid, x+1, y, group, queue)
	dfsShortestBridge(grid, x-1, y, group, queue)
}
