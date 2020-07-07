package tools

type Stack interface {
	Push(elem interface{})
	Pop() interface{}
	Peek() interface{}
	Len() int
}

type StackInterface struct {
	container []interface{}
}

func (s *StackInterface) Push(elem interface{}) {
	s.container = append(s.container, elem)
}

func (s *StackInterface) Pop() interface{} {
	elem := s.container[len(s.container)-1]
	s.container = s.container[:len(s.container)-1]
	return elem
}

func (s *StackInterface) Peek() interface{} {
	return s.container[len(s.container)-1]
}

func (s *StackInterface) Len() int {
	return len(s.container)
}

type StackInt struct {
	container []int
}

func (s *StackInt) Push(elem interface{}) {
	s.container = append(s.container, elem.(int))
}

func (s *StackInt) Pop() interface{} {
	elem := s.container[len(s.container)-1]
	s.container = s.container[:len(s.container)-1]
	return elem
}

func (s *StackInt) Peek() interface{} {
	return s.container[len(s.container)-1]
}

func (s *StackInt) Len() int {
	return len(s.container)
}

func InitStack(stackType Kind) Stack {
	switch stackType {
	case Int:
		return &StackInt{
			container: make([]int, 0),
		}
	case Interface:
		return &StackInterface{
			container: make([]interface{}, 0),
		}
	}
	return nil
}
