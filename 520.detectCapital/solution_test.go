package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=520 lang=golang
 *
 * [520] Detect Capital
 */

// @lc code=start
func isLower(s string) bool {
	for i := 0; i < len(s); i++ {
		if !(s[i] >= 'a' && s[i] <= 'z') {
			return false
		}
	}
	return true
}

func isUpper(s string) bool {
	for i := 0; i < len(s); i++ {
		if !(s[i] >= 'A' && s[i] <= 'Z') {
			return false
		}
	}
	return true
}

func detectCapitalUse(word string) bool {
	return isUpper(word) || isLower(word) || (isUpper(word[:1]) && isLower(word[1:]))
}

// @lc code=end

func Test(t *testing.T) {
	testcases := []struct {
		in  string
		out bool
	}{
		{in: "USA", out: true},
		{in: "FlaG", out: false},
		{in: "leetcode", out: true},
		{in: "Google", out: true},
		{in: "GooGle", out: false},
	}
	for _, tc := range testcases {
		assert.Equal(t, tc.out, detectCapitalUse(tc.in))
	}
}
