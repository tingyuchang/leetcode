package Tree

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	if !isBalanced(root.Left) || !isBalanced(root.Right) {
		return false
	}

	if abs(MaxDepth(root.Left)-MaxDepth(root.Right)) > 1 {
		return false
	}

	return true
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}