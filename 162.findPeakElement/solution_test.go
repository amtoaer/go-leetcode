package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=162 lang=golang
 *
 * [162] Find Peak Element
 */

// @lc code=start
func findPeakElement(nums []int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)>>1
		var gtLeft, gtRight bool
		if mid == 0 || nums[mid] > nums[mid-1] {
			gtLeft = true
		}
		if mid == len(nums)-1 || nums[mid] > nums[mid+1] {
			gtRight = true
		}
		if gtLeft && gtRight {
			return mid
		}
		if gtLeft {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return 0
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		input  []int
		output int
	}{
		{[]int{1, 2, 3, 1}, 2},
		{[]int{1, 2, 1, 3, 5, 6, 4}, 5},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.output, findPeakElement(tt.input))
	}
}
