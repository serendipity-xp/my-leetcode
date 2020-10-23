/**
* @DateTime   : 2020/10/22 17:31
* @Author     : xumamba
* @Description:
**/
package stack


type Stack []int

func (s *Stack) Push(elem int) {
	*s = append(*s, elem)
}

func (s *Stack) Pop() int {
	elem := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return elem
}

func (s *Stack) Peek() int {
	return (*s)[(*s).Len()-1]
}

func (s *Stack) Len() int {
	return len(*s)
}

func NewStack() Stack {
	return make([]int, 0)
}