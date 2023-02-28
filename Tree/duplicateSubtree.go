package Tree

import "strconv"

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	subtrees := make(map[string]int, 0)
	var result []*TreeNode
	inOrderInDuplicateSubtrees(root, &subtrees, &result)

	return result
}

func inOrderInDuplicateSubtrees(root *TreeNode, subtrees *map[string]int, result *[]*TreeNode) string {
	if root == nil {
		return ""
	}

	left := inOrderInDuplicateSubtrees(root.Left, subtrees, result)
	right := inOrderInDuplicateSubtrees(root.Right, subtrees, result)

	code := "(" + left + ")" + strconv.Itoa(root.Val) + "(" + right + ")"

	(*subtrees)[code]++

	if (*subtrees)[code] == 2 {
		*result = append(*result, root)
	}
	return code

}
