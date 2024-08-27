package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=131 lang=golang
 *
 * [131] Palindrome Partitioning
 */

// @lc code=start
func partition(s string) [][]string {
	var res [][]string
	if len(s) == 0 {
		return res
	}
	palindrome := make(map[int]struct{})
	spread := func(i, j int) {
		for i >= 0 && j < len(s) {
			if s[i] != s[j] {
				return
			}
			palindrome[hash(i, j)] = struct{}{}
			i--
			j++
		}
	}
	for i := 0; i < len(s); i++ {
		spread(i, i)
		spread(i, i+1)
	}
	var backtrack func(int, []string)
	backtrack = func(i int, cur []string) {
		if i >= len(s) {
			res = append(res, append([]string{}, cur...))
			return
		}
		for j := i; j < len(s); j++ {
			if _, ok := palindrome[hash(i, j)]; ok {
				cur = append(cur, s[i:j+1])
				backtrack(j+1, cur)
				cur = cur[:len(cur)-1]
			}
		}
	}
	backtrack(0, []string{})
	return res
}

func hash(i, j int) int {
	// i 和 j 都在 16 以内，int 足够表示了
	return i<<10 + j
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		s      string
		expect [][]string
	}{
		{"aab", [][]string{{"aa", "b"}, {"a", "a", "b"}}},
		{"a", [][]string{{"a"}}},
	}
	for _, tt := range tc {
		assert.ElementsMatch(t, partition(tt.s), tt.expect)
	}
}
