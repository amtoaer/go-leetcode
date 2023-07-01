package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] Longest Palindromic Substring
 */

// @lc code=start
func expandAroundCenter(s string, left, right int) string {
	for left >= 0 && right < len(s) {
		if s[left] == s[right] {
			left--
			right++
			continue
		}
		break
	}
	return s[left+1 : right]
}

func longestPalindrome(s string) string {
	var (
		maxLength int
		result    string
	)
	for i := 0; i < len(s); i++ {
		odd := expandAroundCenter(s, i, i)
		even := expandAroundCenter(s, i, i+1)
		if len(odd) > maxLength {
			maxLength = len(odd)
			result = odd
		}
		if len(even) > maxLength {
			maxLength = len(even)
			result = even
		}
	}
	return result
}

// @lc code=end

func Test(t *testing.T) {
	assert.Equal(t, "bab", longestPalindrome("babad"))
	assert.Equal(t, "bb", longestPalindrome("cbbd"))
}
