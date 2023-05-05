package Tree

func LongestZigZag(root *TreeNode) int {
	return longestZigZag(root)
}

func longestZigZag(root *TreeNode) int {
	if root == nil {
		return 0
	}

	res := 0
	zigZag(root.Left, false, &res, 0)
	zigZag(root.Right, true, &res, 0)

	return res
}

// direction true => Right, false => Left
func zigZag(root *TreeNode, direction bool, res *int, currentLength int) {
	if root == nil {
		return
	}

	currentLength++

	if currentLength > *res {
		*res = currentLength
	}

	if direction {
		zigZag(root.Left, false, res, currentLength)
		zigZag(root.Right, true, res, 0)
	} else {
		zigZag(root.Left, false, res, 0)
		zigZag(root.Right, true, res, currentLength)
	}
}
