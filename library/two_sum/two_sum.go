package two_sum

func TwoSumByFor(nums []int, target int) []int {
	result := make([]int, 2)
	for i, num := range nums {
		aim := target - num
		for j, n := range nums {
			if n == aim && i != j {
				result[0] = i
				result[1] = j
			}
		}
	}
	return result
}

func TwoSumByHash(nums []int, target int) []int {
	numsMap := map[int]int{}
	result := make([]int, 2)
	for i, num := range nums {
		numsMap[num] = i + 1
	}
	for i, num := range nums {
		index := numsMap[target-num]
		if index != 0 && index != i+1 {
			result[0] = index - 1
			result[1] = i
		}
	}
	return result
}

func TwoSumByHashOnce(nums []int, target int) []int {
	numsMap := map[int]int{}
	result := make([]int, 2)
	for i, num := range nums {
		index := numsMap[target-num]
		if index != 0 && index != i+1 {
			result[0] = index - 1
			result[1] = i
			break
		}
		numsMap[num] = i + 1
	}
	return result
}
