package Math

import "strconv"

func calculate(s string) int {
	ms := myStack{}
	for i := 0; i < len(s); i++ {
		input := s[i]

		switch input {
		case '+':
			a := ms.Pop()
			b := ms.Pop()
			ms.Puch(a + b)
		case '-':
			a := ms.Pop()
			b := ms.Pop()
			ms.Puch(b - a)
		case '*':
			a := ms.Pop()
			b := ms.Pop()
			ms.Puch(a * b)
		case '/':
			a := ms.Pop()
			b := ms.Pop()
			ms.Puch(b / a)
		case ' ':
			break
		case ')':
			
		default:
			ms.Puch(input)
		}
	}
	ans, _ := strconv.Atoi(string(ms.Pop()))
	return ans
}

type myStack struct {
	data []byte
}

func (s *myStack) Pop() byte {
	val := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return val
}

func (s *myStack) Puch(n byte) {
	s.data = append(s.data, n)
}
