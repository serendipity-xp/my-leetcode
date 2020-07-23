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

	t.Run("两数之和", func(t *testing.T) {
		numbers := []int{2, 7, 11, 15}
		target := 9
		fmt.Println(twoSum2(numbers, target))
		fmt.Println(twoSumBinarySearch(numbers, target))
		fmt.Println(twoSumDoublePointer(numbers, target))
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

	t.Run("戳气球", func(t *testing.T) {
		testCase := []struct {
			input  []int
			output int
		}{{
			[]int{3, 1, 5, 8},
			167,
		},
		}
		for _, tc := range testCase {
			assert.Equal(t, maxCoins(tc.input), tc.output)
		}
	})

	t.Run("旋转数组最小数字", func(t *testing.T) {
		testCase := []struct {
			input  []int
			output int
		}{
			{
				[]int{1, 3, 5},
				1,
			},
			{
				[]int{3, 4, 5, 1, 2},
				1,
			},
			{
				[]int{2, 2, 2, 0, 1},
				0,
			},
		}
		for _, tc := range testCase {
			assert.Equal(t, minArray(tc.input), tc.output)
			assert.Equal(t, minArrayBinarySearch(tc.input), tc.output)
		}
	})

	t.Run("最小路径和", func(t *testing.T) {
		grid := [][]int{
			{1, 3, 4, 8},
			{3, 2, 2, 4},
			{5, 7, 1, 9},
			{2, 3, 2, 3},
		}
		sum := minPathSum(grid)
		fmt.Println(sum)
	})
}
