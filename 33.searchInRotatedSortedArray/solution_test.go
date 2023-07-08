package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=33 lang=golang
 *
 * [33] Search in Rotated Sorted Array
 */

// @lc code=start
func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	var (
		left, right = 0, len(nums) - 1
		mid         int
	)
	for left <= right {
		mid = left + (right-left)>>1
		if nums[mid] == target {
			return mid
		}
		if nums[0] <= nums[mid] {
			// 左边有序
			if nums[0] <= target && target < nums[mid] {
				// target 在左边
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
		if nums[mid] <= nums[len(nums)-1] {
			// 右边有序
			if nums[mid] < target && target <= nums[len(nums)-1] {
				// target 在右边
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}

// @lc code=end

func Test(t *testing.T) {
	assert.Equal(t, 4, search([]int{4, 5, 6, 7, 0, 1, 2}, 0))
	assert.Equal(t, -1, search([]int{4, 5, 6, 7, 0, 1, 2}, 3))
	assert.Equal(t, -1, search([]int{}, 3))
	assert.Equal(t, 0, search([]int{1}, 1))
}
