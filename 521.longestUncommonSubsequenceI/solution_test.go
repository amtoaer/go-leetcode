package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=521 lang=golang
 *
 * [521] Longest Uncommon Subsequence I
 */

// @lc code=start
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func findLUSlength(a string, b string) int {
	if a == b {
		return -1
	}
	return max(len(a), len(b))
}

// @lc code=end

func Test(t *testing.T) {
	testcases := []struct {
		a        string
		b        string
		expected int
	}{
		{"aba", "cdc", 3},
		{"aaa", "bbb", 3},
		{"aaa", "aaa", -1},
	}
	for _, tc := range testcases {
		assert.Equal(t, tc.expected, findLUSlength(tc.a, tc.b))
	}
}
