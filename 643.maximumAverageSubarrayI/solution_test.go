package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=643 lang=golang
 *
 * [643] Maximum Average Subarray I
 */

// @lc code=start
func findMaxAverage(nums []int, k int) float64 {
	if len(nums) < k {
		return 0
	}
	var (
		left, right = 0, k - 1
		sum, maxSum = 0, 0
	)
	for i := left; i <= right; i++ {
		sum += nums[i]
	}
	maxSum = sum
	for right+1 < len(nums) {
		sum = sum + nums[right+1] - nums[left]
		maxSum = max(sum, maxSum)
		left++
		right++
	}
	return float64(maxSum) / float64(k)
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		nums []int
		k    int
		want float64
	}{
		{nums: []int{1, 12, -5, -6, 50, 3}, k: 4, want: 12.75},
		{nums: []int{5}, k: 1, want: 5},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.want, findMaxAverage(tt.nums, tt.k))
	}
}
