package binary

/**
* @DateTime   : 2020/11/26
* @Author     : xumamba
* @Description:
**/

// reverseBits 颠倒给定的32位无符号整数的二进制位
func reverseBits(num uint32) uint32  {
	var result uint32
	var pow = 31
	for num != 0 {
		result += (num & 1) << pow
		num >>= 1
		pow --
	}
	return result
}