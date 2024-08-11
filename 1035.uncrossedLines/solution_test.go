package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=1035 lang=golang
 *
 * [1035] Uncrossed Lines
 */

// @lc code=start
func maxUncrossedLines(nums1 []int, nums2 []int) int {
	dp := make([][]int, len(nums1)+1)
	for i := range dp {
		dp[i] = make([]int, len(nums2)+1)
	}
	for i := 1; i <= len(nums1); i++ {
		for j := 1; j <= len(nums2); j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[len(nums1)][len(nums2)]
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		nums1 []int
		nums2 []int
		want  int
	}{
		{[]int{1, 4, 2}, []int{1, 2, 4}, 2},
		{[]int{2, 5, 1, 2, 5}, []int{10, 5, 2, 1, 5, 2}, 3},
		{[]int{1, 3, 7, 1, 7, 5}, []int{1, 9, 2, 5, 1}, 2},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.want, maxUncrossedLines(tt.nums1, tt.nums2))
	}
}
