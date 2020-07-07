package tree

import (
	"fmt"
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
	stack := make([]*TreeNode, 0)

	for root != nil || len(stack) != 0 {
		for root != nil {
			// 前序遍历，所以先保留结果
			result = append(result, root.Val)
			stack = append(stack, root)
			root = root.Left
		}
		// pop
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		root = node.Right
	}
	return result
}

// 中序非递归遍历
func InorderTraversal(root *TreeNode) []int {
	if root == nil{
		return nil
	}
	result := make([]int,0)
	stack := make([]*TreeNode,0)

	for root != nil || len(stack) != 0{
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		currNode := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
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
	stack := make([]*TreeNode, 0)

	var lastVisit *TreeNode
	for root != nil || len(stack) != 0{
		for root != nil{
			stack = append(stack, root)
			root = root.Left
		}
		currNode := stack[len(stack)-1]
		if currNode.Right == nil || currNode.Right == lastVisit{
			stack = stack[:len(stack)-1]
			result = append(result, currNode.Val)
			lastVisit = currNode
		}else{
			root = currNode.Right
		}
	}
	return result
}