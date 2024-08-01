package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=392 lang=golang
 *
 * [392] Is Subsequence
 */

// @lc code=start
func isSubsequence(s string, t string) bool {
	var sPtr int
	for tPtr := 0; sPtr < len(s) && tPtr < len(t); tPtr++ {
		if s[sPtr] == t[tPtr] {
			sPtr++
		}
	}
	return sPtr == len(s)
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		s      string
		t      string
		expect bool
	}{
		{"abc", "ahbgdc", true},
		{"axc", "ahbgdc", false},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.expect, isSubsequence(tt.s, tt.t))
	}
}
