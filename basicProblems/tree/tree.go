package tree

type Node struct {
	Value    int
	Children []*Node
}

func (n *Node) Add(data int)  {
	n.Children = append(n.Children, &Node{data, nil})
}

func (n *Node) Remove(data interface{})  {
	for i,v := range n.Children {
		if (*v).Value == data {
			n.Children = append(n.Children[:i], n.Children[i+1:]...)
			return
		}
	}
}

type Tree struct {
	Root *Node
}

func (t *Tree) BreadthFirst() []*Node  {
	var result []*Node
	temp := []*Node{t.Root}

	for len(temp) > 0 {
		result = append(result, temp[0])
		if len(temp[0].Children) > 0 {
			temp = append(temp, temp[0].Children...)
		}
		temp = temp[1:]
	}

	return result
}


func (t *Tree) DepthFirst() []*Node {
	var result []*Node
	temp := []*Node{t.Root}

	for len(temp) > 0 {
		result = append(result, temp[0])
		if len(temp[0].Children) > 0 {
			temp = append(temp[:1], append(temp[0].Children, temp[1:]...)...)
		}
		temp = temp[1:]
	}
	return result
}