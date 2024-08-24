package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=264 lang=golang
 *
 * [264] Ugly Number II
 */

// @lc code=start
func nthUglyNumber(n int) int {
	dp := make([]int, n)
	dp[0] = 1
	var p2, p3, p5 int
	for i := 1; i < n; i++ {
		vp2, vp3, vp5 := 2*dp[p2], 3*dp[p3], 5*dp[p5]
		v := min(vp2, vp3, vp5)
		dp[i] = v
		if v == vp2 {
			p2++
		}
		if v == vp3 {
			p3++
		}
		if v == vp5 {
			p5++
		}
	}
	return dp[n-1]
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		input  int
		output int
	}{
		{10, 12},
		{1, 1},
		{2, 2},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.output, nthUglyNumber(tt.input))
	}
}
