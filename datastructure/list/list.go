package list
/**
 * @DateTime   : 2020/7/22 16:35
 * @Author     : xulp
 * @Description: list
 **/

import (
	"fmt"
)


// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(val ...int) *ListNode {
	l := len(val)
	if l == 0 {
		return nil
	}
	ln := &ListNode{}
	tempHead := ln
	for i, v := range val{
		tempHead.Val = v
		if i != l-1{
			tempHead.Next = &ListNode{}
			tempHead = tempHead.Next
		}
	}
	return ln
}

func PrintListNode(head *ListNode) {
	result := make([]int, 0)
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	fmt.Println(result)
}

/*
	有序链表 删除重复元素
*/
func deleteDuplicates(head *ListNode) *ListNode {
	current := head
	for current != nil {
		for current.Next != nil && current.Val == current.Next.Val {
			current.Next = current.Next.Next
		}
		current = current.Next
	}
	return head
}

// 8ms 3.1MB

func deleteDuplicatesRecursion(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	head.Next = deleteDuplicates(head.Next)
	if head.Val == head.Next.Val {
		head = head.Next
	}
	return head
}

// 4ms 3.2MB

func deleteDuplicatesDoublePointer(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	front, back := head, head.Next
	for back != nil {
		if front.Val == back.Val {
			back = back.Next
		} else {
			front.Next = back
			front = back
			back = back.Next
		}
	}
	front.Next = nil
	return head
}

// 0ms  3.1MB

/*
	给定一个排序链表，删除所有含有重复数字的节点，只保留原始链表中 没有重复出现 的数字。
*/
func deleteDuplicates2(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	front, back := head, head.Next
	for back != nil {
		if front.Val == back.Val {
			front = back.Next
			back = front.Next
		} else {
			front.Next = back
			front = back
			back = back.Next
		}
	}
	return head
}
