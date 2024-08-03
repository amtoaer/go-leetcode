package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=172 lang=golang
 *
 * [172] Factorial Trailing Zeroes
 */

// @lc code=start
func trailingZeroes(n int) (ans int) {
	for n > 0 {
		n /= 5
		ans += n
	}
	return
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		input  int
		output int
	}{
		{input: 3, output: 0},
		{input: 5, output: 1},
		{input: 0, output: 0},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.output, trailingZeroes(tt.input))
	}
}
