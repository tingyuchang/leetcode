package Array

func ValidateStackSequences(pushed []int, popped []int) bool {
	return validateStackSequences(pushed, popped)
}

func validateStackSequences(pushed []int, popped []int) bool {
	stack := make([]int, 0)

	for len(popped) > 0 {

		pop := popped[0]
		// find pop value in stack
		if len(stack) > 0 && stack[len(stack)-1] == pop {
			popped = popped[1:]
			stack = stack[:len(stack)-1]

			continue
		}

		// if not found, try push element until meet same value
		for len(pushed) > 0 {
			push := pushed[0]
			pushed = pushed[1:]
			stack = append(stack, push)

			if push == pop {
				break
			}
		}

		// try again, if not found, then return false

		if len(stack) > 0 && stack[len(stack)-1] == pop {
			popped = popped[1:]
			stack = stack[:len(stack)-1]
		} else {
			return false
		}

	}
	return true
}
