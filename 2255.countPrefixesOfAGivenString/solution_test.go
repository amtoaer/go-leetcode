/*
 * @lc app=leetcode.cn id=2255 lang=golang
 * @lcpr version=30104
 *
 * [2255] 统计是给定字符串前缀的字符串数目
 */

package main

import (
	"strings"
	"testing"
)

// @lc code=start
func countPrefixes(words []string, s string) int {
	var cnt int
	for _, word := range words {
		if strings.HasPrefix(s, word) {
			cnt++
		}
	}
	return cnt
}

// @lc code=end

func Test(t *testing.T) {
	testCases := []struct {
		words    []string
		s        string
		expected int
	}{
		{[]string{"a", "b", "c", "ab", "bc", "abc"}, "abc", 3},
		{[]string{"a", "a"}, "aa", 2},
		{[]string{}, "abc", 0},
		{[]string{"abc", "abcd", "ab"}, "abcde", 3},
		{[]string{"xyz", "abc"}, "def", 0},
		{[]string{"x", "y", "z"}, "xyz", 1},
	}

	for _, tc := range testCases {
		result := countPrefixes(tc.words, tc.s)
		if result != tc.expected {
			t.Errorf("countPrefixes(%v, %s) = %d, expected %d", tc.words, tc.s, result, tc.expected)
		}
	}
}

/*
// @lcpr case=start
// ["a","b","c","ab","bc","abc"]\n"abc"\n
// @lcpr case=end

// @lcpr case=start
// ["a","a"]\n"aa"\n
// @lcpr case=end

*/
