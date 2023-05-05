package Tree

type NodeN struct {
	Val   int
	Left  *NodeN
	Right *NodeN
	Next  *NodeN
}

func connectNext(root *NodeN) *NodeN {
	if root == nil {
		return nil
	}

	current := root
	var next, head *NodeN

	for current != nil {
		for current != nil {
			if current.Left != nil {
				if next != nil {
					next.Next = current.Left
				} else {
					head = current.Left
				}
				next = current.Left
			}

			if current.Right != nil {
				if next != nil {
					next.Next = current.Right
				} else {
					head = current.Right
				}
				next = current.Right
			}

			current = current.Next
		}
		current = head
		head = nil
		next = nil
	}

	return root
}

/*

Approach1
BFS: ->

queue := []*Node{root}

    backQueue := make([]*Node, 0)
    for len(queue) > 0 {
        node := queue[0]
        queue = queue[1:]

        if node.Left != nil {
            backQueue = append(backQueue, node.Left)
        }

        if node.Right != nil {
            backQueue = append(backQueue, node.Right)
        }

        if len(queue) == 0 {
            queue, backQueue = backQueue, make([]*Node, 0)
        } else {
            node.Next = queue[0]
        }
    }

    return root

*/
