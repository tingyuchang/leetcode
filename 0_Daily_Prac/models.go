package __Daily_Prac

type Node struct {
	Val       int
	Neighbors []*Node
}

type ListNode struct {
	Val  int
	Next *ListNode
}

type NodeR struct {
	Val    int
	Next   *NodeR
	Random *NodeR
}
