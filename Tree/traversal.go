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

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if q == nil {
		return false
	}

	if p == nil {
		return false
	}

	if p.Val != q.Val {
		return false
	}

	if !isSameTree(p.Left, q.Left) {
		return false
	}

	if !isSameTree(p.Right, q.Right) {
		return false
	}

	return true
}

func preorderTraversal(root *TreeNode) []int {
	nums := make([]int, 0)

	preOrder(root, &nums)

	return nums
}

func preOrder(node *TreeNode, nums *[]int) {
	if node == nil {
		return
	}
	*nums = append(*nums, node.Val)
	preOrder(node.Left, nums)
	preOrder(node.Right, nums)
}

func postorderTraversal(root *TreeNode) []int {
	nums := make([]int, 0)

	postOrder(root, &nums)

	return nums
}

func postOrder(node *TreeNode, nums *[]int) {
	if node == nil {
		return
	}
	postOrder(node.Left, nums)
	postOrder(node.Right, nums)
	*nums = append(*nums, node.Val)
}
