package main

import (
	"testing"
)

/*
 * @lc app=leetcode.cn id=3 lang=golang
 * @lcpr version=20003
 *
 * [3] 无重复字符的最长子串
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func lengthOfLongestSubstring(s string) int {
	var (
		res   int
		empty = struct{}{}
		m     = make(map[byte]struct{})
	)
	for left, right := 0, 0; right < len(s); right++ {
		for {
			if _, ok := m[s[right]]; ok {
				delete(m, s[left])
				left++
			} else {
				break
			}
		}
		m[s[right]] = empty
		res = max(res, right-left+1)
	}
	return res
}

// @lc code=end

/*
// @lcpr case=start
// "abcabcbb"\n
// @lcpr case=end

// @lcpr case=start
// "bbbbb"\n
// @lcpr case=end

// @lcpr case=start
// "pwwkew"\n
// @lcpr case=end

*/

func Test(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
		{"", 0},
		{"abcdef", 6},
		{"aab", 2},
		{"dvdf", 3},
	}

	for _, test := range tests {
		if result := lengthOfLongestSubstring(test.input); result != test.expected {
			t.Errorf("For input '%s', expected %d but got %d", test.input, test.expected, result)
		}
	}
}
