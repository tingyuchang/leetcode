package shortestAlternatingPaths

type Edge struct {
	From  int
	To    int
	Color int
	D     int
	Next  *Edge
}

const (
	BLUE = iota
	RED
)

func ShortestAlternatingPaths(n int, redEdges [][]int, blueEdges [][]int) []int {
	result := make([]int, n)

	for i, _ := range result {
		if i == 0 {
			result[i] = 0
		} else {
			result[i] = -1
		}
	}

	edgeIndex := make(map[int]map[int][]*Edge, 2)
	edgeIndex[RED] = make(map[int][]*Edge, n)
	edgeIndex[BLUE] = make(map[int][]*Edge, n)

	var node, last *Edge
	// setup
	for _, v := range redEdges {
		edge := &Edge{v[0], v[1], RED, 0, nil}
		if edge.From == 0 {
			edge.D = 1
			if node == nil && last == nil {
				node = edge
				last = edge
			} else {
				last.Next = edge
				last = last.Next
			}

			if result[edge.To] == -1 {
				result[edge.To] = 1
			}
		} else {
			// using From as index
			edgeIndex[RED][v[0]] = append(edgeIndex[RED][v[0]], edge)
		}
	}

	for _, v := range blueEdges {
		edge := &Edge{v[0], v[1], BLUE, 0, nil}
		if edge.From == 0 {
			edge.D = 1
			if node == nil && last == nil {
				node = edge
				last = edge
			} else {
				last.Next = edge
				last = last.Next
			}
			if result[edge.To] == -1 {
				result[edge.To] = 1
			}
		} else {
			// using From as index
			edgeIndex[BLUE][v[0]] = append(edgeIndex[BLUE][v[0]], edge)
		}
	}

	/*
		{
			// 0 is BLUE
			0: {
				1: []*Edge{From: 0, To:1, Color: Blue, D: 0, Next: nil}
				2: []*Edge{}
				}
			// 1 is RED
			1: {}
		}
	*/

	for node != nil {
		// add edge to last
		if node.Color == BLUE {
			for _, v := range edgeIndex[RED][node.To] {
				if v.D == 0 {
					v.D = node.D + 1
					last.Next = v
					last = v
				}
				if result[v.To] == -1 {
					result[v.To] = v.D
				} else {
					if result[v.To] > v.D {
						result[v.To] = v.D
					}
				}
			}
		} else {
			for _, v := range edgeIndex[BLUE][node.To] {
				if v.D == 0 {
					v.D = node.D + 1
					last.Next = v
					last = v
				}
				if result[v.To] == -1 {
					result[v.To] = v.D
				} else {
					if result[v.To] > v.D {
						result[v.To] = v.D
					}
				}
			}
		}

		node = node.Next
	}

	return result
}
