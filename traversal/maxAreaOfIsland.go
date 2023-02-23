package traversal

func MaxAreaOfIsland(grid [][]int) int {
	max := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			area := dfsMaxAreaOfIsland(i, j, grid)
			if area > max {
				max = area
			}
		}
	}

	return max
}

func dfsMaxAreaOfIsland(x, y int, grid [][]int) int {
	if (x >= 0 && x < len(grid)) && (y >= 0 && y < len(grid[0])) && grid[x][y] == 1 {
		grid[x][y] = 0
		sum := 1

		sum += dfsMaxAreaOfIsland(x, y+1, grid)
		sum += dfsMaxAreaOfIsland(x, y-1, grid)
		sum += dfsMaxAreaOfIsland(x+1, y, grid)
		sum += dfsMaxAreaOfIsland(x-1, y, grid)

		return sum
	}
	return 0
}
