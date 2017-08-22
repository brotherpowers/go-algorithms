package stack

import "errors"

type Stack struct {
	Array []int
}

func (s Stack) IsEmpty() bool {
	return len(s.Array) == 0
}

func (s Stack) Count() int {
	return len(s.Array)
}

func (s *Stack) Push(element int) {
	s.Array = append(s.Array, element)
}

func (s *Stack) Pop() (int, error) {
	length := len(s.Array)
	if length == 0 {
		return 0, errors.New("Stack Underflow")
	}
	element := s.Array[length-1]
	s.Array = s.Array[:length-1]
	return element, nil
}
