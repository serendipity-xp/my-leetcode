package binary

/**
* @DateTime   : 2020/11/26
* @Author     : xumamba
* @Description:
**/

// hammingWeight 汉明重叠，，输入是一个无符号整数，返回其二进制表达式中数字位数为 ‘1’ 的个数
func hammingWeight(num uint32) int {
	result := 0
	for num != 0 {
		if num&1 == 1 {
			result++
		}
		num = num >> 1
	}
	return result
}

func hammingWeight2(num uint32) int {
	result := 0
	for num != 0 {
		num = num & (num - 1)
		result++
	}
	return result
}

// countBits 给定一个非负整数 num。对于 0 ≤ i ≤ num 范围中的每个数字 i ，计算其二进制数中的 1 的数目并将它们作为数组返回。
func countBits(num int) []int {
	result := make([]int, num+1)
	for i := 0; i <= num; i++ {
		count := 0
		j := i
		for j != 0 {
			if j&1 == 1 {
				count++
			}
			j = j >> 1
		}
		result[i] = count
	}
	return result
}

// countBits2 DP
func countBits2(num int) []int {
	result := make([]int, num+1)
	for i := 1; i <= num; i++ {
		// 上一个缺一的元素+1即可
		result[i] = result[i&(i-1)] + 1
	}
	return result
}
