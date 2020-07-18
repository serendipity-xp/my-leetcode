package tree

/**
* @Time       : 2020/7/18 4:27 下午
* @Author     : xulp
* @Description: 如果两个树在结构上相同，并且节点具有相同的值，则认为它们是相同的。
 */

import (
	"my-leetcode/datastructure/tree"
)

func isSameTree(p *tree.TreeNode, q *tree.TreeNode) bool {
	if p == nil && q == nil {
		return true
	} else if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
