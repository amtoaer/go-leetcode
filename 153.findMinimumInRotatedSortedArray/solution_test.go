package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=153 lang=golang
 *
 * [153] Find Minimum in Rotated Sorted Array
 */

// @lc code=start
func findMin(nums []int) int {
	if nums[0] <= nums[len(nums)-1] {
		return nums[0]
	}
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		value := nums[mid]
		if value <= nums[len(nums)-1] {
			if mid > 0 && value < nums[mid-1] {
				return value
			}
			right = mid - 1
		} else {
			if mid < len(nums)-1 && value > nums[mid+1] {
				return nums[mid+1]
			}
			left = mid + 1
		}
	}
	return 0
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		input []int
		want  int
	}{
		{input: []int{3, 4, 5, 1, 2}, want: 1},
		{input: []int{4, 5, 6, 7, 0, 1, 2}, want: 0},
		{input: []int{11, 13, 15, 17}, want: 11},
		{input: []int{5, 1, 2, 3, 4}, want: 1},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.want, findMin(tt.input))
	}
}
