package Tree

import "math"

func MaxPathSum(root *TreeNode) int {
	res := math.MinInt
	findMax(root, &res)
	return res
}

func findMax(root *TreeNode, res *int) int {
	if root == nil {
		return 0
	}

	left := findMax(root.Left, res)
	right := findMax(root.Right, res)

	sum := root.Val

	if left > 0 {
		sum += left
	}

	if right > 0 {
		sum += right
	}

	if sum > *res {
		*res = sum
	}

	rootMax := root.Val

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	rootMax += max(0, max(left, right))

	return rootMax
}
