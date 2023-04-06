package traversal

import "fmt"

func ClosedIsland(grid [][]int) int {
	return closedIsland(grid)
}

func closedIsland(grid [][]int) int {
	res := 0
	for i := 1; i < len(grid)-1; i++ {
		for j := 1; j < len(grid[i])-1; j++ {
			if grid[i][j] == 0 {
				if dfsClosedIsland(i, j, grid) {
					res++
				}
			}
		}
	}
	return res
}

func dfsClosedIsland(x, y int, grid [][]int) bool {
	// if on the edge, must be 1(water)
	if y == 0 || y == len(grid[0])-1 || x == 0 || x == len(grid)-1 {
		if grid[x][y] == 0 {
			return false
		}

		return true
	}
	fmt.Println(x, y)
	if grid[x][y] == 1 {
		return true
	}

	grid[x][y] = 1

	a := dfsClosedIsland(x+1, y, grid)
	b := dfsClosedIsland(x-1, y, grid)
	c := dfsClosedIsland(x, y+1, grid)
	d := dfsClosedIsland(x, y-1, grid)

	return a && b && c && d
}

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
