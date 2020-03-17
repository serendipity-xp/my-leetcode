package two_sum

func two_sum(nums []int, target int) []int {
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
