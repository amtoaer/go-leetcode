package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=242 lang=golang
 *
 * [242] Valid Anagram
 */

// @lc code=start
func isAnagram(s string, t string) bool {
	m := make([]int, 26)
	for i := 0; i < len(s); i++ {
		m[int(s[i]-'a')]++
	}
	for j := 0; j < len(t); j++ {
		m[int(t[j]-'a')]--
	}
	for _, v := range m {
		if v != 0 {
			return false
		}
	}
	return true
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		s    string
		t    string
		want bool
	}{
		{"anagram", "nagaram", true},
		{"rat", "car", false},
		{"a", "ab", false},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.want, isAnagram(tt.s, tt.t))
	}
}
