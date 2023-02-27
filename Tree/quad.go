package Tree

func construct(grid [][]int) *QuadNode {

	return recursiveRange(0, 0, len(grid), grid)
}

func recursiveRange(x, y, n int, grid [][]int) *QuadNode {

	if n == 1 {
		return &QuadNode{
			Val:         grid[x][y] == 1,
			IsLeaf:      true,
			TopLeft:     nil,
			TopRight:    nil,
			BottomLeft:  nil,
			BottomRight: nil,
		}
	}
	topLeft := recursiveRange(x, y, n/2, grid)
	topRight := recursiveRange(x, y+n/2, n/2, grid)
	bottomLeft := recursiveRange(x+n/2, y, n/2, grid)
	bottomRight := recursiveRange(x+n/2, y+n/2, n/2, grid)

	children := []*QuadNode{topLeft, topRight, bottomLeft, bottomRight}

	isLeaf := true
	for _, v := range children {
		if v.Val != topLeft.Val || !v.IsLeaf {
			isLeaf = false
			break
		}
	}
	if isLeaf {
		return &QuadNode{
			Val:         topLeft.Val,
			IsLeaf:      true,
			TopLeft:     nil,
			TopRight:    nil,
			BottomLeft:  nil,
			BottomRight: nil,
		}
	}
	return &QuadNode{
		false, false, topLeft, topRight, bottomLeft, bottomRight,
	}
}
