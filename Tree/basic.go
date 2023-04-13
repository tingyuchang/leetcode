package Tree

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == val {
		return root
	}

	if val < root.Val {
		return searchBST(root.Left, val)
	}

	if val > root.Val {
		return searchBST(root.Right, val)
	}

	return nil
}

func MinDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 1
	}
	var left, right = 100000, 100000
	if root.Left != nil {
		left = MinDepth(root.Left)
	}
	if root.Right != nil {
		right = MinDepth(root.Right)
	}

	return min(left, right) + 1

}
func MaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := MaxDepth(root.Left)
	right := MaxDepth(root.Right)

	return max(left, right) + 1
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func PreOrder(root *TreeNode, nums *[]int) {
	if root != nil {
		*nums = append(*nums, root.Val)
		PreOrder(root.Left, nums)
		PreOrder(root.Right, nums)
	}
}

func InOrder(root *TreeNode, nums *[]int) {
	if root != nil {
		InOrder(root.Left, nums)
		*nums = append(*nums, root.Val)
		InOrder(root.Right, nums)
	}
}

func PostOrder(root *TreeNode, nums *[]int) {
	if root != nil {
		PostOrder(root.Left, nums)
		PostOrder(root.Right, nums)
		*nums = append(*nums, root.Val)
	}
}
