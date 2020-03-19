package two_add

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		res  = ListNode{}
		curr = &res
		c    int
	)

	for {

		var v1, v2 int

		if l1 != nil {
			v1 = l1.Val
		}
		if l2 != nil {
			v2 = l2.Val
		}

		val := v1 + v2 + c
		c = 0

		// if need to carry, adjust current and bring carry over to next iteration.
		if r := val % 10; val > 9 {
			c = 1
			val = r
		}

		curr.Val = val

		// we are done when we have no nexts left.
		if (l1 == nil || l1.Next == nil) && (l2 == nil || l2.Next == nil) {
			// in case we are done but we have a carry, we create another node with it as a value
			if c > 0 {
				n := &ListNode{Val: c}
				curr.Next = n
				curr = curr.Next
			}
			break
		}

		curr.Next = &ListNode{}
		curr = curr.Next

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}

	return &res
}
