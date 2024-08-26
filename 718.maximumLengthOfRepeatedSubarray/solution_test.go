package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=718 lang=golang
 *
 * [718] Maximum Length of Repeated Subarray
 */

// @lc code=start
func findLength(nums1 []int, nums2 []int) int {
	dp := make([][]int, len(nums1)+1)
	for i := range dp {
		dp[i] = make([]int, len(nums2)+1)
	}
	var res int
	for i := 1; i <= len(nums1); i++ {
		for j := 1; j <= len(nums2); j++ {
			if nums1[i-1] != nums2[j-1] {
				dp[i][j] = 0
			} else {
				dp[i][j] = dp[i-1][j-1] + 1
				res = max(res, dp[i][j])
			}
		}
	}
	return res
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		nums1 []int
		nums2 []int
		want  int
	}{
		{[]int{1, 2, 3, 2, 1}, []int{3, 2, 1, 4, 7}, 3},
		{[]int{0, 0, 0, 0, 0}, []int{0, 0, 0, 0, 0}, 5},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.want, findLength(tt.nums1, tt.nums2))
	}
}
