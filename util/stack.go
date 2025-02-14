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

func (s *Stack) PushString(str string) {
	for _, char := range str {
		s.Push(int(char)) // Store ASCII value of each character
	}
}

// Convert ASCII values back to a string
func (s *Stack) PopString(length int) string {
	var result []rune
	for i := 0; i < length; i++ {
		val, ok := s.Pop()
		if !ok {
			break
		}
		result = append([]rune{rune(val)}, result...) // Reverse order since it's a stack
	}
	return string(result)
}
