package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=300 lang=golang
 *
 * [300] Longest Increasing Subsequence
 */

// @lc code=start
func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums))
	dp[0] = 1
	res := 1
	for i := 1; i < len(nums); i++ {
		tmp := 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				tmp = max(tmp, dp[j]+1)
			}
		}
		res = max(res, tmp)
		dp[i] = tmp
	}
	return res
}

// @lc code=end

func Test(t *testing.T) {
	type TestCase struct {
		nums []int
		want int
	}
	testCases := []TestCase{
		{nums: []int{10, 9, 2, 5, 3, 7, 101, 18}, want: 4},
		{nums: []int{0, 1, 0, 3, 2, 3}, want: 4},
		{nums: []int{7, 7, 7, 7, 7, 7, 7}, want: 1},
		{nums: []int{1, 3, 6, 7, 9, 4, 10, 5, 6}, want: 6},
	}
	for _, tc := range testCases {
		assert.Equal(t, tc.want, lengthOfLIS(tc.nums))
	}
}
