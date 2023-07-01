package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] Longest Substring Without Repeating Characters
 */

// @lc code=start

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func lengthOfLongestSubstring(s string) int {
	var (
		m      = make(map[byte]int)
		maxLen = 0
		start  = 0
	)
	for i := 0; i < len(s); i++ {
		// 如果当前扫到的字符已经在子字符串间出现过，那么子字符串应该从出现位置的下个字符开始
		if v, ok := m[s[i]]; ok && v >= start {
			start = v + 1
		}
		// 记录下这个字符最后出现的位置
		m[s[i]] = i
		// 计算一次最大距离
		maxLen = max(maxLen, i-start+1)
	}
	return maxLen
}

// @lc code=end

func Test(t *testing.T) {
	assert.Equal(t, 3, lengthOfLongestSubstring("abcabcbb"))
	assert.Equal(t, 1, lengthOfLongestSubstring("bbbbb"))
}
