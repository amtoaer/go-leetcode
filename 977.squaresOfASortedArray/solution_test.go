package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=977 lang=golang
 *
 * [977] Squares of a Sorted Array
 */

// @lc code=start
func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func sortedSquares(nums []int) []int {
	res := make([]int, len(nums))
	left, right := 0, len(nums)-1
	cur := len(nums) - 1
	for left <= right {
		if abs(nums[left]) >= abs(nums[right]) {
			res[cur] = nums[left] * nums[left]
			left++
		} else {
			res[cur] = nums[right] * nums[right]
			right--
		}
		cur--
	}
	return res
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		input  []int
		output []int
	}{
		{input: []int{-4, -1, 0, 3, 10}, output: []int{0, 1, 9, 16, 100}},
		{input: []int{-7, -3, 2, 3, 11}, output: []int{4, 9, 9, 49, 121}},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.output, sortedSquares(tt.input))
	}
}
