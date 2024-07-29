package main

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=2740 lang=golang
 *
 * [2740] Find the Value of the Partition
 */

// @lc code=start
func findValueOfPartition(nums []int) int {
	res := math.MaxInt
	for i := 1; i < len(nums); i++ {
		res = min(res, nums[i]-nums[i-1])
	}
	return res
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		input  []int
		output int
	}{
		{[]int{3, 6, 9, 15}, 3},
		{[]int{1, 3, 2, 4}, 1},
		{[]int{100, 1, 10}, 9},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.output, findValueOfPartition(tt.input))
	}
}
