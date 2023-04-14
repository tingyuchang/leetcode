package Array

func BackspaceCompare(s string, t string) bool {
	return backspaceCompare(s, t)
}

func backspaceCompare(s string, t string) bool {
	stack := make([]byte, 0)

	for i := 0; i < len(s); i++ {
		if s[i] == '#' {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
			continue
		}
		stack = append(stack, s[i])
	}

	s = string(stack)
	stack = []byte{}

	for i := 0; i < len(t); i++ {
		if t[i] == '#' {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
			continue
		}
		stack = append(stack, t[i])
	}

	return s == string(stack)
}
