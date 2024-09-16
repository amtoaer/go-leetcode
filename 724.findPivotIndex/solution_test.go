package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=724 lang=golang
 *
 * [724] Find Pivot Index
 */

// @lc code=start
func pivotIndex(nums []int) int {
	leftSum, rightSum := make([]int, len(nums)), make([]int, len(nums))
	var leftVal, rightVal int
	for i := 0; i < len(nums); i++ {
		leftSum[i] = leftVal
		leftVal += nums[i]
	}
	for i := len(nums) - 1; i >= 0; i-- {
		rightSum[i] = rightVal
		rightVal += nums[i]
	}
	for i := 0; i < len(nums); i++ {
		if leftSum[i] == rightSum[i] {
			return i
		}
	}
	return -1
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		nums []int
		out  int
	}{
		{[]int{1, 7, 3, 6, 5, 6}, 3},
		{[]int{1, 2, 3}, -1},
		{[]int{2, 1, -1}, 0},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.out, pivotIndex(tt.nums))
	}
}
