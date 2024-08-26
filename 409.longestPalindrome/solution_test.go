package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=409 lang=golang
 *
 * [409] Longest Palindrome
 */

// @lc code=start
func longestPalindrome(s string) int {
	m := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		m[s[i]]++
	}
	res := 0
	hasOdd := false
	for _, v := range m {
		if v&1 == 0 {
			res += v
		} else if hasOdd {
			res += v - 1
		} else {
			res += v
			hasOdd = true
		}
	}
	return res
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		input  string
		output int
	}{
		{"abccccdd", 7},
		{"a", 1},
		{"bb", 2},
		{"ab", 1},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.output, longestPalindrome(tt.input))
	}
}
