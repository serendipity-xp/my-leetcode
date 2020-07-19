package tree

import (
	"my-leetcode/datastructure/tree"
)

/**
* @Time       : 2020/7/19 5:56 下午
* @Author     : xulp
* @Description: 给定一个二叉树，检查它是否是镜像对称的。
 */

func isSymmetric(root *tree.TreeNode) bool {
	if root == nil {
		return true
	}
	return isMirror(root.Left, root.Right)
}

func isMirror(a, b *tree.TreeNode) bool {
	if a == nil && b == nil {
		return true
	} else if a == nil || b == nil {
		return false
	}
	if a.Val == b.Val {
		return isMirror(a.Left, b.Right) && isMirror(a.Right, b.Left)
	}
	return false
}
