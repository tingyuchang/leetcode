package traversal

func NumEnclaves(grid [][]int) int {
	return numEnclaves(grid)
}

func numEnclaves(grid [][]int) int {
	res := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				sum, notBoundary := dfsNumEnclaves(grid, i, j, 0)
				if notBoundary {
					res += sum
				}
			}
		}
	}

	return res
}

func dfsNumEnclaves(grid [][]int, x, y, sum int) (int, bool) {
	// 0 is Water
	// 1 is land
	//fmt.Println(x, y, sum)

	if x >= 0 && x < len(grid) && y >= 0 && y < len(grid[x]) && grid[x][y] == 0 {
		return sum, true
	}
	if x <= 0 || x >= len(grid)-1 || y <= 0 || y >= len(grid[x])-1 {
		return 0, false
	}

	grid[x][y] = 0
	sum += 1

	aSum, aBool := dfsNumEnclaves(grid, x+1, y, sum)
	//fmt.Println("A", x+1, y, aSum, aBool)
	if aBool {
		sum = aSum
	}

	bSum, bBool := dfsNumEnclaves(grid, x-1, y, sum)
	//fmt.Println("B", x-1, y, aSum, aBool)
	if bBool {
		sum = bSum
	}

	cSum, cBool := dfsNumEnclaves(grid, x, y+1, sum)
	//fmt.Println("C", x, y+1, aSum, aBool)
	if cBool {
		sum = cSum
	}

	dSum, dBool := dfsNumEnclaves(grid, x, y-1, sum)
	//fmt.Println("D", x, y-1, aSum, aBool)
	if dBool {
		sum = dSum
	}

	if aBool && bBool && cBool && dBool {
		return sum, true
	}

	return 0, false

}
