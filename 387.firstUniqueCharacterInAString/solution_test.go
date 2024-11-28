package main

import (
	"testing"
)

/*
 * @lc app=leetcode.cn id=387 lang=golang
 * @lcpr version=20003
 *
 * [387] 字符串中的第一个唯一字符
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start

func firstUniqChar(s string) int {
	cnt := [26]int{}
	for _, ch := range s {
		cnt[ch-'a']++
	}
	for i, ch := range s {
		if cnt[ch-'a'] == 1 {
			return i
		}
	}
	return -1
}

// @lc code=end

/*
// @lcpr case=start
// "leetcode"\n
// @lcpr case=end

// @lcpr case=start
// "loveleetcode"\n
// @lcpr case=end

// @lcpr case=start
// "aabb"\n
// @lcpr case=end

*/

func Test(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"leetcode", 0},
		{"loveleetcode", 2},
		{"aabb", -1},
		{"", -1},
		{"a", 0},
		{"ab", 0},
		{"aabbccddeeffg", 12},
		{"aadadaad", -1},
	}

	for _, test := range tests {
		result := firstUniqChar(test.input)
		if result != test.expected {
			t.Errorf("For input '%s', expected %d but got %d", test.input, test.expected, result)
		}
	}
}
