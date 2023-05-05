package Tree

func LowestCommonAncestorOfBST(root, p, q *TreeNode) *TreeNode {
	// Left vs. Right
	if p.Val <= root.Val && q.Val >= root.Val {
		return root
	}

	if p.Val < root.Val && q.Val < root.Val {
		return LowestCommonAncestorOfBST(root.Left, p, q)
	}
	if p.Val > root.Val && q.Val > root.Val {
		return LowestCommonAncestorOfBST(root.Right, p, q)
	}

	return root
}

func LowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root.Val == p.Val || root.Val == q.Val {
		return root
	}

	left := LowestCommonAncestor(root.Left, p, q)
	right := LowestCommonAncestor(root.Right, p, q)

	if left != nil && right != nil {
		return root
	}

	if left == nil {
		return right
	}
	return left
}

// LowestCommonAncestorWithUnCertain
// 1. Check both a, b exist in the tree, if not exist, can direct return -1
// 2. dfs traversal base on 2 targets exist.
// 2a. : if root == a || root == b, we can return root
// 2b. if root is not a || b, them try root.Left & root.Right
// 2c. if both Left & Right find a, b, the root will be answer, otherwise return which one has Value
func LowestCommonAncestorWithUnCertain(root *TreeNode, a, b int) int {
	if !findAncesator(root, a) || !findAncesator(root, b) {
		return -1
	}
	return lcaHelper(root, a, b).Val
}

func lcaHelper(root *TreeNode, a, b int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == a || root.Val == b {
		return root
	}
	left := lcaHelper(root.Left, a, b)
	right := lcaHelper(root.Right, a, b)

	if left == nil {
		return right
	}

	if right == nil {
		return left
	}

	return root

}

func findAncesator(root *TreeNode, n int) bool {
	if root == nil {
		return false
	}

	if root.Val == n {
		return true
	}

	return findAncesator(root.Left, n) || findAncesator(root.Right, n)
}
