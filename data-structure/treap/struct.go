package treap

import (
	"fmt"
	"math/rand"
)

type TreapNode struct {
	Parent   *TreapNode
	Left     *TreapNode
	Right    *TreapNode
	Val      int
	Priority int
}

/*
	 left

				r                      R
	          /   \                 /     \
			L  	   R        ->     r      RR
				/     \          /    \
	           RL     RR       L       RL
*/
func (t *TreapNode) rotateLeft() *TreapNode {
	right := t.Right
	t.Right = right.Left
	right.Left = t
	return right
}

/*
		 right

					r                      L
		          /   \                 /     \
				L  	   R        ->     LL      r
	          /   \          				/    \
		    LL    LR       				   LR     R
*/
func (t *TreapNode) rotateRight() *TreapNode {
	left := t.Left
	t.Left = t.Left.Right
	left.Right = t
	return left
}

func newTreapNode(val int) *TreapNode {
	return &TreapNode{
		Parent:   nil,
		Left:     nil,
		Right:    nil,
		Val:      val,
		Priority: rand.Int() % 100,
	}
}

func Insert(root *TreapNode, val int) *TreapNode {
	if root == nil {
		root = newTreapNode(val)
		fmt.Printf("val: %d, priority: %d\n", root.Val, root.Priority)
		return root
	}

	if val < root.Val {
		root.Left = Insert(root.Left, val)

		if root.Left != nil && root.Left.Priority < root.Priority {
			fmt.Println("trigger right rotate ", root.Val)
			root = root.rotateRight()
		}
	} else {
		root.Right = Insert(root.Right, val)
		if root.Right != nil && root.Right.Priority < root.Priority {
			fmt.Println("trigger left rotate ", root.Val)
			root = root.rotateLeft()
		}
	}

	fmt.Println("-----")
	PrintTreap(root, 0)
	fmt.Println("----")
	return root
}

func PrintTreap(root *TreapNode, space int) {
	queue := []*TreapNode{root}

	for len(queue) != 0 {
		currentSize := len(queue)

		for i := 0; i < currentSize; i++ {
			currentNode := queue[i]
			if currentNode == nil {
				fmt.Printf("(nil)")
				continue
			}

			fmt.Printf("%d(%d)", currentNode.Val, currentNode.Priority)
			queue = append(queue, currentNode.Left)
			queue = append(queue, currentNode.Right)
		}

		queue = queue[currentSize:]
		fmt.Println()

	}

}
