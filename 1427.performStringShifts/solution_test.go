/*
 * @lc app=leetcode.cn id=1427 lang=golang
 * @lcpr version=30202
 *
 * [1427] 字符串的左右移
 */

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// @lc code=start
func stringShift(s string, shift [][]int) string {
	var movement int
	for _, s := range shift {
		if s[0] == 1 {
			movement -= s[1]
		} else {
			movement += s[1]
		}
	}
	movement %= len(s)
	if movement < 0 {
		movement += len(s)
	}
	return s[movement:] + s[:movement]
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		s     string
		shift [][]int
		want  string
	}{
		{"abc", [][]int{{0, 1}, {1, 2}}, "cab"},
		{"abcdefg", [][]int{{1, 1}, {1, 1}, {0, 2}, {1, 3}}, "efgabcd"},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.want, stringShift(tt.s, tt.shift))
	}
}

/*
// @lcpr case=start
// "abc"\n[[0,1],[1,2]]\n
// @lcpr case=end

// @lcpr case=start
// "abcdefg"\n[[1,1],[1,1],[0,2],[1,3]]\n
// @lcpr case=end

*/
