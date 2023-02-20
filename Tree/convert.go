package Tree

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
