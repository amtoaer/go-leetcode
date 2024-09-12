package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=1768 lang=golang
 *
 * [1768] Merge Strings Alternately
 */

// @lc code=start
func mergeAlternately(word1 string, word2 string) string {
	var sb strings.Builder
	var l, r int
	var flag bool
	for l < len(word1) && r < len(word2) {
		if flag {
			sb.WriteByte(word2[r])
			r++
		} else {
			sb.WriteByte(word1[l])
			l++
		}
		flag = !flag
	}
	if l < len(word1) {
		sb.WriteString(word1[l:])
	}
	if r < len(word2) {
		sb.WriteString(word2[r:])
	}
	return sb.String()
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		word1  string
		word2  string
		output string
	}{
		{"abc", "pqr", "apbqcr"},
		{"ab", "pqrs", "apbqrs"},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.output, mergeAlternately(tt.word1, tt.word2))
	}
}
