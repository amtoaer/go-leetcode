package main

import (
	"testing"
)

/*
 * @lc app=leetcode.cn id=316 lang=golang
 * @lcpr version=20003
 *
 * [316] 去除重复字母
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func removeDuplicateLetters(s string) string {
	m := make(map[rune]int)
	exist := make(map[rune]struct{})
	for _, r := range s {
		m[r]++
	}
	var res []rune
	for _, r := range s {
		if _, ok := exist[r]; !ok {
			for len(res) > 0 && res[len(res)-1] > r && m[res[len(res)-1]] > 0 {
				delete(exist, res[len(res)-1])
				res = res[:len(res)-1]
			}
			res = append(res, r)
			exist[r] = struct{}{}
		}
		m[r]--
	}
	return string(res)
}

// @lc code=end

/*
// @lcpr case=start
// "bcabc"\n
// @lcpr case=end

// @lcpr case=start
// "cbacdcbc"\n
// @lcpr case=end

*/

func Test(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"bcabc", "abc"},
		{"cbacdcbc", "acdb"},
		{"abacb", "abc"},
		{"bbcaac", "bac"},
		{"", ""},
	}

	for _, test := range tests {
		result := removeDuplicateLetters(test.input)
		if result != test.expected {
			t.Errorf("removeDuplicateLetters(%q) = %q; want %q", test.input, result, test.expected)
		}
	}
}
