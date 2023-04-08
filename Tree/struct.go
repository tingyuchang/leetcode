package Tree

type Node struct {
	Val      int
	Children []*Node
}

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}

type TreeNodeWithNext struct {
	Left  *TreeNodeWithNext
	Right *TreeNodeWithNext
	Val   int
	Next  *TreeNodeWithNext
}

// GenerateTreeBySlice
// input ["1', "", "2"]  empty string means nil node
func GenerateTreeBySlice(input []string) *TreeNode {
	//n := len(input)
	//i := 0
	var node *TreeNode

	return node
}

type QuadNode struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *QuadNode
	TopRight    *QuadNode
	BottomLeft  *QuadNode
	BottomRight *QuadNode
}
