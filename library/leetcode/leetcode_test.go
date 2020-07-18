package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Leetcode(t *testing.T) {
	t.Run("两数之和", func(t *testing.T) {
		nums := []int{0, -2, 0, 2, -5} // [-1,-2,-3,-4,-5] -8
		target := 0

		result := twoSumHash(nums, target)
		fmt.Println(result)
	})

	t.Run("两数相加", func(t *testing.T) {
		l1 := &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val:  3,
					Next: nil,
				},
			},
		}
		l2 := &ListNode{
			Val: 5,
			Next: &ListNode{
				Val: 6,
				Next: &ListNode{
					Val:  4,
					Next: nil,
				},
			},
		}

		result := addTwoNumbersOld(l1, l2)
		fmt.Println(result)
		fmt.Println(result.Val)
		if result.Next != nil {
			fmt.Println(result.Next.Val)
			if result.Next.Next != nil {
				fmt.Println(result.Next.Next.Val)
			}
		}
	})

	t.Run("交错字符串", func(t *testing.T) {
		testCase := []struct {
			s1, s2, s3 string
			result     bool
		}{
			{"aabcc", "dbbca", "aadbbcbcac", true},
			{"aabcc", "dbbca", "aadbbbaccc", false},
			{"", "", "", true},
			{"a", "", "a", true},
		}
		for _, tc := range testCase {
			assert.Equal(t, isInterleave(tc.s1, tc.s2, tc.s3), tc.result)
		}
	})

	t.Run("thinker", func(t *testing.T) {
		sliceObj := make([]bool, 1)
		sliceObj[0] = true
	})
}
