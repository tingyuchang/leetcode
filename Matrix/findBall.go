package Matrix

func FindBall(grid [][]int) []int {
	return findBall(grid)
}

/*

Pass:

1. if 1 then right adjacent must be 1 and cannot touch right edge until the end of grid
2. if -1 then left adjacent must be -1 and cannot touch left edge until the end of edge

ex :
{1, 1, 1, -1, -1},
{1, 1, 1, -1, -1},
{-1, -1, -1, 1, 1},
{1, 1, 1, 1, -1},
{1, -1, -1, -1, -1},

try {0,0}


*/

func findBall(grid [][]int) []int {
	ans := make([]int, len(grid[0]))

	for i := 0; i < len(ans); i++ {

		direction := grid[0][i]
		position := i
		down := 0

		for down < len(grid) {
			// cannot touch right edge until the end of grid
			if direction == 1 && (position == len(grid[0])-1 || grid[down][position+1] == -1) {
				position = -1
				break
			}
			// cannot touch left edge until the end of edge
			if direction == -1 && (position == 0 || grid[down][position-1] == 1) {
				position = -1
				break
			}

			// the last row
			if down == len(grid)-1 {
				if direction == 1 {
					position++
				} else {
					position--
				}

				direction = grid[down][position]

				break
			}
			down++

			if direction == 1 {
				position++
			} else {
				position--
			}

			direction = grid[down][position]
		}

		ans[i] = position
	}

	return ans
}
