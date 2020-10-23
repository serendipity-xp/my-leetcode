/**
* @DateTime   : 2020/10/22 17:53
* @Author     : xumamba
* @Description:
**/
package stack

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	stack := NewStack()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	fmt.Println(stack.Peek())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Peek())
	fmt.Println(stack.Len())
}