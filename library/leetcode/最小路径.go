package leetcode

import (
	"my-leetcode/tools"
)

/**
* @DateTime   : 2020/7/23 11:35
* @Author     : xulp
* @Description: 给定一个包含非负整数的 m x n 网格，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
说明：每次只能向下或者向右移动一步。
示例:
输入:
[
  [1,3,1],
  [1,5,1],
  [4,2,1]
]
输出: 7
解释: 因为路径 1→3→1→1→1 的总和最小。
**/

func minPathSum(grid [][]int) int {
	dp := make([][]int, len(grid))
	for i, v := range grid {
		dp[i] = make([]int, len(v))
	}
	dp[0][0] = grid[0][0]
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if i > 0 && j > 0 {
				dp[i][j] = tools.Min(dp[i][j-1], dp[i-1][j]) + grid[i][j]
			} else if i > 0 {
				dp[i][j] = dp[i-1][j] + grid[i][j]
			} else if j > 0 {
				dp[i][j] = dp[i][j-1] + grid[i][j]
			}
		}
	}
	return dp[len(grid)-1][len(grid[len(grid)-1])-1]
}

// 8ms 4.4MB
