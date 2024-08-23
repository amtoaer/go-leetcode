package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=1047 lang=golang
 *
 * [1047] Remove All Adjacent Duplicates In String
 */

// @lc code=start
func removeDuplicates(s string) string {
	var stack []byte
	for i := 0; i < len(s); i++ {
		if len(stack) > 0 && stack[len(stack)-1] == s[i] {
			stack = stack[:len(stack)-1]
			continue
		}
		stack = append(stack, s[i])
	}
	return string(stack)
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		input  string
		output string
	}{
		{"abbaca", "ca"},
		{"azxxzy", "ay"},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.output, removeDuplicates(tt.input))
	}
}
