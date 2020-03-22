package median_num

import (
	"sort"
)

func GetMedianNum(list1, list2 []int) float64 {
	list1 = append(list1, list2...)
	sort.Ints(list1)
	length := len(list1)
	if length%2 == 0 {
		return float64(list1[length/2] + list1[length/2-1]) / 2
	}
	return float64(list1[length/2])
}
