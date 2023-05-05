package Tree

func Flatten(root *TreeNode) {
	flatten(root)
}

func flatten(root *TreeNode) {
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
