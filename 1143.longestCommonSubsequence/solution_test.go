package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=1143 lang=golang
 *
 * [1143] Longest Common Subsequence
 */

// @lc code=start
func longestCommonSubsequence(text1 string, text2 string) int {
	dp := make([][]int, len(text1)+1)
	for i := range dp {
		dp[i] = make([]int, len(text2)+1)
	}
	for i := 1; i <= len(text1); i++ {
		for j := 1; j <= len(text2); j++ {
			dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			if text1[i-1] == text2[j-1] {
				dp[i][j] = max(dp[i][j], dp[i-1][j-1]+1)
			}
		}
	}
	return dp[len(text1)][len(text2)]
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		text1  string
		text2  string
		output int
	}{
		{"abcde", "ace", 3},
		{"abc", "abc", 3},
		{"abc", "def", 0},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.output, longestCommonSubsequence(tt.text1, tt.text2))
	}
}
