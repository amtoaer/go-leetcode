package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=383 lang=golang
 *
 * [383] Ransom Note
 */

// @lc code=start
func canConstruct(ransomNote string, magazine string) bool {
	m := make([]int, 26)
	for i := 0; i < len(magazine); i++ {
		m[magazine[i]-'a']++
	}
	for i := 0; i < len(ransomNote); i++ {
		m[ransomNote[i]-'a']--
		if m[ransomNote[i]-'a'] < 0 {
			return false
		}
	}
	return true
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		ransomNote string
		magazine   string
		want       bool
	}{
		{"a", "b", false},
		{"aa", "ab", false},
		{"aa", "aab", true},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.want, canConstruct(tt.ransomNote, tt.magazine))
	}
}
