/**
* @DateTime   : 2020/10/23 15:29
* @Author     : xumamba
* @Description:
**/
package list

import (
	"my-leetcode/datastructure/list"
)

// 0 ms	2 MB
func getDecimalValue(head *list.ListNode) int {
	result := 0
	for head != nil{
		result = (result << 1) | head.Val
		head = head.Next
	}
	return result
}

// 0 ms	2 MB
func getDecimalValue2(head *list.ListNode) int {
	result := 0
	for head != nil{
		result = result*2 + head.Val
		head = head.Next
	}
	return result
}