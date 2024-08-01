package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=290 lang=golang
 *
 * [290] Word Pattern
 */

// @lc code=start
func wordPattern(pattern string, s string) bool {
	left, right := 0, 0
	m := make(map[byte]string)
	reverseM := make(map[string]byte)
	for left < len(pattern) && right < len(s) {
		tmp := right
		for tmp < len(s) && s[tmp] != ' ' {
			tmp++
		}
		if w, ok := m[pattern[left]]; ok {
			if w != s[right:tmp] {
				return false
			}
		} else {
			if b, ok := reverseM[s[right:tmp]]; ok && b != pattern[left] {
				return false
			}
			m[pattern[left]] = s[right:tmp]
			reverseM[s[right:tmp]] = pattern[left]
		}
		left++
		right = tmp + 1
	}
	return left >= len(pattern) && right >= len(s)
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		pattern string
		s       string
		want    bool
	}{
		{"abba", "dog cat cat dog", true},
		{"abba", "dog cat cat fish", false},
		{"aaaa", "dog cat cat dog", false},
		{"abba", "dog dog dog dog", false},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.want, wordPattern(tt.pattern, tt.s))
	}
}
