package basicProblems

type ListNode struct {
	Val interface{}
	Next *ListNode
}

func CreateLinkList(data []interface{}) *ListNode {
	dummyNode := &ListNode{}
	tmpNode := dummyNode

	for _,v := range data {
		tmpNode.Next = &ListNode{
			Val: v,
		}
		tmpNode = tmpNode.Next
	}
	return dummyNode.Next
}
