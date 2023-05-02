package Array

import "fmt"

func Trap(height []int) int {
	return trapStack(height)
}

func trapStack(height []int) int {
	ans := 0
	current := 0
	stack := make([]int, 0)

	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}

	for current < len(height) {
		for len(stack) != 0 && height[current] > height[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
			distance := current - stack[len(stack)-1] - 1
			boundedHeight := min(height[current], height[stack[len(stack)-1]]) - height[top]
			ans += distance * boundedHeight
		}
		stack = append(stack, current)
		current++
	}

	return ans
}

/*
TO(n)
SO(n)
*/
func trapDP(height []int) int {
	if len(height) == 0 {
		return 0
	}

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}

	ans := 0
	leftMax := make([]int, len(height))
	rightMax := make([]int, len(height))

	leftMax[0] = height[0]
	for i := 1; i < len(height); i++ {
		leftMax[i] = max(leftMax[i-1], height[i])
	}

	rightMax[len(height)-1] = height[len(height)-1]

	for i := len(height) - 2; i >= 0; i-- {
		rightMax[i] = max(rightMax[i+1], height[i])
	}

	for i := 1; i < len(height)-1; i++ {
		ans += min(leftMax[i], rightMax[i]) - height[i]
	}

	return ans
}

/*
TO(n^2)
SO(1)
*/
func trapBruteFroce(height []int) int {
	ans := 0

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}

	for i := 1; i < len(height)-1; i++ {
		left, right := 0, 0

		for j := i; j >= 0; j-- {
			left = max(left, height[j])
		}

		for j := i; j < len(height); j++ {
			right = max(right, height[j])
		}
		fmt.Println("ans: ", ans)
		ans += min(left, right) - height[i]
		fmt.Println(i, ": ", left, right, height[i])

	}

	return ans
}
