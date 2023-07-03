package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] Valid Parentheses
 */

// @lc code=start
func isValid(s string) bool {
	var bytes []byte
	pairs := map[byte]byte{
		')': '(',
		'}': '{',
		']': '[',
	}
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '(', '{', '[':
			bytes = append(bytes, s[i])
		case ')', '}', ']':
			if len(bytes) == 0 || bytes[len(bytes)-1] != pairs[s[i]] {
				return false
			}
			bytes = bytes[:len(bytes)-1]
		}
	}
	return len(bytes) == 0
}

// @lc code=end

func Test(t *testing.T) {
	assert.Equal(t, true, isValid("()"))
	assert.Equal(t, true, isValid("()[]{}"))
	assert.Equal(t, false, isValid("(]"))
	assert.Equal(t, false, isValid("([)]"))
	assert.Equal(t, true, isValid("{[]}"))
}
