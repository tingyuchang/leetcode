package Tree

type NodeP struct {
	Value  int
	Left   *NodeP
	Right  *NodeP
	parent *NodeP
}

func GenerateNodeP(s string, target int) *NodeP {
	codex := Codec{}
	// "20,9,25,5,12,null,null,null,null,11,14"
	root := codex.DeserializeP(s)

	return searchNodeP(root, target)
}

func searchNodeP(root *NodeP, target int) *NodeP {
	if root == nil {
		return nil
	}

	if root.Value == target {
		return root
	}

	if target > root.Value {
		return searchNodeP(root.Right, target)
	}

	return searchNodeP(root.Left, target)
}

func FindInOrderSuccessor(inputNode *NodeP) *NodeP {
	if inputNode.Right == nil {
		if inputNode.parent == nil {
			return inputNode
		} else {
			parent := inputNode.parent
			child := inputNode
			for parent != nil && child == parent.Right {
				child = parent
				parent = child.parent
			}
			return parent
		}
	}
	//
	return findSmallest(inputNode.Right)
}

func findSmallest(root *NodeP) *NodeP {
	if root.Left == nil {
		return root
	}
	return findSmallest(root.Left)
}
