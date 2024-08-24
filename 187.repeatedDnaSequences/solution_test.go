package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=187 lang=golang
 *
 * [187] Repeated DNA Sequences
 */

// @lc code=start
func findRepeatedDnaSequences(s string) []string {
	res := []string{}
	m := make(map[string]int)
	if len(s) <= 10 {
		return res
	}
	left, right := 0, 0
	for right-left < 10 && right < len(s) {
		right++
		if right-left == 10 {
			m[s[left:right]]++
			left++
		}
	}
	for k, v := range m {
		if v > 1 {
			res = append(res, k)
		}
	}
	return res
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		input  string
		output []string
	}{
		{"AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT", []string{"AAAAACCCCC", "CCCCCAAAAA"}},
		{"AAAAAAAAAAAAA", []string{"AAAAAAAAAA"}},
	}
	for _, tt := range tc {
		assert.ElementsMatch(t, tt.output, findRepeatedDnaSequences(tt.input))
	}
}
