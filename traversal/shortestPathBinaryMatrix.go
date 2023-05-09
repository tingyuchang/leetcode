package traversal

import "fmt"

func ShortestPathBinaryMatrix(grid [][]int) int {
	return shortestPathBinaryMatrix(grid)
}

func shortestPathBinaryMatrix(grid [][]int) int {
	if grid[0][0] == 1 {
		return -1
	}

	n := len(grid)

	if n == 1 {
		return 1
	}

	queue := []point{point{0, 0, 1}}
	grid[0][0] = 1

	directions := [][]int{
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{0, -1},
		{0, 1},
		{1, -1},
		{1, 0},
		{1, 1},
	}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		for _, direction := range directions {
			newX := current.x + direction[0]
			newY := current.y + direction[1]

			if newX >= 0 && newX < n &&
				newY >= 0 && newY < n &&
				grid[newX][newY] == 0 {

				if newX == n-1 && newY == n-1 {
					fmt.Println(current.step)
					return current.step + 1
				}
				queue = append(queue, point{x: newX, y: newY, step: current.step + 1})
				grid[newX][newY] = 1
			}
		}
	}
	return -1
}

type point struct {
	x    int
	y    int
	step int
}
