package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=522 lang=golang
 *
 * [522] Longest Uncommon Subsequence II
 */

// @lc code=start
func isSubSequence(a, b string) bool {
	left, right := 0, 0
	for left < len(a) && right < len(b) {
		if a[left] == b[right] {
			left++
		}
		right++
	}
	return left == len(a)
}

func findLUSlength(strs []string) int {
	res := -1
label:
	for i := 0; i < len(strs); i++ {
		for j := 0; j < len(strs); j++ {
			if i != j && isSubSequence(strs[i], strs[j]) {
				continue label
			}
		}
		if res < len(strs[i]) {
			res = len(strs[i])
		}
	}
	return res
}

// @lc code=end

func Test(t *testing.T) {
	testcases := []struct {
		strs []string
		want int
	}{
		{strs: []string{"aba", "cdc", "eae"}, want: 3},
		{strs: []string{"aaa", "aaa", "aa"}, want: -1},
	}
	for _, tc := range testcases {
		assert.Equal(t, tc.want, findLUSlength(tc.strs))
	}
}
