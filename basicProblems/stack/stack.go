package stack

type Stack struct {
	data []interface{}
}

func (s *Stack) Push(record interface{}) {
	s.data = append(s.data, record)
}

func (s *Stack) Pop() interface{} {
	record := s.data[:len(s.data)-1]
	return record
}

func (s *Stack) length() int {
	return len(s.data)
}

type FakeQueue struct {
	stackA *Stack // input pool
	StackB *Stack // output pool
}

func (fq *FakeQueue) Add(record interface{}) {
	fq.stackA.Push(record)
}

func (fq *FakeQueue) Remove() interface{} {
	if fq.StackB.length()  == 0 {
		// output A's Data to B
		for fq.stackA.length() > 0 {
			fq.StackB.Push(fq.stackA.Pop())
		}
	}
	return fq.StackB.Pop()
}