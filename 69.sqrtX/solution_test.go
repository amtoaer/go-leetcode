package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=69 lang=golang
 *
 * [69] Sqrt(x)
 */

// @lc code=start
func mySqrt(x int) int {
	left, right := 0, x
	ans := -1
	for left <= right {
		mid := left + (right-left)/2
		if mid*mid <= x {
			ans = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return ans
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		input  int
		output int
	}{
		{4, 2},
		{8, 2},
		{9, 3},
		{10, 3},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.output, mySqrt(tt.input))
	}
}
