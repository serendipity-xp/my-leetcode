package stack

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	t.Run("有效的括号", func(t *testing.T) {
		isValid("(){}[]")
		edge := hasEdge("a", "b")
		fmt.Println(edge)
	})
}
