package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=198 lang=golang
 *
 * [198] House Robber
 */

// @lc code=start
func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	left, right := nums[0], max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		left, right = right, max(right, left+nums[i])
	}
	return right
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		input  []int
		output int
	}{
		{input: []int{1, 2, 3, 1}, output: 4},
		{input: []int{2, 7, 9, 3, 1}, output: 12},
		{input: []int{2, 1, 1, 2}, output: 4},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.output, rob(tt.input))
	}
}
