package tree

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTraversal(t *testing.T) {
	treeObj := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 1,
				},
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 7,
			Left: &TreeNode{
				Val: 6,
			},
			Right: &TreeNode{
				Val: 8,
			},
		},
	}

	t.Run("PreorderTraversal", func(t *testing.T) {
		PreorderTraversal(treeObj)
	})

	t.Run("PreorderNonRecursionTraversal", func(t *testing.T) {
		assert.Equal(t, PreorderNonRecursionTraversal(treeObj), []int{5, 3, 2, 1, 4, 7, 6, 8})
	})

	t.Run("InorderTraversal", func(t *testing.T) {
		assert.Equal(t, InorderTraversal(treeObj), []int{1, 2, 3, 4, 5, 6, 7, 8})
	})

	t.Run("PostOrderTraversal", func(t *testing.T) {
		assert.Equal(t, PostOrderTraversal(treeObj), []int{1, 2, 4, 3, 6, 8, 7, 5})
	})

	t.Run("深度优先遍历", func(t *testing.T) {
		assert.Equal(t, DFSPreorderTraversal(treeObj), []int{5, 3, 2, 1, 4, 7, 6, 8})
	})

	t.Run("分治法", func(t *testing.T) {
		assert.Equal(t, DivideAndConquer(treeObj), []int{5, 3, 2, 1, 4, 7, 6, 8})
	})

	t.Run("层次优先遍历", func(t *testing.T) {
		assert.Equal(t, levelOrder(treeObj), [][]int{{5}, {3, 7}, {2, 4, 6, 8}, {1}})
	})

	t.Run("反向层序优先遍历", func(t *testing.T) {
		assert.Equal(t, levelOrderBottom(treeObj), [][]int{{1}, {2, 4, 6, 8}, {3, 7}, {5}})
	})

	t.Run("Z字形层序遍历", func(t *testing.T) {
		assert.Equal(t, zigzagLevelOrder(treeObj), [][]int{{5}, {7, 3}, {2, 4, 6, 8}, {1}})
	})

	t.Run("分治法遍历二叉树", func(t *testing.T) {
		assert.Equal(t, perOrderTraversal(treeObj), []int{5, 3, 2, 1, 4, 7, 6, 8})
	})

	t.Run("最大深度", func(t *testing.T) {
		assert.Equal(t, maxDepth(treeObj), 4)
	})

	t.Run("是否平衡二叉树", func(t *testing.T) {
		assert.Equal(t, isBalanced(treeObj), true)
	})

	t.Run("最大路径和", func(t *testing.T) {
		assert.Equal(t, maxPathSum(treeObj), 27)
	})

	t.Run("最近公共祖先", func(t *testing.T) {
		// p := &TreeNode{
		// 	Val: 7,
		// }
		// q := &TreeNode{
		// 	Val: 3,
		// }
		// result := &TreeNode{
		// 	Val:   5,
		// 	Left:  q,
		// 	Right: p,
		// }
		// assert.Equal(t, lowestCommonAncestor(treeObj, p, q), result)
	})

	t.Run("是否是有效的二分搜索树", func(t *testing.T) {
		assert.Equal(t, isValidBST(treeObj), true)
		assert.Equal(t, isValidBST2(treeObj), true)
	})

	t.Run("插入二叉树", func(t *testing.T) {
		fmt.Println(insertIntoBST(treeObj,5))
	})
}
