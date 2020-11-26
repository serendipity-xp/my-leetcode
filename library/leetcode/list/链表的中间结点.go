/**
* @DateTime   : 2020/10/24 17:15
* @Author     : xumamba
* @Description:
**/
package list

import (
	"my-leetcode/datastructure/list"
)

func middleNode(head *list.ListNode) *list.ListNode {
	var container []*list.ListNode
	for head != nil{
		container = append(container, head)
		head = head.Next
	}
	return container[len(container)/2]
}