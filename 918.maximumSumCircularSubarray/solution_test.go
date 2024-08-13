package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=918 lang=golang
 *
 * [918] Maximum Sum Circular Subarray
 */

// @lc code=start
func maxSubarraySumCircular(nums []int) int {
	dpMax, dpMin := nums[0], nums[0]
	maxRes, minRes := nums[0], nums[0]
	sum := nums[0]
	for i := 1; i < len(nums); i++ {
		dpMax = max(nums[i], dpMax+nums[i])
		maxRes = max(maxRes, dpMax)
		dpMin = min(nums[i], dpMin+nums[i])
		minRes = min(minRes, dpMin)
		sum += nums[i]
	}
	if minRes == sum {
		return maxRes
	}
	return max(maxRes, sum-minRes)
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		input  []int
		output int
	}{
		{[]int{1, -2, 3, -2}, 3},
		{[]int{5, -3, 5}, 10},
		{[]int{3, -1, 2, -1}, 4},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.output, maxSubarraySumCircular(tt.input))
	}
}
