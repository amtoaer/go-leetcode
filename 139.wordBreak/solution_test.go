package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=139 lang=golang
 *
 * [139] Word Break
 */

// @lc code=start
func wordBreak(s string, wordDict []string) bool {
	m := make(map[string]struct{})
	for _, word := range wordDict {
		m[word] = struct{}{}
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if !dp[j] {
				continue
			}
			if _, ok := m[s[j:i]]; ok {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		s        string
		wordDict []string
		want     bool
	}{
		{"leetcode", []string{"leet", "code"}, true},
		{"applepenapple", []string{"apple", "pen"}, true},
		{"catsandog", []string{"cats", "dog", "sand", "and", "cat"}, false},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.want, wordBreak(tt.s, tt.wordDict))
	}
}
