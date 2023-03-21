package Tree

func IsSubtree(root *TreeNode, subRoot *TreeNode) bool {

	if root == nil && subRoot != nil {
		return false
	}

	if isSameTree(root, subRoot) {
		return true
	}
	if isSameTree(root.Left, subRoot) {
		return true
	}

	if isSameTree(root.Right, subRoot) {
		return true
	}

	return IsSubtree(root.Right, subRoot) || IsSubtree(root.Left, subRoot)
}
func levelOrderBottom(root *TreeNode) [][]int {
	queue1 := make([]*TreeNode, 0)
	queue2 := make([]*TreeNode, 0)

	result := make([][]int, 0)

	if root == nil {
		return result
	}

	var currentQueue *[]*TreeNode
	var backQueue *[]*TreeNode

	queue1 = append(queue1, root)
	currentQueue = &queue1
	backQueue = &queue2
	temp := make([]int, 0)
	for len(queue1)+len(queue2) != 0 {
		// queue POP
		node := (*currentQueue)[0]
		*currentQueue = append([]*TreeNode{}, (*currentQueue)[1:]...)
		temp = append(temp, node.Val)
		if node.Left != nil {
			*backQueue = append(*backQueue, node.Left)
		}
		if node.Right != nil {
			*backQueue = append(*backQueue, node.Right)
		}

		if len(*currentQueue) == 0 {
			currentQueue, backQueue = backQueue, currentQueue

			result = append([][]int{temp}, result...)
			temp = make([]int, 0)
		}
	}
	return result
}

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
