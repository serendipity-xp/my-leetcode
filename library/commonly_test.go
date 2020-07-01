package library

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"testing"
)

// todo attention:
// 参数传递,只能修改，不能新增或者删除原始数据
// 默认 s=s[0:len(s)], 取下限不取上限，左闭右开 [)
// map键需要可比较，不能为slice、map、function
// map值都有默认值，可以直接操作默认值，如：m[age]++ 值由0 变为1
// 比较两个map需要遍历，其中的kv是否相同，因为有默认值关系，所以需要检查val和ok两个值

func Test_commonly_used_libraries(t *testing.T) {
	t.Run("stack", func(t *testing.T) {
		// 栈
		// 创建栈
		stack := make([]int, 0)
		// push压栈
		stack = append(stack, 1)
		// pop弹出
		element := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		fmt.Println(element)
		// 判空
		fmt.Println(len(stack) == 0)
	})

	t.Run("queue", func(t *testing.T) {

		// 队列
		// 创建队列
		queue := make([]int, 0)
		// queue入队
		queue = append(queue, 2)
		// dequeue出队
		element := queue[0]
		queue = queue[1:]
		fmt.Println(element)
		// 判空
		fmt.Println(len(queue) == 0)

	})

	t.Run("dict", func(t *testing.T) {
		// 創建
		dict := make(map[string]int)
		// set key value
		dict["one"] = 1
		dict["two"] = 2
		// 遍历
		for k, v := range dict{
			println(k, v)
		}
		// delete key
		delete(dict, "one")
		fmt.Println(dict)

		val, ok := dict["two"]
		fmt.Println(val, ok)
	})
}

func Test_standard_library(t *testing.T) {

	t.Run("sort", func(t *testing.T) {
		// int排序
		sort.Ints([]int{})
		// 字符串排序
		sort.Strings([]string{})
		// 自定义排序
		// sort.Slice(s, func(i, j int) bool {
		// 	return s[i] < s[j]
		// })
	})

	t.Run("math", func(t *testing.T) {
		_ = math.MaxInt32
		_ = math.MinInt64
	})

	t.Run("copy", func(t *testing.T) {
		// 删除a[i]，可以用 copy 将i+1到末尾的值覆盖到i,然后末尾-1
		a := []int{1,2,3}
		i := 1
		fmt.Println(a)
		copy(a[i:], a[i+1:])
		a = a[:len(a)-1]
		fmt.Println(a)
		a[1] = 4
		a = append(a, 5)
		fmt.Println(a)
	})

	t.Run("type_change", func(t *testing.T) {
		s := "12345"
		fmt.Println(int(s[0]+'a'))
		fmt.Println(s[0])
		fmt.Println(byte(s[0])+'0')

		// 字符串转数字
		atoi, err := strconv.Atoi(s)
		fmt.Println(atoi, err)
		// 数字转字符串
		itoa := strconv.Itoa(atoi)
		fmt.Println(itoa)
	})
}


