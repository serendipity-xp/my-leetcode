package str

// 从srcStr中匹配aimStr，返回匹配位置索引
func FindNeedleStr(srcStr, aimStr string) (index int) {
	if len(aimStr) == 0 {
		return 0
	}
	var i, j int
	for i = 0; i < len(srcStr)-len(aimStr); i++ {
		for j = 0; j < len(aimStr); j++{
			if srcStr[i+j] != aimStr[j]{
				break
			}
		}
		if len(aimStr) == j{
			return i
		}
	}
	return -1
}
