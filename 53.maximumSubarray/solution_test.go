package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=53 lang=golang
 *
 * [53] Maximum Subarray
 */

// @lc code=start
func maxSubArray(nums []int) int {
	res := nums[0]
	dp := nums[0]
	for i := 1; i < len(nums); i++ {
		if dp < 0 {
			dp = nums[i]
		} else {
			dp = dp + nums[i]
		}
		res = max(res, dp)
	}
	return res
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		nums []int
		want int
	}{
		{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
		{[]int{1}, 1},
		{[]int{5, 4, -1, 7, 8}, 23},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.want, maxSubArray(tt.nums))
	}
}
