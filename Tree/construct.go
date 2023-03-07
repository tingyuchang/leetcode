package Tree

func buildTree(preorder []int, inorder []int) *TreeNode {
	cache := make(map[int]int)
	currentIndex := 0
	for i, v := range inorder {
		cache[v] = i
	}

	return constructFromPreorderAndInorder(preorder, inorder, 0, len(preorder)-1, &currentIndex, &cache)
}

func constructFromPreorderAndInorder(preorder, inorder []int, start, end int, currentIndex *int, cache *map[int]int) *TreeNode {
	if *currentIndex >= len(preorder) {
		return nil
	}

	// find gead from preorder
	headValue := preorder[*currentIndex]
	head := &TreeNode{Val: headValue}
	*currentIndex++

	target := (*cache)[headValue]

	// find left subtree form inorder

	if target-start > 0 {
		head.Left = constructFromPreorderAndInorder(preorder, inorder, start, target-1, currentIndex, cache)
	}
	// find right subtree from inorder
	if end-target > 0 {
		head.Right = constructFromPreorderAndInorder(preorder, inorder, target+1, end, currentIndex, cache)
	}

	return head
}

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	// find middle as root
	slow, fast := head, head
	var tail *ListNode

	for fast != nil && fast.Next != nil {
		tail = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	root := &TreeNode{Val: slow.Val}
	if slow.Next != nil {
		root.Right = sortedListToBST(slow.Next)
	}
	if tail != nil {
		tail.Next = nil
		root.Left = sortedListToBST(head)
	}
	return root
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	mid := len(nums) / 2
	node := &TreeNode{Val: nums[mid]}
	node.Left = sortedArrayToBST(nums[:mid])
	node.Right = sortedArrayToBST(nums[mid+1:])

	return node
}
