package Tree

import "math"

func isCompleteTree(root *TreeNode) bool {
	q1 := make([]*TreeNode, 0)
	q2 := make([]*TreeNode, 0)
	currentQueue := &q1
	backQueue := &q2
	q1 = append(q1, root)
	count := 0
	height := 0
	stop := false
	for len(*currentQueue) > 0 {
		node := (*currentQueue)[0]
		*currentQueue = (*currentQueue)[1:len(*currentQueue)]
		count++

		if node.Left != nil {
			if stop == true {
				return false
			}
			*backQueue = append(*backQueue, node.Left)
		} else {
			stop = true
		}

		if node.Right != nil {
			if stop == true {
				return false
			}
			*backQueue = append(*backQueue, node.Right)
		} else {
			stop = true
		}

		// exchange
		if len(*currentQueue) == 0 {
			// check
			if count != int(math.Pow(float64(2), float64(height))) && len(*backQueue) != 0 {
				return false
			}
			currentQueue, backQueue = backQueue, currentQueue
			height++
			stop = false
			count = 0
		}
	}

	return true
}
