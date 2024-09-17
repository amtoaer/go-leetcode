package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=394 lang=golang
 *
 * [394] Decode String
 */

// @lc code=start
func decodeString(s string) string {
	var (
		stack []byte
		ints  []int
		pos   int
	)
	for pos < len(s) {
		cur := s[pos]
		if cur >= '0' && cur <= '9' {
			val := 0
			for s[pos] >= '0' && s[pos] <= '9' {
				val = val*10 + int(s[pos]-'0')
				pos++
			}
			ints = append(ints, val)
		} else if (cur >= 'a' && cur <= 'z') || (cur >= 'A' && cur <= 'Z') || cur == '[' {
			stack = append(stack, cur)
			pos++
		} else {
			// 遇到了 ]
			i := len(stack) - 1
			for stack[i] != '[' {
				i--
			}
			// IMPORTANT: 需要重复的部分引用了 stack 的后面一部分，需要深拷贝，不然 append 过程中会被覆盖
			needRepeated := append([]byte{}, stack[i+1:]...)
			stack = stack[:i]
			val := ints[len(ints)-1]
			ints = ints[:len(ints)-1]
			for j := 0; j < val; j++ {
				stack = append(stack, needRepeated...)
			}
			pos++
		}
	}
	return string(stack)
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		input  string
		output string
	}{
		{"3[a]2[bc]", "aaabcbc"},
		{"3[a2[c]]", "accaccacc"},
		{"2[abc]3[cd]ef", "abcabccdcdcdef"},
		{"abc3[cd]xyz", "abccdcdcdxyz"},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.output, decodeString(tt.input))
	}
}
