package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTree(t *testing.T) {
	t.Run("搜索插入位置", func(t *testing.T) {
		testCase := []struct {
			input  []int
			target int
			output int
		}{
			{[]int{1, 3, 5, 6}, 5, 2},
			{[]int{1, 3, 5, 6}, 2, 1},
			{[]int{1, 3, 5, 6}, 7, 4},
			{[]int{1, 3, 5, 6}, 0, 0},
			{[]int{}, 2, 0},
		}
		for _, testcase := range testCase {
			// 二分搜索
			assert.Equal(t, searchInsert(testcase.input, testcase.target), testcase.output)
			// 暴力破解
			assert.Equal(t, searchInsertViolence(testcase.input, testcase.target), testcase.output)
		}
	})

}
