package tree

import (
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
}
