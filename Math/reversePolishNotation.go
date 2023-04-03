package Math

import (
	"strconv"
)

func EvalRPN(tokens []string) int {
	return evalRPN(tokens)
}

type stack []int

func (s *stack) pop() int {
	res := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return res
}

func (s *stack) push(n int) {
	*s = append(*s, n)
}

func evalRPN(tokens []string) int {
	s := stack{}

	for i := 0; i < len(tokens); i++ {

		num, err := strconv.Atoi(tokens[i])

		if err != nil {
			a, b := s.pop(), s.pop()
			switch tokens[i] {
			case "+":
				s.push(a + b)
			case "-":
				s.push(b - a)
			case "*":
				s.push(a * b)
			case "/":
				s.push(b / a)
			}
		} else {
			s.push(num)
		}

	}

	return s.pop()
}
