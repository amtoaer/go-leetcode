package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=34 lang=golang
 *
 * [34] Find First and Last Position of Element in Sorted Array
 */

// @lc code=start
func searchRange(nums []int, target int) []int {
	binarySearch := func(left, right int, findLeft bool) int {
		res := len(nums)
		for left <= right {
			mid := left + (right-left)>>1
			if nums[mid] > target || (nums[mid] == target && findLeft) {
				right = mid - 1
				res = mid
			} else {
				left = mid + 1
			}
		}
		return res
	}
	left, right := binarySearch(0, len(nums)-1, true), binarySearch(0, len(nums)-1, false)-1
	if left <= right && right < len(nums) && nums[left] == target && nums[right] == target {
		return []int{left, right}
	}
	return []int{-1, -1}
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		input  []int
		target int
		output []int
	}{
		{[]int{5, 7, 7, 8, 8, 10}, 8, []int{3, 4}},
		{[]int{5, 7, 7, 8, 8, 10}, 6, []int{-1, -1}},
		{[]int{}, 0, []int{-1, -1}},
		{[]int{1}, 1, []int{0, 0}},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.output, searchRange(tt.input, tt.target))
	}
}
