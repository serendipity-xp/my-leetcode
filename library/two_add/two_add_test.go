package two_add

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddTwoNumbers(t *testing.T) {

	l1 := ListNode{
		Val: 2,
		Next: &ListNode{
			Val:  4,
			Next: &ListNode{Val: 3},
		},
	}
	l2 := ListNode{
		Val: 5,
		Next: &ListNode{
			Val:  6,
			Next: &ListNode{Val: 4},
		},
	}

	numbers := addTwoNumbers(&l1, &l2)
	assert.Equal(t, 7, numbers.Val)
	assert.Equal(t, 0, numbers.Next.Val)
	assert.Equal(t, 8, numbers.Next.Next.Val)
}
