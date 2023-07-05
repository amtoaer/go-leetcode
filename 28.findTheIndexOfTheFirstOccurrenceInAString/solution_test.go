package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=28 lang=golang
 *
 * [28] Find the Index of the First Occurrence in a String
 */

// @lc code=start
func buildNext(s string) []int {
	res := make([]int, len(s))
	res[0] = 0
	i := 1
	now := 0
	for i < len(s) {
		if s[i] == s[now] {
			now++
			res[i] = now
			i++
		} else if now != 0 {
			now = res[now-1]
		} else {
			res[i] = 0
			i++
		}
	}
	return res
}

func strStr(haystack string, needle string) int {
	// kmp 算法（写得不熟练，需要继续练）
	if len(needle) == 0 {
		return 0
	}
	next := buildNext(needle)
	pos := 0
	i := 0
	for i < len(haystack) {
		if needle[pos] == haystack[i] {
			pos++
			i++
		} else if pos != 0 {
			pos = next[pos-1]
		} else {
			i++
		}
		if pos == len(needle) {
			return i - pos
		}
	}
	return -1
}

// @lc code=end

func Test(t *testing.T) {
	assert.Equal(t, 2, strStr("hello", "ll"))
	assert.Equal(t, -1, strStr("aaaaa", "bba"))
	assert.Equal(t, 4, strStr("mississippi", "issip"))
}
