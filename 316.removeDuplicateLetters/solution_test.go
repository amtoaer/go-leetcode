package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=316 lang=golang
 *
 * [316] Remove Duplicate Letters
 */

// @lc code=start
func removeDuplicateLetters(s string) string {
	left := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		left[s[i]]++
	}
	var stack []byte
	used := make(map[byte]bool)
	for i := 0; i < len(s); i++ {
		b := s[i]
		// 字母已经被用过了，直接跳过
		if used[b] {
			left[b]--
			continue
		}
		for len(stack) > 0 && // 栈不为空（显而易见）
			stack[len(stack)-1] >= b && // 栈顶元素大于当前元素（否则没必要出栈）
			left[stack[len(stack)-1]] > 0 { // 栈顶元素后面还有（如果后面没有了，这里弹了结果会少一个字母）
			used[stack[len(stack)-1]] = false
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, b)
		used[b] = true
		left[b]--
	}
	return string(stack)
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		input  string
		output string
	}{
		{"bcabc", "abc"},
		{"cbacdcbc", "acdb"},
		{"abacb", "abc"},
		{"cdadabcc", "adbc"},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.output, removeDuplicateLetters(tt.input))
	}
}
