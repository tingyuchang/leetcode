package Tree

func findTarget(root *TreeNode, k int) bool {
	return checkTarget(root, k, make(map[int]struct{}))
}

func checkTarget(root *TreeNode, k int, cache map[int]struct{}) bool {
	if root == nil {
		return false
	}

	if root.Left != nil && (root.Val+root.Left.Val) == k {
		return true
	}

	if root.Right != nil && (root.Val+root.Right.Val) == k {
		return true
	}

	if root.Left != nil && root.Right != nil && (root.Right.Val+root.Left.Val) == k {
		return true
	}

	if _, ok := cache[root.Val]; ok {
		return true
	}

	cache[k-root.Val] = struct{}{}
	return checkTarget(root.Left, k, cache) || checkTarget(root.Right, k, cache)

}
