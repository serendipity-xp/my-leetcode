package sort

import (
	"fmt"
	"testing"
)

/**
* @DateTime   : 2020/12/10
* @Author     : xumamba
* @Description:
**/

// 冒泡排序
func bubbleSort(arr []int) []int {
	length := len(arr)
	for i := 0; i < length; i++ {
		for j := 0; j < length-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

// 选择排序
func selectSort(arr []int) []int {
	length := len(arr)
	for i := 0; i < length; i++ {
		minIndex := i
		for j := i + 1; j < length; j++ {
			if arr[minIndex] > arr[j] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
	return arr
}

// 插入排序
func insertSort(arr []int) []int {
	for i := range arr {
		preIndex := i - 1
		currentValue := arr[i]
		for preIndex >= 0 && arr[preIndex] > currentValue {
			arr[preIndex+1] = arr[preIndex]
			preIndex -= 1
		}
		arr[preIndex+1] = currentValue
	}
	return arr
}

// 希尔排序

// 归并排序

func TestSort(t *testing.T) {
	testCase := []int{6, 2, 3, 1, 0, 5}
	t.Run("bubble sort", func(t *testing.T) {
		fmt.Println(bubbleSort(testCase))
	})
	t.Run("select sort", func(t *testing.T) {
		fmt.Println(selectSort(testCase))
	})
	t.Run("insert sort", func(t *testing.T) {
		fmt.Println(insertSort(testCase))
	})

	t.Run("Sqrt", func(t *testing.T) {
		sqrt := mySqrt(9)
		fmt.Println(sqrt)
	})
}

func reverseLeftWords(s string, n int) string {
	runes := []rune(s)
	reserve(runes)
	reserve(runes[len(runes)-n:])
	reserve(runes[:len(runes)-n])
	return string(runes)
}

func reserve(str []rune) []rune {
	length := len(str) - 1
	for i := 0; i < len(str)/2; i++ {
		str[i], str[length-i] = str[length-i], str[i]
	}
	return str
}

func mySqrt(x int) int {
	result := 0
	left, right := 0, x
	for left <= right {
		mid := (right-left)/2 + left
		temp := mid * mid
		if temp > x {
			right = mid - 1
		} else {
			result = mid
			left = mid + 1
		}
	}
	return result
}
