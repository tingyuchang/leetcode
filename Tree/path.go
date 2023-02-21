package Tree

import (
	"strconv"
)

func binaryTreePaths(root *TreeNode) []string {
	result := make([]string, 0)
	if root == nil {
		return result
	}
	if root.Left == nil && root.Right == nil {
		return []string{strconv.Itoa(root.Val)}
	}
	path := binaryTreePaths(root.Left)
	path = append(path, binaryTreePaths(root.Right)...)

	for i := 0; i < len(path); i++ {
		result = append(result, strconv.Itoa(root.Val)+"->"+path[i])
	}

	return result
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	if root.Val == targetSum && root.Left == nil && root.Right == nil {
		return true
	}

	return hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val)
}
