package stack


type Stack struct{
	Array []interface{}
}

func (s Stack) IsEmpty() bool{
	return len(s.Array) == 0
}

func (s Stack) Count() int{
	return len(s.Array)
}

func (s *Stack) Push(element interface{}){
	s.Array = append(s.Array, element)
}
