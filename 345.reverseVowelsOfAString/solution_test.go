package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=345 lang=golang
 *
 * [345] Reverse Vowels of a String
 */

// @lc code=start

func should_handle(b byte) bool {
	valid := []byte{'a', 'e', 'i', 'o', 'u'}
	for _, item := range valid {
		if item == b || item-byte(32) == b {
			return true
		}
	}
	return false
}

func reverseVowels(s string) string {
	res := []byte(s)
	left, right := 0, len(s)-1
	for {
		for left < right && !should_handle(s[left]) {
			left++
		}
		for left < right && !should_handle(s[right]) {
			right--
		}
		if left >= right {
			return string(res)
		}
		res[left], res[right] = res[right], res[left]
		left++
		right--
	}
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		input  string
		output string
	}{
		{"IceCreAm", "AceCreIm"},
		{"leetcode", "leotcede"},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.output, reverseVowels(tt.input))
	}
}
