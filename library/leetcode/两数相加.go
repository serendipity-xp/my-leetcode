package leetcode

/*
   输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
   输出：7 -> 0 -> 8
   原因：342 + 465 = 807
*/

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		result = new(ListNode)
		dummy  = result
		carry  = 0
	)
	for l1 != nil || l2 != nil || carry != 0 {
		dummy.Next = new(ListNode)
		dummy = dummy.Next
		if l1 != nil{
			carry += l1.Val
			l1 = l1.Next
		}
		if l2 != nil{
			carry += l2.Val
			l2 = l2.Next
		}
		dummy.Val = carry % 10
		carry = carry / 10
	}
	return result.Next
}

func addTwoNumbersOld(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		result = new(ListNode)
		dummy = result
		carry int
	)
	for {

		if l1 == nil && l2 == nil {
			if carry > 0 {
				dummy.Next = &ListNode{Val: carry}
			}
			break
		}

		var x, y, sum = 0, 0, 0

		if l1 != nil {
			x = l1.Val
		}
		if l2 != nil {
			y = l2.Val
		}

		sum = x + y + carry
		if sum >= 10 {
			carry = 1
		} else {
			carry = 0
		}

		dummy.Next = &ListNode{
			Val:  sum % 10,
			Next: nil,
		}

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
		dummy = dummy.Next
	}

	return result.Next
}
