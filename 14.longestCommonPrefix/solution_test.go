package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=14 lang=golang
 *
 * [14] Longest Common Prefix
 */

// @lc code=start
func longestCommonPrefix(strs []string) string {
	var sb strings.Builder
loop:
	for i := 0; i < len(strs[0]); i++ {
		for idx := range strs {
			if i >= len(strs[idx]) {
				break loop
			}
			if strs[idx][i] != strs[0][i] {
				break loop
			}
		}
		sb.WriteByte(strs[0][i])
	}
	return sb.String()
}

// @lc code=end

func Test(t *testing.T) {
	assert.Equal(t, "fl", longestCommonPrefix([]string{"flower", "flow", "flight"}))
	assert.Equal(t, "", longestCommonPrefix([]string{"dog", "racecar", "car"}))
}
