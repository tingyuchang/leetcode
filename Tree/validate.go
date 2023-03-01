package Tree

import "math"

func isValidBST(root *TreeNode) bool {
	//nums := make([]int, 0)
	//inorderValidate(root, &nums)
	//
	//temp := nums[0]
	//for i := 1; i < len(nums); i++ {
	//	if nums[i] <= temp {
	//		return false
	//	}
	//
	//	temp = nums[i]
	//}
	//
	//return true

	return check(root, math.MinInt, math.MaxInt)
}

func inorderValidate(root *TreeNode, nums *[]int) {
	if root == nil {
		return
	}

	inorderValidate(root.Left, nums)
	*nums = append(*nums, root.Val)
	inorderValidate(root.Right, nums)
}

func check(root *TreeNode, min, max int) bool {
	if root.Val <= min || root.Val >= max {
		return false
	}

	if root.Left != nil {
		if !check(root.Left, min, root.Val) {
			return false
		}
	}

	if root.Right != nil {
		if !check(root.Right, root.Val, max) {
			return false
		}
	}

	return true
}
