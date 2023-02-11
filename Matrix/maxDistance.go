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

type PointV2 struct {
	X    int
	Y    int
	D    int
	Next *PointV2
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

func MaxDistanceV2(grid [][]int) int {
	n := len(grid)
	count := 0
	result := 0
	head := &PointV2{}
	node := head
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				node.X = i
				node.Y = j
				node.D = 0
				node.Next = &PointV2{}
				node = node.Next
				count++
			}
		}
	}

	if count == 0 || count == n*n {
		return -1
	}

	node2 := head

	for node2 != nil {
		// looking for up, right, down and left cell
		if node2.Y-1 >= 0 {
			q := grid[node2.X][node2.Y-1]
			if q == 0 {
				node.X = node2.X
				node.Y = node2.Y - 1
				node.D = node2.D + 1
				node.Next = &PointV2{}
				node = node.Next
				grid[node2.X][node2.Y-1] = 1
			}
		}
		if node2.X+1 < n {
			q := grid[node2.X+1][node2.Y]
			if q == 0 {
				node.X = node2.X + 1
				node.Y = node2.Y
				node.D = node2.D + 1
				node.Next = &PointV2{}
				node = node.Next
				grid[node2.X+1][node2.Y] = 1
			}
		}

		if node2.Y+1 < n {
			q := grid[node2.X][node2.Y+1]
			if q == 0 {
				node.X = node2.X
				node.Y = node2.Y + 1
				node.D = node2.D + 1
				node.Next = &PointV2{}
				node = node.Next
				grid[node2.X][node2.Y+1] = 1
			}
		}

		if node2.X-1 >= 0 {
			q := grid[node2.X-1][node2.Y]
			if q == 0 {
				node.X = node2.X - 1
				node.Y = node2.Y
				node.D = node2.D + 1
				node.Next = &PointV2{}
				node = node.Next
				grid[node2.X-1][node2.Y] = 1
			}
		}
		if node2.Next == node {
			result = node2.D
			break
		}

		node2 = node2.Next

	}

	return result
}
