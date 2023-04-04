package LinkedList

import (
	"leetcode/Tree"
)

func IsSubPath(head *ListNode, root *Tree.TreeNode) bool {
	return check(head, root)
}

func check(head *ListNode, root *Tree.TreeNode) bool {
	if root == nil {
		return false
	}

	if dfsIsSubPath(head, root) {
		return true
	}

	if check(head, root.Left) {
		return true
	}

	if check(head, root.Right) {
		return true
	}

	return false
}

func dfsIsSubPath(head *ListNode, root *Tree.TreeNode) bool {
	if head == nil {
		return true
	}

	if root == nil || head.Val != root.Val {
		return false
	}

	if dfsIsSubPath(head.Next, root.Left) {
		return true
	}

	if dfsIsSubPath(head.Next, root.Right) {
		return true
	}

	return false
}

