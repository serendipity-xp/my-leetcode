package median_num

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMedianNum(t *testing.T) {
	var (
		list1 = []int{3, 1, 0, 1}
		list2 = []int{2, 2, 3, 1}

		list3 []int
		list4 []int

		list5 = []int{1}
	)
	medianNum := GetMedianNum(list1, list2)
	assert.Equal(t, 1.5, medianNum)
	medianNum = GetMedianNum(list3, list4)
	assert.Equal(t, 0.0, medianNum)
	medianNum = GetMedianNum(list4, list5)
	assert.Equal(t, 0.5, medianNum)
}
