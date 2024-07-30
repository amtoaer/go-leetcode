package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=45 lang=golang
 *
 * [45] Jump Game II
 */

// @lc code=start
func jump(nums []int) int {
	var step, end, tmpEnd int
	for idx, num := range nums[:len(nums)-1] {
		tmpEnd = max(tmpEnd, idx+num)
		if idx == end {
			end = tmpEnd
			step++
		}
	}
	return step
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		input  []int
		output int
	}{
		{input: []int{2, 3, 1, 1, 4}, output: 2},
		{input: []int{2, 3, 0, 1, 4}, output: 2},
		{input: []int{0}, output: 0},
		{input: []int{1}, output: 0},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.output, jump(tt.input))
	}
}
