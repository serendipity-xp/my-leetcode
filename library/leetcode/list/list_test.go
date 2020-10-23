/**
* @DateTime   : 2020/10/22 14:09
* @Author     : xumamba
* @Description:
**/
package list

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"my-leetcode/datastructure/list"
)

var testCase = &list.ListNode{
	Val: 4,
	Next: &list.ListNode{
		Val: 1,
		Next: &list.ListNode{
			Val: 5,
			Next: &list.ListNode{
				Val:  9,
				Next: nil,
			},
		},
	},
}

func TestDeleteNode(t *testing.T) {
	node := deleteNode(testCase, 1)
	list.PrintListNode(node)

	node2 := deleteNode2(testCase, 1)
	list.PrintListNode(node2)

	listNode := list.NewListNode(4, 1, 5, 9)
	list.PrintListNode(listNode)
}

func TestReverseList(t *testing.T) {
	node := reverseList(testCase)
	list.PrintListNode(node)
}

func TestReversePrint(t *testing.T)  {
	node := reversePrint(testCase)
	fmt.Println(node)

	node2 := reversePrint2(testCase)
	fmt.Println(node2)

	node3 := reversePrint3(testCase)
	fmt.Println(node3)
}

func TestGetDecimalValue(t *testing.T) {
	listNode := list.NewListNode(1, 0, 1)
	value := getDecimalValue(listNode)
	assert.Equal(t, 5, value)

	value = getDecimalValue2(listNode)
	assert.Equal(t, 5, value)
}

func TestKthToLast(t *testing.T)  {
	listNode := list.NewListNode(1, 2, 3)
	result := kthToLast(listNode, 2)
	fmt.Println(result)
}
