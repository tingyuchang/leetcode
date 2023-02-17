package Tree

import (
	"fmt"
)

func MinDiffInBST(root *TreeNode) int {
	nums := &[]int{}
	inOrder(root, nums)
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}

	minVal := 106
	for i := 1; i < len(*nums); i++ {
		diff := (*nums)[i] - (*nums)[i-1]
		minVal = min(minVal, diff)
		fmt.Println(diff)
	}

	return minVal
}

func inOrder(node *TreeNode, nums *[]int) {
	if node != nil {
		inOrder(node.Left, nums)
		*nums = append(*nums, node.Val)
		inOrder(node.Right, nums)
	}
}

func MinDiffInBSTV2(root *TreeNode) int {
	minVal := 100000
	preVal := -100000

	inOrderV2(root, &minVal, &preVal)

	return minVal
}

func inOrderV2(node *TreeNode, minVal *int, preVal *int) {
	if node != nil {
		inOrderV2(node.Left, minVal, preVal)
		if *minVal > node.Val-*preVal {
			*minVal = node.Val - *preVal
		}
		*preVal = node.Val
		inOrderV2(node.Right, minVal, preVal)
	}
}
