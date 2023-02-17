package Tree

func InorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	inorder(root, &result)
	return result
}

func inorder(root *TreeNode, nums *[]int) {
	if root == nil {
		return
	}
	inorder(root.Left, nums)
	*nums = append(*nums, root.Val)
	inorder(root.Right, nums)
}
