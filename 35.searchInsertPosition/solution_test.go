package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=35 lang=golang
 *
 * [35] Search Insert Position
 */

// @lc code=start
func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1
	var mid int
	for left <= right {
		mid = left + (right-left)>>1
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}

// @lc code=end

func Test(t *testing.T) {
	assert.Equal(t, 2, searchInsert([]int{1, 3, 5, 6}, 5))
	assert.Equal(t, 1, searchInsert([]int{1, 3, 5, 6}, 2))
	assert.Equal(t, 4, searchInsert([]int{1, 3, 5, 6}, 7))
}
