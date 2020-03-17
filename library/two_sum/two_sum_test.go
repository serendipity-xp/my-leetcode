package two_sum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	nums   = []int{2, 7, 11, 15}
	target = 9
	expect = []int{1, 0}
)

func TestTwoSum(t *testing.T) {
	result := two_sum(nums, target)
	assert.Equal(t, result, expect)
}

func BenchmarkTwoSum(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		two_sum(nums, target)
	}
}

// BenchmarkTwoSum-4   	25894044	        47.8 ns/op	      16 B/op	       1 allocs/op
