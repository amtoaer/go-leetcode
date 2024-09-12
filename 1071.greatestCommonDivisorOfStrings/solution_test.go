package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=1071 lang=golang
 *
 * [1071] Greatest Common Divisor of Strings
 */

// @lc code=start
func gcdOfStrings(str1 string, str2 string) string {
	if str1+str2 != str2+str1 {
		return ""
	}
	var end int
	for end = 0; end < len(str1) && end < len(str2); end++ {
		if str1[end] != str2[end] {
			break
		}
	}
	public_prefix := str1[:end]
label:
	for len(public_prefix) > 0 {
		if len(str1)%len(public_prefix) != 0 || len(str2)%len(public_prefix) != 0 {
			public_prefix = public_prefix[:len(public_prefix)-1]
			continue
		}
		m := make(map[string]struct{})
		for i := 0; i < len(str1); i += len(public_prefix) {
			m[str1[i:i+len(public_prefix)]] = struct{}{}
			if len(m) > 1 {
				public_prefix = public_prefix[:len(public_prefix)-1]
				continue label
			}
		}
		for i := 0; i < len(str2); i += len(public_prefix) {
			m[str2[i:i+len(public_prefix)]] = struct{}{}
			if len(m) > 1 {
				public_prefix = public_prefix[:len(public_prefix)-1]
				continue label
			}
		}
		return public_prefix
	}
	return ""
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		str1 string
		str2 string
		want string
	}{
		{"ABCABC", "ABC", "ABC"},
		{"ABABAB", "ABAB", "AB"},
		{"LEET", "CODE", ""},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.want, gcdOfStrings(tt.str1, tt.str2))
	}
}
