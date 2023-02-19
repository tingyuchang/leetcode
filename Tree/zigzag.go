package Tree

import "fmt"

type Queue struct {
	Q       []*TreeNode
	S       []*TreeNode
	Current bool // false is Q, true is S
}

func (q *Queue) Len() int {
	return len(q.Q) + len(q.S)
}

func (q *Queue) PushQ(node *TreeNode) {
	q.Q = append(q.Q, node)
}
func (q *Queue) PushS(node *TreeNode) {
	q.S = append(q.S, node)
}

func (q *Queue) PopQ() *TreeNode {
	node := q.Q[len(q.Q)-1]
	q.Q = q.Q[:len(q.Q)-1]
	return node
}

func (q *Queue) PopS() *TreeNode {
	node := q.S[len(q.S)-1]
	q.S = q.S[:len(q.S)-1]
	return node
}

// Check return true when  current status changed
func (q *Queue) Check() bool {
	if !q.Current {
		if len(q.Q) == 0 {
			q.Current = true
			return true
		}
	} else {
		if len(q.S) == 0 {
			q.Current = false
			return true
		}
	}

	return false
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	nums := make([][]int, 0)
	if root == nil {
		return nums
	}
	queue := Queue{}

	queue.PushQ(root)
	queue.Current = false
	var temp []int
	for queue.Len() > 0 {
		// Current is false: Queue
		// Current is true: Stack
		var node *TreeNode
		// fmt.Printf("Q: %v, S%v\n", len(queue.Q), len(queue.S))
		if !queue.Current {
			node = queue.PopQ()
			// Push to Stack

			if node.Left != nil {
				queue.PushS(node.Left)
			}
			if node.Right != nil {
				queue.PushS(node.Right)
			}

		} else {
			node = queue.PopS()
			if node.Right != nil {
				queue.PushQ(node.Right)
			}
			if node.Left != nil {
				queue.PushQ(node.Left)
			}

		}
		fmt.Println(node.Val)
		temp = append(temp, node.Val)
		if queue.Check() {
			nums = append(nums, temp)
			temp = make([]int, 0)
		}

	}

	return nums
}
