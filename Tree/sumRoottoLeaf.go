package Tree

import (
	"fmt"
	"math"
)

func sumNumbers(root *TreeNode) int {
	res := make([]int, 0)
	sum := 0
	dfsSumNumbers(root, &res, &sum)
	return sum
}

func dfsSumNumbers(root *TreeNode, res *[]int, sum *int) {
	if root == nil {
		return
	}
	*res = append(*res, root.Val)
	if root.Left == nil && root.Right == nil {
		leaf := 0
		count := 0
		for i := len(*res) - 1; i >= 0; i-- {
			leaf += (*res)[i] * int(math.Pow(float64(10), float64(count)))
			count++
		}
		fmt.Println(leaf)

		*sum += leaf
	}
	dfsSumNumbers(root.Left, res, sum)
	dfsSumNumbers(root.Right, res, sum)
	*res = (*res)[:len(*res)-1]

}