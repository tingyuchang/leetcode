package Tree

func MaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := MaxDepth(root.Left)
	right := MaxDepth(root.Right)

	return max(left, right) + 1
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