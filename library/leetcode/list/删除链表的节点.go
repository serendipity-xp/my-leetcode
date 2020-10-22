/**
* @DateTime   : 2020/10/22 13:38
* @Author     : xumamba
* @Description:
**/
package list

import (
	"my-leetcode/datastructure/list"
)

// 4ms  2.8MB
func deleteNode(head *list.ListNode, val int) *list.ListNode {
	dummyHead := &list.ListNode{
		Next: head,
	}
	current := dummyHead
	for current.Next != nil{
		next := current.Next
		if next.Val == val{
			current.Next = next.Next
		}
		current = current.Next
		if current == nil{
			break
		}
	}

	return dummyHead.Next
}

// 4ms  2.8MB
func deleteNode2(head *list.ListNode, val int) *list.ListNode {
	dummyHead := &list.ListNode{
		Next: head,
	}
	current := dummyHead
	for head != nil{
		next := head.Next
		if head.Val == val{
			current.Next = next
			break
		}
		current = head
		head = next
	}
	return dummyHead.Next
}