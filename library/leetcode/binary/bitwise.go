package binary

/**
* @DateTime   : 2020/11/26
* @Author     : xumamba
* @Description:
**/

// rangeBitwiseAnd 给定范围 [m, n]，其中 0 <= m <= n <= 2147483647，返回此范围内所有数字的按位与（包含 m, n 两端点）。
func rangeBitwiseAnd(m int, n int) int {
	var count int
	for m != n{
		m >>= 1
		n >>= 1
		count++
	}
	return m << count
}
