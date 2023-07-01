package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] Container With Most Water
 */

// @lc code=start
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxArea(height []int) int {
	var (
		left, right     = 0, len(height) - 1
		area, minHeight = 0, 0
	)
	for left < right {
		minHeight = 0
		if height[left] > height[right] {
			minHeight = height[right]
			right--
		} else {
			minHeight = height[left]
			left++
		}
		area = max(area, minHeight*(right-left+1))
	}
	return area
}

// @lc code=end

func Test(t *testing.T) {
	assert.Equal(t, 49, maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	assert.Equal(t, 1, maxArea([]int{1, 1}))
}
