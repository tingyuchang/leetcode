package Tree

import "fmt"

func connect(root *TreeNodeWithNext) *TreeNodeWithNext {
	if root == nil {
		return nil
	}

	current := root

	for current.Left != nil {
		node := current

		for node != nil {
			node.Left.Next = node.Right
			if node.Next != nil {
				node.Right.Next = node.Next.Left
			}

			node = node.Next
		}

		current = current.Left
	}

	return root
}

func connectV1(root *TreeNodeWithNext) *TreeNodeWithNext {
	if root == nil {
		return nil
	}

	q1 := []*TreeNodeWithNext{root}
	q2 := []*TreeNodeWithNext{}

	current := &q1
	background := &q2

	for len(*current) > 0 {
		node := (*current)[0]
		*current = (*current)[1:]

		// add child to background
		if node.Left != nil {
			*background = append(*background, node.Left)
		}

		if node.Right != nil {
			*background = append(*background, node.Right)
		}

		// link Next to next node
		if len(*current) > 0 {
			node.Next = (*current)[0]
		}

		if len(*current) == 0 {
			current, background = background, current
			fmt.Println(*current)
		}
	}

	return root
}
