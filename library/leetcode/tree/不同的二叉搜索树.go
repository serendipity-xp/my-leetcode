package tree

import (
	"my-leetcode/datastructure/tree"
)

/**
* @DateTime   : 2020/7/21 11:44
* @Author     : xulp
* @Description: 给定一个整数 n，生成所有由 1 ... n 为节点所组成的 二叉搜索树 。
**/

func generateTrees(n int) []*tree.TreeNode {
	if n == 0 {
		return nil
	}
	return generateTree(1, n)
}

func generateTree(start, end int) []*tree.TreeNode {
	if start > end {
		return []*tree.TreeNode{nil}
	}
	result := make([]*tree.TreeNode,0)
	for i := start; i <= end; i++ {
		leftTree := generateTree(start, i-1)
		rightTree := generateTree(i+1, end)
		for _, left := range leftTree {
			for _, right := range rightTree {
				tempTree := &tree.TreeNode{
					Val:   i,
				}
				tempTree.Left = left
				tempTree.Right = right
				result = append(result, tempTree)
			}
		}
	}
	return result
}

func helper(start, end int) []*tree.TreeNode {
	if start > end {
		return []*tree.TreeNode{nil}
	}
	allTrees := []*tree.TreeNode{}
	// 枚举可行根节点
	for i := start; i <= end; i++ {
		// 获得所有可行的左子树集合
		leftTrees := helper(start, i - 1)
		// 获得所有可行的右子树集合
		rightTrees := helper(i + 1, end)
		// 从左子树集合中选出一棵左子树，从右子树集合中选出一棵右子树，拼接到根节点上
		for _, left := range leftTrees {
			for _, right := range rightTrees {
				currTree := &tree.TreeNode{i, nil, nil}
				currTree.Left = left
				currTree.Right = right
				allTrees = append(allTrees, currTree)
			}
		}
	}
	return allTrees
}