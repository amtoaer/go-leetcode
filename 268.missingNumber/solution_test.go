package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=268 lang=golang
 *
 * [268] Missing Number
 */

// @lc code=start
func missingNumber(nums []int) int {
	var res int
	for _, v := range nums {
		res ^= v
	}
	for i := 0; i <= len(nums); i++ {
		res ^= i
	}
	return res
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		input  []int
		output int
	}{
		{input: []int{3, 0, 1}, output: 2},
		{input: []int{0, 1}, output: 2},
		{input: []int{9, 6, 4, 2, 3, 5, 7, 0, 1}, output: 8},
		{input: []int{0}, output: 1},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.output, missingNumber(tt.input))
	}
}
