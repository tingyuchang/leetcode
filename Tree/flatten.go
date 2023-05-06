package Tree

func Flatten(root *TreeNode) {
	flatten(root)
}

func flatten(root *TreeNode) {
	// preorder
	queue := make([]*TreeNode, 0)
	preorderFlatten(root, &queue)

	for i := 0; i < len(queue)-1; i++ {
		node := queue[i]
		node.Right = queue[i+1]
		node.Left = nil
	}

}

func preorderFlatten(root *TreeNode, queue *[]*TreeNode) {
	if root == nil {
		return
	}
	*queue = append(*queue, root)
	preorderFlatten(root.Left, queue)
	preorderFlatten(root.Right, queue)
}

func flattenRecursive(root *TreeNode) {
	if root == nil {
		return
	}
	var temp *TreeNode
	if root.Right != nil {
		flatten(root.Right)
		temp = root.Right
	}

	if root.Left != nil {
		flatten(root.Left)
		root.Right = root.Left
		root.Left = nil

		current := root.Right

		for current != nil {
			if current.Right == nil {
				break
			}
			current = current.Right
		}

		current.Right = temp

	}
}
