package tree

import (
	"my-leetcode/datastructure/tree"
)

var result [][]int

func levelOrder(root *tree.TreeNode) [][]int {
	if root == nil{
		return nil
	}
	result = make([][]int, 0)
	dfs(root, 0)
	return result
}

func dfs(root *tree.TreeNode, level int) {
	if root == nil{
		return
	}
	if level == len(result){
		result = append(result, []int{})
	}
	result[level] = append(result[level], root.Val)
	dfs(root.Left, level + 1)
	dfs(root.Right, level + 1)
}

