package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=1493 lang=golang
 *
 * [1493] Longest Subarray of 1's After Deleting One Element
 */

// @lc code=start
func longestSubarray(nums []int) int {
	var (
		left, res int
		deleted   int
	)
	for right := 0; right < len(nums); right++ {
		if nums[right] != 1 {
			deleted += 1
		}
		for deleted > 1 {
			if nums[left] != 1 {
				deleted--
			}
			left++
		}
		res = max(res, right-left)
	}
	return res
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		input  []int
		output int
	}{
		{[]int{1, 1, 0, 1}, 3},
		{[]int{0, 1, 1, 1, 0, 1, 1, 0, 1}, 5},
		{[]int{1, 1, 1}, 2},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.output, longestSubarray(tt.input))
	}
}
