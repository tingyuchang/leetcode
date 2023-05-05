package Tree

// this problem is not difficult, but the interest part is tree height
// the advanced condition can think about the tree balance

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}

	// recursive solution
	if val < root.Val {
		root.Left = insertIntoBST(root.Left, val)
	}

	if val > root.Val {
		root.Right = insertIntoBST(root.Right, val)
	}

	// or for loop to find the insert node
	/*
		var node *treeNode
		if val < node.Val => node = node.Left
		if val > node.Val => node = node.Right

		until node == nil
	*/

	return root
}
