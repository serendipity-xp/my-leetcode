package tree

import (
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
	t.Run("相同的树", func(t *testing.T) {
		assert.Equal(t, isSameTree(treeObj1, treeObj2), false)
	})
}
