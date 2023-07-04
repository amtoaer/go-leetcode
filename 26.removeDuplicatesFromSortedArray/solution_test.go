package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=26 lang=golang
 *
 * [26] Remove Duplicates from Sorted Array
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	start := 0
	for idx := 1; idx < len(nums); idx++ {
		if nums[idx] != nums[start] {
			start++
			nums[start] = nums[idx]
		}
	}
	return start + 1
}

// @lc code=end

func Test(t *testing.T) {
	assert.Equal(t, 2, removeDuplicates([]int{1, 1, 2}))
	assert.Equal(t, 5, removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
}
