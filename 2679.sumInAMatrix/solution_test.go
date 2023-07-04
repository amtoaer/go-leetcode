package main

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=2679 lang=golang
 *
 * [2679] Sum in a Matrix
 */

// @lc code=start
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func matrixSum(nums [][]int) int {
	if len(nums) == 0 || len(nums[0]) == 0 {
		return 0
	}
	for i := 0; i < len(nums); i++ {
		sort.Ints(nums[i])
	}
	sum := 0
	for j := 0; j < len(nums[0]); j++ {
		maxVal := nums[0][j]
		for i := 1; i < len(nums); i++ {
			maxVal = max(maxVal, nums[i][j])
		}
		sum += maxVal
	}
	return sum
}

// @lc code=end

func Test(t *testing.T) {
	assert.Equal(t, 12, matrixSum([][]int{{1, 2, 3}, {3, 4, 5}}))
	assert.Equal(t, 21, matrixSum([][]int{{1, 2, 3}, {3, 4, 5}, {6, 7, 8}}))
}
