package Math

func TwoSumII(numbers []int, target int) []int {
	return twoSumII(numbers, target)
}

func twoSumII(numbers []int, target int) []int {

	l := 0
	r := len(numbers) - 1

	for r > l {
		sum := numbers[l] + numbers[r]
		if sum == target {
			return []int{l + 1, r + 1}
		}

		if sum > target {
			r--
		} else {
			l++
		}
	}

	return nil
}
