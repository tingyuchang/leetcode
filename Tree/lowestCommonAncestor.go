package Tree

func LowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// left vs. right
	if p.Val <= root.Val && q.Val >= root.Val {
		return root
	}

	if p.Val < root.Val && q.Val < root.Val {
		return LowestCommonAncestor(root.Left, p, q)
	}
	if p.Val > root.Val && q.Val > root.Val {
		return LowestCommonAncestor(root.Right, p, q)
	}

	return root
}
