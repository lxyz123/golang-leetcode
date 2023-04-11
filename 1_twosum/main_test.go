package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		res    []int
	}{
		{[]int{1, 2, 3, 5}, 5, []int{1, 2}},
		{[]int{1, 2, 3, 5}, 10, nil},
		{[]int{1, 2, 3, 5}, 6, []int{0, 3}},
		{[]int{1, 2, 3, 5}, 7, []int{1, 3}},
		{[]int{1, 2, 3, 5}, 3, []int{0, 1}},
		{[]int{1, 2, 3, 5}, 4, []int{0, 2}},
	}
	for _, tt := range tests {
		assert.Equal(t, twoSum(tt.nums, tt.target), tt.res)
	}
}

func TestTwoSum1(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		res    []int
	}{
		{[]int{1, 2, 3, 5}, 5, []int{1, 2}},
		{[]int{1, 2, 3, 5}, 10, nil},
		{[]int{1, 2, 3, 5}, 6, []int{0, 3}},
		{[]int{1, 2, 3, 5}, 7, []int{1, 3}},
		{[]int{1, 2, 3, 5}, 3, []int{0, 1}},
		{[]int{1, 2, 3, 5}, 4, []int{0, 2}},
	}
	for _, tt := range tests {
		assert.Equal(t, twoSum1(tt.nums, tt.target), tt.res)
	}
}
