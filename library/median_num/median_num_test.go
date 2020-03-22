package median_num

import (
	`testing`

	`github.com/stretchr/testify/assert`
)

func TestGetMedianNum(t *testing.T) {
	var(
		list1 = []int{3,1,0,1}
		list2 = []int{2,2,3,1}
	)
	medianNum := GetMedianNum(list1, list2)
	assert.Equal(t,1.5,medianNum)
}
