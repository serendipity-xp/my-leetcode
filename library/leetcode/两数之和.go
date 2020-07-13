package leetcode

func twoSum(nums []int, target int) []int {
	result := make([]int,2)
	for index1, value1 := range nums{
		// if value1 <= target{
			other := target - value1
			for index2 := index1 + 1; index2 < len(nums); index2++{
				if other == nums[index2]{
					result[0] = index1
					result[1] = index2
				}
			}
		}
	// }
	return result
}

func twoSumHash(nums []int, target int) []int {
	result := make([]int,2)
	hashMap := make(map[int]int)
	for index, value := range nums{
		other := target - value
		if otherIndex, exist := hashMap[other]; exist {
			result[0] = otherIndex
			result[1] = index
			break
		}
		hashMap[value] = index
	}
	return result
}