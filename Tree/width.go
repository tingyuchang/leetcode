package Tree

func WidthOfBinaryTree(root *TreeNode) int {
	return widthOfBinaryTree(root)
}

/*
max width := 2 ^ high

BFS
if has Left the width = 2^hight+1
if has Right the winth = 2^hight+2


*/

func widthOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	backQueue := make([]*TreeNode, 0)
	cache := make(map[*TreeNode]int)
	cache[root] = 1
	res := 1
	start := 1
	end := 0
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		parentIndex := cache[node]

		if node.Left != nil {
			backQueue = append(backQueue, node.Left)
			cache[node.Left] = 2*(parentIndex-1) + 1
		}

		if node.Right != nil {
			backQueue = append(backQueue, node.Right)
			cache[node.Right] = 2*(parentIndex-1) + 2
		}
		if len(queue) == 0 {
			//fmt.Println(cache)
			end = cache[node]
			queue, backQueue = backQueue, make([]*TreeNode, 0)

			if end-start+1 > res {
				res = end - start + 1
			}
			if len(queue) != 0 {
				start, end = cache[queue[0]], 0
			}

		}

	}

	return res
}
