package tools

/**
* @DateTime   : 2020/7/23 11:42
* @Author     : xulp
* @Description:
**/

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
