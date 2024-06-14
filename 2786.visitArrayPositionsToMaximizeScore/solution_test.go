package main

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=2786 lang=golang
 *
 * [2786] Visit Array Positions to Maximize Score
 */

// @lc code=start
func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func maxScore(nums []int, x int) int64 {
	dp := [2]int{math.MinInt32, math.MinInt32}
	dp[nums[0]%2] = nums[0]
	for i := 1; i < len(nums); i++ {
		pos := nums[i] % 2
		aPos := 1 - pos
		dp[pos] = max(dp[pos]+nums[i], dp[aPos]+nums[i]-x)
	}
	return int64(max(dp[0], dp[1]))
}

// @lc code=end

func Test(t *testing.T) {
	testcases := []struct {
		nums []int
		x    int
		want int64
	}{
		{nums: []int{2, 3, 6, 1, 9, 2}, x: 5, want: 13},
		{nums: []int{2, 4, 6, 8}, x: 3, want: 20},
		{
			nums: []int{
				9,
				58,
				17,
				54,
				91,
				90,
				32,
				6,
				13,
				67,
				24,
				80,
				8,
				56,
				29,
				66,
				85,
				38,
				45,
				13,
				20,
				73,
				16,
				98,
				28,
				56,
				23,
				2,
				47,
				85,
				11,
				97,
				72,
				2,
				28,
				52,
				33,
			},
			x:    90,
			want: 886,
		},
	}
	for _, tc := range testcases {
		assert.Equal(t, tc.want, maxScore(tc.nums, tc.x))
	}
}
