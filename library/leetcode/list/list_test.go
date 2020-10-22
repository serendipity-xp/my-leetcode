/**
* @DateTime   : 2020/10/22 14:09
* @Author     : xumamba
* @Description:
**/
package list

import (
	"testing"

	"my-leetcode/datastructure/list"
)

func TestDeleteNode(t *testing.T) {
	testCase := &list.ListNode{
		Val: 4,
		Next: &list.ListNode{
			Val: 1,
			Next: &list.ListNode{
				Val: 1,
				Next: &list.ListNode{
					Val:  9,
					Next: nil,
				},
			},
		},
	}

	node := deleteNode(testCase, 1)
	list.PrintListNode(node)

	node2 := deleteNode2(testCase, 1)
	list.PrintListNode(node2)

	listNode := list.NewListNode(4, 1, 5, 9)
	list.PrintListNode(listNode)
}
