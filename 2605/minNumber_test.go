package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	tests := []struct {
		nums1  []int
		nums2  []int
		target int
	}{
		{[]int{1, 2, 3}, []int{1, 2, 3}, 1},
		{[]int{1, 2, 3}, []int{4, 5, 3}, 3},
		{[]int{1, 2, 3}, []int{4, 5, 6}, 14},
		{[]int{7, 8, 9}, []int{1, 2, 3}, 17},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8}, []int{9}, 19},
	}
	for _, tt := range tests {
		assert.Equal(t, minNumber(tt.nums1, tt.nums2), tt.target)
	}
}
