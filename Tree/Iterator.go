package Tree

type BSTIterator struct {
	nodes []*TreeNode
}

func BSTIteratorConstructor(root *TreeNode) BSTIterator {
	nodes := make([]*TreeNode, 0)

	inOrderBSTIterator(root, &nodes)

	return BSTIterator{
		nodes: nodes,
	}
}

func inOrderBSTIterator(root *TreeNode, nodes *[]*TreeNode) {
	if root == nil {
		return
	}

	*nodes = append(*nodes, root)
	inOrderBSTIterator(root.Left, nodes)
}

func (this *BSTIterator) Next() int {
	node := this.nodes[len(this.nodes)-1]
	this.nodes = this.nodes[:len(this.nodes)-1]
	inOrderBSTIterator(node.Right, &this.nodes)
	return node.Val
}

func (this *BSTIterator) HasNext() bool {
	return len(this.nodes) > 0
}
