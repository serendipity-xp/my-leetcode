package leetcode

/*
* @Time       : 2020/7/18 11:28 上午
* @Author     : xulp
* @Description: 给定三个字符串 s1, s2, s3, 验证 s3 是否是由 s1 和 s2 交错组成的。
 */

func isInterleave(s1 string, s2 string, s3 string) bool {
	ls1, ls2, ls3 := len(s1), len(s2), len(s3)
	if ls1+ls2 != ls3 {
		return false
	}
	f := make([][]bool, ls1+1)
	for i := 0; i < ls1+1; i++ {
		f[i] = make([]bool, ls2+1)
	}
	f[0][0] = true
	for i := 0; i < ls1+1; i++ {
		for j := 0; j < ls2+1; j++ {
			if i > 0 {
				f[i][j] = f[i][j] || (f[i-1][j] && s1[i-1] == s3[i+j-1])
			}
			if j > 0 {
				f[i][j] = f[i][j] || (f[i][j-1] && s2[j-1] == s3[i+j-1])
			}
		}
	}
	return f[ls1][ls2]
}
