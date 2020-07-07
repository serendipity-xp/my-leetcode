package tree

import (
	"fmt"

	"my-leetcode/tools"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 前序遍历
func PreorderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Println(root.Val)
	PreorderTraversal(root.Left)
	PreorderTraversal(root.Right)
}

// 前序非递归遍历
func PreorderNonRecursionTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	result := make([]int, 0)
	stack := tools.InitStack(tools.Interface)

	for root != nil || stack.Len() != 0 {
		for root != nil {
			// 前序遍历，所以先保留结果
			result = append(result, root.Val)
			stack.Push(root)
			root = root.Left
		}
		// pop
		currNode := stack.Pop().(*TreeNode)
		root = currNode.Right
	}
	return result
}

// 中序非递归遍历
func InorderTraversal(root *TreeNode) []int {
	if root == nil{
		return nil
	}
	result := make([]int,0)
	stack := tools.InitStack(tools.Interface)

	for root != nil || stack.Len() != 0{
		for root != nil {
			stack.Push(root)
			root = root.Left
		}
		currNode := stack.Pop().(*TreeNode)
		result = append(result, currNode.Val)
		root = currNode.Right
	}
	return result
}

// 后序非递归遍历
func PostOrderTraversal(root *TreeNode) []int {
	if root == nil{
		return nil
	}
	result := make([]int, 0)
	stack := tools.InitStack(tools.Interface)

	var lastVisit *TreeNode
	for root != nil || stack.Len() != 0{
		for root != nil{
			stack.Push(root)
			root = root.Left
		}
		currNode := stack.Peek().(*TreeNode)
		if currNode.Right == nil || currNode.Right == lastVisit{
			stack.Pop()
			result = append(result, currNode.Val)
			lastVisit = currNode
		}else{
			root = currNode.Right
		}
	}
	return result
}