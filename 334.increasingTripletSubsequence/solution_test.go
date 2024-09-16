package main

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=334 lang=golang
 *
 * [334] Increasing Triplet Subsequence
 */

// @lc code=start
func increasingTriplet(nums []int) bool {
	if len(nums) < 3 {
		return false
	}
	leftMin, rightMax := make([]int, len(nums)), make([]int, len(nums))
	minVal, maxVal := math.MaxInt, math.MinInt
	for i := 0; i < len(nums); i++ {
		leftMin[i] = minVal
		minVal = min(minVal, nums[i])
	}
	for i := len(nums) - 1; i >= 0; i-- {
		rightMax[i] = maxVal
		maxVal = max(maxVal, nums[i])
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] > leftMin[i] && nums[i] < rightMax[i] {
			return true
		}
	}
	return false
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		nums []int
		want bool
	}{
		{nums: []int{1, 2, 3, 4, 5}, want: true},
		{nums: []int{5, 4, 3, 2, 1}, want: false},
		{nums: []int{2, 1, 5, 0, 4, 6}, want: true},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.want, increasingTriplet(tt.nums))
	}
}
