package leetcode

/**
* @DateTime: 2020/7/20 16:48
* @Author     : xulp
* @Description:给定一个已按照升序排列 的有序数组，找到两个数使得它们相加之和等于目标数。
函数应该返回这两个下标值 index1 和 index2，其中 index1 必须小于 index2。

说明:

返回的下标值（index1 和 index2）不是从零开始的。
你可以假设每个输入只对应唯一的答案，而且你不可以重复使用相同的元素。
**/

func twoSum2(numbers []int, target int) []int {
	result := make([]int, 2)
	dict := make(map[int]int)
	for index, value := range numbers {
		other := target - value
		if firstIndex, ok := dict[other]; ok {
			result[0] = firstIndex + 1
			result[1] = index + 1
		}
		dict[value] = index
	}
	return result
}

func twoSumBinarySearch(numbers []int, target int) []int {
	result := make([]int, 2)
	for i := 0; i < len(numbers); i++ {
		other := target - numbers[i]
		left, right := i+1, len(numbers)-1
		for left <= right {
			mid := (right-left)/2 + left
			if numbers[mid] == other {
				result[0] = i + 1
				result[1] = mid + 1
				return result
			}
			if numbers[mid] > other {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}
	return result
}

func twoSumDoublePointer(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	tempSum := 0
	for left <= right {
		tempSum = numbers[left] + numbers[right]
		if tempSum == target {
			return []int{left + 1, right + 1}
		}else if tempSum > target{
			right--
		}else {
			left++
		}
	}
	return nil
}
