package main

import (
	"testing"
)

/*
 * @lc app=leetcode.cn id=409 lang=golang
 * @lcpr version=20004
 *
 * [409] 最长回文串
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func longestPalindrome(s string) int {
	m := make(map[rune]int)
	for _, r := range s {
		m[r]++
	}
	var (
		res    int
		hasOdd bool
	)
	for _, c := range m {
		res += c - c%2
		hasOdd = hasOdd || c&1 == 1
	}
	if hasOdd {
		res += 1
	}
	return res
}

// @lc code=end

/*
// @lcpr case=start
// "abccccdd"\n
// @lcpr case=end

// @lcpr case=start
// "a"\n
// @lcpr case=end

*/

func Test(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"abccccdd", 7},
		{"a", 1},
		{"bb", 2},
		{"ccc", 3},
		{"bananas", 5},
	}

	for _, test := range tests {
		if result := longestPalindrome(test.input); result != test.expected {
			t.Errorf("For input '%s', expected %d but got %d", test.input, test.expected, result)
		}
	}
}
