/**
* @DateTime   : 2020/10/22 15:31
* @Author     : xumamba
* @Description:
**/
package list

import (
	"my-leetcode/datastructure/list"
)

//  0ms 2.6MB
func reverseList(head *list.ListNode) *list.ListNode {
	var result *list.ListNode
	for head != nil{
		result = &list.ListNode{Val: head.Val, Next: result}
		head = head.Next
	}
	return result
}


// 0ms 4.7MB
func reversePrint(head *list.ListNode) []int {
	if head == nil{
		return []int{}
	}
	res := reversePrint(head.Next)
	res = append(res, head.Val)
	return res
}

// 4ms  3.2MB
func reversePrint2(head *list.ListNode) []int {
	var result *list.ListNode
	for head != nil{
		result = &list.ListNode{Val: head.Val, Next: result}
		head = head.Next
	}
	res := make([]int, 0)
	for result != nil {
		res = append(res, result.Val)
		result = result.Next
	}
	return res
}

// 4ms 3.1MB
func reversePrint3(head *list.ListNode) []int {
	res := make([]int, 0)
	for head != nil{
		res = append(res, head.Val)
		head = head.Next
	}
	temp := 0
	j := len(res)
	for i:=0; i < j/2; i++{
		temp = res[j-i-1]
		res[j-i-1] = res[i]
		res[i] = temp
	}
	return res
}