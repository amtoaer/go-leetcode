package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=72 lang=golang
 *
 * [72] Edit Distance
 */

// @lc code=start
func minDistance(word1 string, word2 string) int {
	dp := make([][]int, len(word1)+1)
	for i := range dp {
		dp[i] = make([]int, len(word2)+1)
	}
	for i := 0; i <= len(word1); i++ {
		for j := 0; j <= len(word2); j++ {
			if i == 0 {
				dp[i][j] = j
			} else if j == 0 {
				dp[i][j] = i
			} else {
				pre := dp[i-1][j-1]
				if word1[i-1] != word2[j-1] {
					pre += 1
				}
				dp[i][j] = min(dp[i-1][j]+1, dp[i][j-1]+1, pre)
			}
		}
	}
	return dp[len(word1)][len(word2)]
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		word1, word2 string
		expect       int
	}{
		{"horse", "ros", 3},
		{"intention", "execution", 5},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.expect, minDistance(tt.word1, tt.word2))
	}
}
