package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=205 lang=golang
 *
 * [205] Isomorphic Strings
 */

// @lc code=start
func isIsomorphic(s string, t string) bool {
	m := make(map[byte]byte)
	reverseM := make(map[byte]byte)
	for i := 0; i < len(s); i++ {
		if x, ok := m[s[i]]; ok {
			if x != t[i] {
				return false
			}
			continue
		}
		if y, ok := reverseM[t[i]]; ok && y != s[i] {
			return false
		}
		m[s[i]] = t[i]
		reverseM[t[i]] = s[i]
	}
	return true
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		s      string
		t      string
		output bool
	}{
		{"egg", "add", true},
		{"foo", "bar", false},
		{"paper", "title", true},
		{"badc", "baba", false},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.output, isIsomorphic(tt.s, tt.t))
	}
}
