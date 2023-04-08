package Tree

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil && root2 == nil {
		return nil
	}

	var root *TreeNode

	if root1 == nil && root2 != nil {
		root = root2
		root.Left = mergeTrees(nil, root2.Left)
		root.Right = mergeTrees(nil, root2.Right)
	} else if root1 != nil && root2 == nil {
		root = root1
		root.Left = mergeTrees(root1.Left, nil)
		root.Right = mergeTrees(root1.Right, nil)
	} else {
		root = &TreeNode{Val: root1.Val + root2.Val}
		root.Left = mergeTrees(root1.Left, root2.Left)
		root.Right = mergeTrees(root1.Right, root2.Right)
	}

	return root
}
