package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	tests := []struct {
		s      string
		length int
	}{
		{"000", 0},
		{"111", 0},
		{"1110", 0},
		{"0001", 2},
		{"010101", 2},
		{"01010011", 4},
		{"00001111000111", 8},
	}
	for _, tt := range tests {
		assert.Equal(t, findTheLongestBalancedSubstring(tt.s), tt.length)
	}
}
