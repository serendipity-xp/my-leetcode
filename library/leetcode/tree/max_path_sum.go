package tree

import (
	"math"

	"my-leetcode/datastructure/tree"
)

/**
* @DateTime   : 2020/12/7
* @Author     : xumamba
* @Description:
**/

func maxPathSum(root *tree.TreeNode) int {
	maxSum := math.MinInt32
	var dfs func(root *tree.TreeNode)int
	dfs = func(root *tree.TreeNode) int {
		if root == nil{
			return 0
		}
		l := dfs(root.Left)
		r := dfs(root.Right)
		sum := l + r + root.Val
		maxSum = max(sum, maxSum)

		innerSum := root.Val + max(l, r)
		return max(innerSum, 0)
	}
	dfs(root)
	return maxSum
}

