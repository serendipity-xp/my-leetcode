package tree

import (
	"fmt"

	"my-leetcode/tools"
)

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
	if root == nil {
		return nil
	}
	result := make([]int, 0)
	stack := tools.InitStack(tools.Interface)

	for root != nil || stack.Len() != 0 {
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
	if root == nil {
		return nil
	}
	result := make([]int, 0)
	stack := tools.InitStack(tools.Interface)

	var lastVisit *TreeNode
	for root != nil || stack.Len() != 0 {
		for root != nil {
			stack.Push(root)
			root = root.Left
		}
		currNode := stack.Peek().(*TreeNode)
		if currNode.Right == nil || currNode.Right == lastVisit {
			stack.Pop()
			result = append(result, currNode.Val)
			lastVisit = currNode
		} else {
			root = currNode.Right
		}
	}
	return result
}

// 深度优先遍历
func DFSPreorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	dfs(root, &result)
	return result
}

func dfs(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}
	*result = append(*result, root.Val)
	dfs(root.Left, result)
	dfs(root.Right, result)
}

// 分治法
func DivideAndConquer(root *TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}
	// 分治
	left := DivideAndConquer(root.Left)
	right := DivideAndConquer(root.Right)
	// 合并结果
	result = append(result, root.Val)
	result = append(result, left...)
	result = append(result, right...)
	return result

}

/*
DFS 深度搜索（从上到下） 和分治法区别：前者一般将最终结果通过指针参数传入，后者一般递归返回结果最后合并
*/

// BFS层次遍历
func levelOrderOld(root *TreeNode) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		list := make([]int, 0)
		// 记录当前层有多少元素（遍历当前层，再添加下一层）
		length := len(queue)
		for i := 0; i < length; i++ {
			// 出队列
			level := queue[0]
			queue = queue[1:]
			list = append(list, level.Val)
			if level.Left != nil {
				queue = append(queue, level.Left)
			}
			if level.Right != nil {
				queue = append(queue, level.Right)
			}
		}
		result = append(result, list)
	}
	return result
}
