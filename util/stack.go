package util

type Stack struct {
	Stack []int
	top   int
}

func (s *Stack) Push(value int) {
	s.Stack = append(s.Stack, value)
	s.top++
}

func (s *Stack) Pop() (int, bool) {
	if s.top == 0 {
		return 0, false
	}
	s.top--
	value := s.Stack[s.top]
	s.Stack = s.Stack[:s.top]
	return value, true
}

func (s *Stack) Peek() (int, bool) {
	if s.top == 0 {
		return 0, false
	}
	return s.Stack[s.top-1], true
}
