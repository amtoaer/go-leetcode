package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=97 lang=golang
 *
 * [97] Interleaving String
 */

// @lc code=start
func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}
	dp := make([]bool, len(s2)+1)
	dp[0] = true
	for i := 0; i <= len(s1); i++ {
		for j := 0; j <= len(s2); j++ {
			idx := i + j - 1
			if i > 0 {
				dp[j] = dp[j] && s1[i-1] == s3[idx]
			}
			if j > 0 {
				dp[j] = dp[j] || (dp[j-1] && s2[j-1] == s3[idx])
			}
		}
	}
	return dp[len(s2)]
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		s1, s2, s3 string
		expect     bool
	}{
		{"aabcc", "dbbca", "aadbbcbcac", true},
		{"aabcc", "dbbca", "aadbbbaccc", false},
		{"", "", "", true},
		{"a", "", "c", false},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.expect, isInterleave(tt.s1, tt.s2, tt.s3))
	}
}
