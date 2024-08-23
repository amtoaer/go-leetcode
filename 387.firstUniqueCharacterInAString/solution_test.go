package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=387 lang=golang
 *
 * [387] First Unique Character in a String
 */

// @lc code=start
func firstUniqChar(s string) int {
	var cnt [26]int
	for i := 0; i < len(s); i++ {
		cnt[s[i]-'a']++
	}
	for i := 0; i < len(s); i++ {
		if cnt[s[i]-'a'] == 1 {
			return i
		}
	}
	return -1
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		input  string
		output int
	}{
		{"leetcode", 0},
		{"loveleetcode", 2},
		{"aabb", -1},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.output, firstUniqChar(tt.input))
	}
}
