package leetcode

import (
	"fmt"
	"testing"
)

func Test_Leetcode(t *testing.T) {
	t.Run("两数之和", func(t *testing.T) {
		nums := []int{0, -2, 0,2,-5} // [-1,-2,-3,-4,-5] -8
		target := 0

		result := twoSumHash(nums, target)
		fmt.Println(result)
	})
}
