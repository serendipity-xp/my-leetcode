package tree

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"my-leetcode/datastructure/tree"
)

func TestTree(t *testing.T) {
	treeObj1 := &tree.TreeNode{
		Val: 5,
		Left: &tree.TreeNode{
			Val: 3,
			Left: &tree.TreeNode{
				Val: 2,
				Left: &tree.TreeNode{
					Val: 1,
				},
			},
			Right: &tree.TreeNode{
				Val: 4,
			},
		},
		Right: &tree.TreeNode{
			Val: 7,
			Left: &tree.TreeNode{
				Val: 6,
			},
			Right: &tree.TreeNode{
				Val: 8,
			},
		},
	}
	treeObj2 := &tree.TreeNode{
		Val: 5,
		Left: &tree.TreeNode{
			Val: 3,
			Left: &tree.TreeNode{
				Val: 2,
				Left: &tree.TreeNode{
					Val: 1,
				},
			},
			Right: &tree.TreeNode{
				Val: 4,
			},
		},
		Right: &tree.TreeNode{
			Val: 7,
			Left: &tree.TreeNode{
				Val: 6,
			},
			Right: &tree.TreeNode{
				Val: 9,
			},
		},
	}
	treeObj3 := &tree.TreeNode{
		Val: 1,
		Left: &tree.TreeNode{
			Val: 2,
			Left: &tree.TreeNode{
				Val: 3,
			},
			Right: &tree.TreeNode{
				Val: 4,
			},
		},
		Right: &tree.TreeNode{
			Val: 2,
			Left: &tree.TreeNode{
				Val: 4,
			},
			Right: &tree.TreeNode{
				Val: 3,
			},
		},
	}
	t.Run("相同的树", func(t *testing.T) {
		assert.Equal(t, isSameTree(treeObj1, treeObj2), false)
	})

	t.Run("对称二叉树", func(t *testing.T) {
		assert.Equal(t, isSymmetric(treeObj1), false)
		assert.Equal(t, isSymmetric(treeObj3), true)
	})

	t.Run("最大深度", func(t *testing.T) {
		assert.Equal(t, maxDepth(treeObj1),4)
	})

	t.Run("不同的二叉搜索树", func(t *testing.T) {
		fmt.Println(generateTrees(8))
		fmt.Println(len([]*tree.TreeNode{nil}))
		fmt.Println(len(make([]*tree.TreeNode,0)))
	})

	t.Run("层序遍历", func(t *testing.T) {
		fmt.Println(levelOrder(treeObj1))
	})
}
