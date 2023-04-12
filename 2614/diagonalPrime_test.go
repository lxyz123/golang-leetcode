package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	tests := []struct {
		nums   [][]int
		target int
	}{
		{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, 7},
		{[][]int{{1, 2, 3}, {5, 6, 7}, {9, 10, 11}}, 11},
	}
	for _, tt := range tests {
		assert.Equal(t, diagonalPrime(tt.nums), tt.target)
	}
}
