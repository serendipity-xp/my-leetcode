package two_sum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	nums   = []int{2, 7, 11, 15}
	target = 9
	expect = []int{0, 1}
)

func TestTwoSum(t *testing.T) {
	t.Run("for", func(t *testing.T) {
		result := TwoSumByFor(nums, target)
		assert.Equal(t, result, expect)
	})
	t.Run("hash", func(t *testing.T) {
		result := TwoSumByHash(nums, target)
		assert.Equal(t, result, expect)
	})
	t.Run("hashOnce", func(t *testing.T) {
		result := TwoSumByHash(nums, target)
		assert.Equal(t, result, expect)
	})
}

func BenchmarkTwoSum(b *testing.B) {
	b.Run("for", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			TwoSumByFor(nums, target)
		}
	})
	b.Run("hash", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			TwoSumByHash(nums, target)
		}
	})
	b.Run("hashOnce", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			TwoSumByHashOnce(nums, target)
		}
	})

}

// BenchmarkTwoSum/for-4         	21800482	        51.6 ns/op	      16 B/op	       1 allocs/op
// BenchmarkTwoSum/hash-4        	 7672957	       152 ns/op	      16 B/op	       1 allocs/op
// BenchmarkTwoSum/hashOnce-4    	17749783	        65.2 ns/op	      16 B/op	       1 allocs/op
