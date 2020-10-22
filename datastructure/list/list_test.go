package list

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

/**
* @DateTime   : 2020/7/22 16:41
* @Author     : xulp
* @Description:
**/

func TestList(t *testing.T) {
	testCase := []struct {
		input, output *ListNode
	}{
		{
			input: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val:  2,
						Next: nil,
					},
				},
			},
			output: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val:  2,
					Next: nil,
				},
			},
		},
		{
			input: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val:  3,
								Next: nil,
							},
						},
					},
				},
			},
			output: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val:  3,
						Next: nil,
					},
				},
			},
		},
	}
	t.Run("删除重复元素", func(t *testing.T) {
		for _, tc := range testCase {
			assert.Equal(t, deleteDuplicates(tc.input), tc.output)
			assert.Equal(t, deleteDuplicatesRecursion(tc.input), tc.output)
			assert.Equal(t, deleteDuplicatesDoublePointer(tc.input), tc.output)
		}
	})
	t.Run("删除重复元素只保留未重复的数字", func(t *testing.T) {
		result := deleteDuplicates2(testCase[0].input)
		fmt.Println(result)
	})
}
