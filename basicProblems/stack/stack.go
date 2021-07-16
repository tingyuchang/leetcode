package stack

type stack struct {
	data []interface{}
}

func (s *stack) push(record interface{}) {
	s.data = append(s.data, record)
}

func (s *stack) pop() interface{} {
	record := s.data[:len(s.data)-1]
	return record
}