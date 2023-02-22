package BFS

func NumIslands(grid [][]byte) int {
	sum := -1

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' {
				dfsIsland(i, j, grid)
				sum++
			}
		}
	}

	return sum + 1
}

func dfsIsland(x, y int, grid [][]byte) {
	if (y >= 0 && y < len(grid[0])) && (x >= 0 && x < len(grid)) && grid[x][y] == '1' {
		grid[x][y] = '0'
		dfsIsland(x, y+1, grid)
		dfsIsland(x, y-1, grid)
		dfsIsland(x+1, y, grid)
		dfsIsland(x-1, y, grid)
	}

}
