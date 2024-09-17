package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=2390 lang=golang
 *
 * [2390] Removing Stars From a String
 */

// @lc code=start
func removeStars(s string) string {
	var stack []rune
	for _, r := range s {
		if r == '*' {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, r)
		}
	}
	return string(stack)
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		s    string
		want string
	}{
		{"leet**cod*e", "lecoe"},
		{"erase*****", ""},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.want, removeStars(tt.s))
	}
}
