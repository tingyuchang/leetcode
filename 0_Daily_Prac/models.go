package __Daily_Prac

import "strconv"

type ListNode struct {
	Val  int
	Next *ListNode
}

type Node struct {
	Val       int
	Neighbors []*Node
}

type NodeR struct {
	Val    int
	Next   *NodeR
	Random *NodeR
}

func (n *ListNode) Show() string {
	dummyNode := &ListNode{}
	var s string
	dummyNode.Next = n
	for dummyNode.Next != nil {
		s = s + strconv.Itoa(dummyNode.Next.Val)
		dummyNode = dummyNode.Next
	}
	return s
}

func GenerateNodeFromArray(nums []int) *ListNode {
	head := &ListNode{}
	node := head

	for i := 0; i < len(nums); i++ {
		node.Val = nums[i]
		if i == len(nums)-1 {
			break
		}
		next := &ListNode{}
		node.Next = next
		node = node.Next
	}

	return head
}

func (n *ListNode) String() string {
	res := ""

	node := n

	for node != nil {
		res += strconv.Itoa(node.Val) + "->"
		node = node.Next
	}
	if len(res) > 2 {
		res = string([]byte(res)[:len(res)-2])
	}

	return res
}
