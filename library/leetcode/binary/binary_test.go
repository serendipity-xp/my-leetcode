package binary

/**
 * @DateTime   : 2020/11/26
 * @Author     : xumamba
 * @Description:
 **/

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingleNumber(t *testing.T) {
	testCase := []struct {
		input  []int
		output int
	}{
		{
			[]int{1, 2, 3, 4, 1, 3, 4},
			2,
		},
		{
			[]int{-1, -2, 0, 4, -1, -2, 4},
			0,
		},
	}
	for _, tc := range testCase {
		number := singleNumber(tc.input)
		assert.Equal(t, number, tc.output)
	}
}

func TestSingleNumber2(t *testing.T) {
	testCase := []struct {
		input  []int
		output int
	}{
		{
			[]int{1, 2, 1, 3, 3, 4, 4, 1, 3, 4},
			2,
		},
		{
			[]int{-1, -1, -2, -2, 0, 4, 4, -1, -2, 4},
			0,
		},
	}
	for _, tc := range testCase {
		number := singleNumber2(tc.input)
		assert.Equal(t, number, tc.output)
	}
}

func TestSingleNumber3(t *testing.T) {
	testCase := []struct {
		input  []int
		output []int
	}{
		{
			[]int{1, 2, 3, 4, 3, 4},
			[]int{1, 2},
		},
		{
			[]int{-1, -2, 0, -1, -2, 4},
			[]int{4, 0},
		},
	}
	for _, tc := range testCase {
		number := singleNumber3(tc.input)
		assert.Equal(t, number, tc.output)
	}
}

func TestHamming(t *testing.T) {
	testCase := []struct {
		input  uint32
		output int
	}{
		{
			3,
			2,
		},
		{
			11,
			3,
		},
	}
	for _, tc := range testCase {
		result := hammingWeight(tc.input)
		assert.Equal(t, result, tc.output)

		result = hammingWeight2(tc.input)
		assert.Equal(t, result, tc.output)
	}
}

func TestCountBits(t *testing.T) {
	testCase := []struct {
		input  int
		output []int
	}{
		{
			0,
			[]int{0},
		},
		{
			5,
			[]int{0, 1, 1, 2, 1, 2},
		},
	}
	for _, tc := range testCase {
		result := countBits(tc.input)
		assert.Equal(t, result, tc.output)

		result = countBits2(tc.input)
		assert.Equal(t, result, tc.output)
	}
}

func TestReserveBits(t *testing.T) {
	testCase := []struct {
		input  uint32
		output uint32
	}{
		{
			0,
			0,
		},
		{
			5,
			0xa0000000,
		},
	}
	for _, tc := range testCase {
		result := reverseBits(tc.input)
		assert.Equal(t, result, tc.output)
	}
}

func TestRangeBitwiseAnd(t *testing.T) {
	res := rangeBitwiseAnd(5, 7)
	fmt.Println(res)
}