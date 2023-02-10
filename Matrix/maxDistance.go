package Matrix

import (
	"fmt"
	"math"
)

type Point struct {
	X int
	Y int
}

func (p Point) distance(anotherP Point) int {
	return int(math.Abs(float64(p.X-anotherP.X)) + math.Abs(float64(p.Y-anotherP.Y)))
}
func MaxDistance(grid [][]int) int {
	max := 0
	n := len(grid)
	var water []Point
	var land []Point

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				// land
				land = append(land, Point{i, j})
			} else {
				// water
				water = append(water, Point{i, j})
			}
		}
	}

	if len(land) == 0 || len(water) == 0 {
		return -1
	}

	for _, w := range water {
		// find land's nearest water
		x := w.X
		y := w.Y
		shortest := (n - 1) * 2
		for i := x; i < n; i++ {
			short := (n - 1) * 2
			for j := y; j < n; j++ {
				q := grid[i][j]
				if q == 1 {
					qp := Point{i, j}
					if w.distance(qp) < short {
						fmt.Printf("%v %v\n", w, qp)
						short = w.distance(qp)
					}
					break
				}
			}

			for j := y - 1; j >= 0; j-- {
				q := grid[i][j]
				if q == 1 {
					qp := Point{i, j}
					if w.distance(qp) < short {
						fmt.Printf("%v %v\n", w, qp)
						short = w.distance(qp)
					}
					break
				}
			}
			if short < shortest {
				shortest = short
			}
		}

		for i := x - 1; i >= 0; i-- {
			short := (n - 1) * 2
			for j := y; j < n; j++ {
				q := grid[i][j]
				if q == 1 {
					qp := Point{i, j}
					if w.distance(qp) < short {
						fmt.Printf("%v %v\n", w, qp)
						short = w.distance(qp)
					}
					break
				}
			}

			for j := y - 1; j >= 0; j-- {
				q := grid[i][j]
				if q == 1 {
					qp := Point{i, j}
					if w.distance(qp) < short {
						fmt.Printf("%v %v\n", w, qp)
						short = w.distance(qp)
					}
					break
				}
			}
			if short < shortest {
				shortest = short
			}

		}

		if shortest > max {
			max = shortest
		}
	}

	return max
}
