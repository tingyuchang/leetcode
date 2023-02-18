package Tree

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isMirror(root.Left, root.Right)
}

func isMirror(l, r *TreeNode) bool {
	if l == nil && r == nil {
		return true
	}
	if l == nil && r != nil || l != nil && r == nil {
		return false
	}
	if !isMirror(l.Left, r.Right) {
		return false
	}
	if !isMirror(l.Right, r.Left) {
		return false
	}

	if l.Val != r.Val {
		return false
	}
	return true
}
