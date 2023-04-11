package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	tests := []struct {
		digits []string
		n      int
		res    int
	}{
		{[]string{"1", "3", "5", "7"}, 100, 20},
		{[]string{"1", "4", "7"}, 1000000000, 29523},
		{[]string{"7"}, 8, 1},
	}
	for _, tt := range tests {
		assert.Equal(t, atMostNGivenDigitSet(tt.digits, tt.n), tt.res)
	}
}
