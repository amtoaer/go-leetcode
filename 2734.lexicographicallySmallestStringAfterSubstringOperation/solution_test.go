package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=2734 lang=golang
 *
 * [2734] Lexicographically Smallest String After Substring Operation
 */

// @lc code=start
func smallestString(s string) string {
	b := []byte(s)
	// 直接遍历字符串，寻找第一个大于 'a' 的字符
	for i := 0; i < len(b); i++ {
		if b[i] > 'a' {
			// 从找到的字符开始，将后续所有大于 'a' 的字符减一
			for j := i; j < len(b) && b[j] > 'a'; j++ {
				b[j] -= 1
			}
			return string(b)
		}
	}
	// 如果所有字符都不大于 'a'，则将最后一个字符设为 'z'
	b[len(b)-1] = 'z'
	return string(b)
}

// @lc code=end

func Test(t *testing.T) {
	testcases := []struct {
		s    string
		want string
	}{
		{"cbabc", "baabc"},
		{"acbbc", "abaab"},
		{"leetcode", "kddsbncd"},
		{"aa", "az"},
	}
	for _, tc := range testcases {
		assert.Equal(t, tc.want, smallestString(tc.s))
	}
}
