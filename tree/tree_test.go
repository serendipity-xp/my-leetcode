package tree

import (
	"fmt"
	"testing"
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

	t.Run("", func(t *testing.T) {
		result := PreorderNonRecursionTraversal(treeObj)
		fmt.Println(result)
	})
}
