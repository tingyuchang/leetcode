package LinkedList

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	node := &Node{
		Val:    head.Val,
		Next:   nil,
		Random: nil,
	}

	firstNode := &Node{}
	firstNode.Next = node

	// copy node mapping
	cache := make(map[*Node]*Node)
	temp := head

	cache[temp] = node

	for temp != nil {
		if temp.Next == nil {
			break
		}
		next := &Node{
			Val:    temp.Next.Val,
			Next:   nil,
			Random: nil,
		}
		cache[temp.Next] = next
		node.Next = next
		node = node.Next
		temp = temp.Next
	}

	temp = head
	node = firstNode.Next
	for temp != nil {

		if _, ok := cache[temp.Random]; ok {
			node.Random = cache[temp.Random]
		}

		node = node.Next
		temp = temp.Next
	}

	return firstNode.Next

}
