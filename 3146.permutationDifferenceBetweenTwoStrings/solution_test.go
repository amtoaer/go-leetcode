package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=3146 lang=golang
 *
 * [3146] Permutation Difference between Two Strings
 */

// @lc code=start
func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}

func findPermutationDifference(s string, t string) int {
	m := make(map[byte]int)
	for i, b := range []byte(s) {
		m[b] = i
	}
	var res int
	for i, b := range []byte(t) {
		res += abs(m[b] - i)
	}
	return res
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		s, t string
		want int
	}{
		{"abc", "abc", 0},
		{"abc", "bac", 2},
		{"abcde", "edbac", 12},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.want, findPermutationDifference(tt.s, tt.t))
	}
}
