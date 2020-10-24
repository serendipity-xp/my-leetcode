/**
* @DateTime   : 2020/10/23 15:53
* @Author     : xumamba
* @Description:
**/
package list

import (
	"my-leetcode/datastructure/list"
)

func kthToLast(head *list.ListNode, k int) int {
	fast := head
	slow := head
	for k > 0 {
		fast = fast.Next
		k--
	}
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}
	return slow.Val
}

func getKthFromEnd(head *list.ListNode, k int) *list.ListNode {
	fast, slow := head, head
	for k > 0 {
		fast = fast.Next
		k--
	}
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}
