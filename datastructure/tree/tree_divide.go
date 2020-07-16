package tree

import (
	"fmt"
	"math"
)

/*
分治法应用
	先分别处理局部，再合并结果
适用场景
	快速排序
	归并排序
	二叉树相关问题
分治法模板
	递归返回条件
	分段处理
	合并结果

func traversal(root *TreeNode) ResultType  {
    // nil or leaf
    if root == nil {
        // do something and return
    }

    // Divide
    ResultType left = traversal(root.Left)
    ResultType right = traversal(root.Right)

    // Conquer
    ResultType result = Merge from left and right

    return result
}
*/

// v2 通过分治法遍历二叉树
func perOrderTraversal(root *TreeNode) []int {
	result := divideAndConquer(root)
	return result
}

func divideAndConquer(root *TreeNode) []int {
	result := make([]int, 0)
	// 递归返回条件
	if root == nil {
		return result
	}
	// 分段处理
	left := divideAndConquer(root.Left)
	right := divideAndConquer(root.Right)
	// 合并结果
	result = append(result, root.Val)
	result = append(result, left...)
	result = append(result, right...)
	return result
}

// 归并排序
func MergeSort(nums []int) []int {
	return mergeSort(nums)
}

func mergeSort(nums []int) []int {
	// 递归返回条件
	if len(nums) <= 1 {
		return nums
	}
	// 分治法：divide 分为两段
	mid := len(nums) / 2
	left := mergeSort(nums[:mid])
	right := mergeSort(nums[mid:])
	// 合并两段数据
	result := merge(left, right)
	// 递归需要返回结果用于合并
	return result
}

func merge(left, right []int) (result []int) {
	// 两边数组合并游标
	l, r := 0, 0
	// 注意下标越界
	for l < len(left) && r < len(right) {
		// 谁小合并谁
		if left[l] > right[r] {
			result = append(result, right[r])
			r++
		} else {
			result = append(result, left[l])
			l++
		}
	}
	// 剩余部分合并
	result = append(result, left[l:]...)
	result = append(result, right[r:]...)
	return
}

// 快速排序
func QuickSort(nums []int) []int {
	// 把一个数组分为左右两段，左段小于右段，类似分治法没有合并过程
	quickSort(nums, 0, len(nums)-1)
	return nums
}

// 原地交换，所以传入交换索引
func quickSort(nums []int, start, end int) {
	if start < end {
		// 分治法，divide
		pivot := partition(nums, start, end)
		quickSort(nums, 0, pivot-1)
		quickSort(nums, pivot+1, end)
	}
}

// 分区
func partition(nums []int, start, end int) int {
	p := nums[end]
	i := start
	for j := start; j < end; j++ {
		if nums[j] < p {
			swap(nums, i, j)
			i++
		}
	}
	// 把中间的值换为用于比较的基准值
	swap(nums, i, end)
	return i
}

// 交换
func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}

/*
快排由于是原地交换所以没有合并过程 传入的索引是存在的索引（如：0、length-1 等），越界可能导致崩溃
*/

// 最大深度
func maxDepth(root *TreeNode) int {
	// 递归返回条件
	if root == nil {
		return 0
	}
	// 分治处理
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	if left == -1 || right == -1 || math.Abs(float64(left-right)) > 1 {
		return -1
	}
	// 合并结果
	if left > right {
		return left + 1
	}
	return right + 1
}

// 判断是否是高度平衡的二叉树 (使用二义性的返回值：-1：代表不平衡； 非-1：代表树的高度）
func isBalanced(root *TreeNode) bool {
	if maxDepth(root) == -1 {
		return false
	}
	return true
}

// 最大路径和
type ResultType struct {
	SinglePath int // 保存单边最大值
	MaxPath    int // 保存最大值（单边或者两个单边+根的值）
}

func maxPathSum(root *TreeNode) int {
	result := helper(root)
	fmt.Println("单边最大值：", result.SinglePath)
	return result.MaxPath
}

func helper(root *TreeNode) ResultType {
	// check
	if root == nil {
		return ResultType{
			SinglePath: 0,
			MaxPath:    -(1 << 31),
		}
	}
	// Divide
	left := helper(root.Left)
	right := helper(root.Right)

	// Conquer
	result := ResultType{}
	// 求单边最大值
	if left.SinglePath > right.SinglePath {
		result.SinglePath = max(left.SinglePath+root.Val, 0)
	} else {
		result.SinglePath = max(right.SinglePath+root.Val, 0)
	}

	// 求两边加根最大值
	maxPath := max(right.MaxPath, left.MaxPath)
	result.MaxPath = max(maxPath, left.SinglePath+right.SinglePath+root.Val)
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 最近公共祖先
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// check
	if root == nil {
		return root
	}
	// 相等 直接返回root节点即可
	if root == p || root == q {
		return root
	}
	// Divide
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)


	// Conquer
	// 左右两边都不为空，则根节点为祖先
	if left != nil && right != nil {
		return root
	}
	if left != nil {
		return left
	}
	if right != nil {
		return right
	}
	return nil
}
