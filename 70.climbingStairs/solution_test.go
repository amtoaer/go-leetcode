package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=70 lang=golang
 *
 * [70] Climbing Stairs
 */

// @lc code=start
func climbStairs(n int) int {
	if n == 1 || n == 2 {
		return n
	}
	l, r := 1, 2
	for i := 2; i < n; i++ {
		l, r = r, l+r
	}
	return r
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		input  int
		output int
	}{
		{2, 2},
		{3, 3},
		{4, 5},
		{5, 8},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.output, climbStairs(tt.input))
	}
}
