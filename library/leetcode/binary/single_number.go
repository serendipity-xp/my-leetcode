package binary

/**
* @DateTime   : 2020/11/26
* @Author     : xumamba
* @Description:
**/

// singleNumber 给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
func singleNumber(nums []int) int {
	result := 0
	for _, num := range nums {
		result = result ^ num
	}
	return result
}

// singleNumber2 给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现了三次。找出那个只出现了一次的元素。
func singleNumber2(nums []int) int {
	var result int
	for i := 0; i < 64; i++ {
		sum := 0
		for j := 0; j < len(nums); j++ {
			// 统计末位1的个数
			sum += (nums[j] >> i) & 1
		}
		// 还原位 00^10 == 10
		result ^= (sum % 3) << i
	}
	return result
}

// singleNumber3 给定一个整数数组 nums，其中恰好有两个元素只出现一次，其余所有元素均出现两次。 找出只出现一次的那两个元素。
func singleNumber3(nums []int) []int {
	diff := 0 // 两个不同元素的异或值
	for i := 0; i < len(nums); i++ {
		diff ^= nums[i]
	}
	result := []int{diff, diff}
	// 计算两个不同值 最后一个1的位置
	diff = (diff & (diff - 1)) ^ diff
	for i := 0; i < len(nums); i++ {
		if diff&nums[i] == 0 {
			// nums 最后一个1 其中一个结果值一样
			// diff = a ^ b
			// a = a ^ b ^ b
			result[0] ^= nums[i]
		} else {
			// diff = a ^ b
			// b = a ^ b ^ b
			result[1] ^= nums[i]
		}
	}
	return result
}