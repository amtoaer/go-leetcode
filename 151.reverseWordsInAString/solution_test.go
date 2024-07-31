package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=151 lang=golang
 *
 * [151] Reverse Words in a String
 */

// @lc code=start
func reverseWords(s string) string {
	var label int
	var sb strings.Builder
	sBytes := []byte(s)
	for i := len(sBytes) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			continue
		}
		if i == len(s)-1 || s[i+1] == ' ' {
			label = i
		}
		if i == 0 || s[i-1] == ' ' {
			sb.Write(sBytes[i : label+1])
			sb.WriteByte(' ')
		}
	}
	res := sb.String()
	return res[:len(res)-1]
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		input  string
		output string
	}{
		{"the sky is blue", "blue is sky the"},
		{"  hello world  ", "world hello"},
		{"a good   example", "example good a"},
		{"  Bob    Loves  Alice   ", "Alice Loves Bob"},
		{"Alice does not even like bob", "bob like even not does Alice"},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.output, reverseWords(tt.input))
	}
}
