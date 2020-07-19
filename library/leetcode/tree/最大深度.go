package tree

import (
	"my-leetcode/datastructure/tree"
)

/**
* @Time       : 2020/7/19 6:18 下午
* @Author     : xulp
* @Description: 给定一个二叉树，找出其最大深度。
 */

func maxDepth(root *tree.TreeNode) int {
	if root == nil{
		return 0
	}
	return max(maxDepth(root.Left),maxDepth(root.Right))+1
}

func max(x, y int) int{
	if x > y {
		return x
	}
	return y
}