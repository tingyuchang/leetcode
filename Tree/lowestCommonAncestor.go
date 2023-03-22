package Tree

func LowestCommonAncestorOfBST(root, p, q *TreeNode) *TreeNode {
	// left vs. right
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
