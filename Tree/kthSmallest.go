package Tree

func KthSmallest(root *TreeNode, k int) int {
	res := 0
	current := 0
	preOrderKth(root, k, &current, &res)
	return res
}

func preOrderKth(root *TreeNode, k int, current, res *int) {
	if root == nil || *current > k {
		return
	}

	preOrderKth(root.Left, k, current, res)
	*current++
	if *current == k {
		*res = root.Val
		return
	} else if *current > k {
		return
	}

	preOrderKth(root.Right, k, current, res)
}
