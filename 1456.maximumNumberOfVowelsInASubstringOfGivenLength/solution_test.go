package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=1456 lang=golang
 *
 * [1456] Maximum Number of Vowels in a Substring of Given Length
 */

// @lc code=start

var m = map[byte]struct{}{
	'a': {},
	'e': {},
	'i': {},
	'o': {},
	'u': {},
}

func isVowel(b byte) bool {
	_, ok := m[b]
	return ok
}

func maxVowels(s string, k int) int {
	var (
		left     int
		cnt, res int
	)
	for right := 0; right < len(s); right++ {
		if isVowel(s[right]) {
			cnt++
		}
		for right-left+1 > k {
			if isVowel(s[left]) {
				cnt--
			}
			left++
		}
		res = max(res, cnt)
	}
	return res
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		s   string
		k   int
		out int
	}{
		{"abciiidef", 3, 3},
		{"aeiou", 2, 2},
		{"leetcode", 3, 2},
		{"rhythms", 4, 0},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.out, maxVowels(tt.s, tt.k))
	}
}
