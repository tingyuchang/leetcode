package CloneGraph

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	cache := make(map[int]*Node)
	queue := make([]*Node, 0)
	queue = append(queue, node)

	for len(queue) > 0 {
		oldNode := queue[0]
		queue = queue[1:]
		// check exist or not
		var copyNode *Node
		if _, ok := cache[oldNode.Val]; ok {
			copyNode = cache[oldNode.Val]
		} else {
			copyNode = &Node{oldNode.Val, make([]*Node, 0)}
			cache[oldNode.Val] = copyNode
		}
		if len(oldNode.Neighbors) != len(copyNode.Neighbors) {
			queue = append(queue, oldNode.Neighbors...)
			for _, v := range oldNode.Neighbors {

				var neighbor *Node

				if _, ok := cache[v.Val]; ok {
					neighbor = cache[v.Val]
				} else {
					neighbor = &Node{v.Val, make([]*Node, 0)}
					cache[v.Val] = neighbor
				}

				copyNode.Neighbors = append(copyNode.Neighbors, neighbor)
			}
		}
	}

	return cache[node.Val]
}
