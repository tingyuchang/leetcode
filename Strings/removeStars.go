package Strings

func RemoveStars(s string) string {
	return removeStars(s)
}

func removeStars(s string) string {
	stack := make([]byte, 0)

	i := 0
	for i < len(s) {
		if s[i] == '*' {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, s[i])
		}
		i++
	}

	return string(stack)

}
