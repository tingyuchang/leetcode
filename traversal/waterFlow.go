package traversal

func PacificAtlantic(heights [][]int) [][]int {
	result := make([][]int, 0)

	rows := len(heights)
	columns := len(heights[0])

	pacific := make([][]bool, rows)
	atlantic := make([][]bool, rows)

	for i := 0; i < rows; i++ {
		pacific[i] = make([]bool, columns)
		atlantic[i] = make([]bool, columns)
	}

	for i := 0; i < columns; i++ {
		dfsPacificAtlantic(0, i, heights, heights[0][i], pacific)
		dfsPacificAtlantic(rows-1, i, heights, heights[rows-1][i], atlantic)
	}

	for i := 0; i < rows; i++ {
		dfsPacificAtlantic(i, 0, heights, heights[i][0], pacific)
		dfsPacificAtlantic(i, columns-1, heights, heights[i][columns-1], atlantic)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			if pacific[i][j] == true && atlantic[i][j] == true {
				result = append(result, []int{i, j})
			}
		}
	}

	return result
}

// from edge to center
func dfsPacificAtlantic(x, y int, heights [][]int, height int, visited [][]bool) {

	if (x >= 0 && x < len(heights)) && (y >= 0 && y < len(heights[x])) && (heights[x][y] >= height) && (visited[x][y] != true) {
		visited[x][y] = true
		dfsPacificAtlantic(x-1, y, heights, heights[x][y], visited)
		dfsPacificAtlantic(x, y+1, heights, heights[x][y], visited)
		dfsPacificAtlantic(x+1, y, heights, heights[x][y], visited)
		dfsPacificAtlantic(x, y-1, heights, heights[x][y], visited)
	}

	return

}
