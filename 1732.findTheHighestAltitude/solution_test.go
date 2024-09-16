package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=1732 lang=golang
 *
 * [1732] Find the Highest Altitude
 */

// @lc code=start
func largestAltitude(gain []int) int {
	var res, cur int
	for _, g := range gain {
		cur += g
		res = max(res, cur)
	}
	return res
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		gain []int
		out  int
	}{
		{[]int{-5, 1, 5, 0, -7}, 1},
		{[]int{-4, -3, -2, -1, 4, 3, 2}, 0},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.out, largestAltitude(tt.gain))
	}
}
