package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=42 lang=golang
 *
 * [42] Trapping Rain Water
 */

// @lc code=start
func trap(height []int) int {
	var stack []int
	var res int
	for i, h := range height {
		for len(stack) > 0 && height[stack[len(stack)-1]] < h {
			mid := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
			left := stack[len(stack)-1]
			res += (i - left - 1) * (min(h, height[left]) - height[mid])
		}
		stack = append(stack, i)
	}
	return res
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		input  []int
		output int
	}{
		{[]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, 6},
		{[]int{4, 2, 0, 3, 2, 5}, 9},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.output, trap(tt.input))
	}
}
