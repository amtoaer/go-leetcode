package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=191 lang=golang
 *
 * [191] Number of 1 Bits
 */

// @lc code=start
func hammingWeight(n int) int {
	cnt := 0
	for n > 0 {
		cnt += n & 1
		n >>= 1
	}
	return cnt
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		input  int
		output int
	}{
		{11, 3},
		{128, 1},
		{2147483645, 30},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.output, hammingWeight(tt.input))
	}
}
