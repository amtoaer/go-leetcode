/*
 * @lc app=leetcode.cn id=161 lang=golang
 * @lcpr version=30202
 *
 * [161] 相隔为 1 的编辑距离
 */

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// @lc code=start
func isOneEditDistance(s string, t string) bool {
	if abs(len(s)-len(t)) > 1 {
		return false
	}
	var left, right int
	for left < len(s) && right < len(t) && s[left] == t[right] {
		left++
		right++
	}
	// 字符串相同
	if left == len(s) && right == len(t) {
		return false
	}
	//有一边到达终点，另一边多一个字符，编辑距离肯定为 1
	if left == len(s) || right == len(t) {
		return true
	}
	return s[left+1:] == t[right:] || s[left:] == t[right+1:] || s[left+1:] == t[right+1:]
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		s    string
		t    string
		want bool
	}{
		{"ab", "acb", true},
		{"cab", "ad", false},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.want, isOneEditDistance(tt.s, tt.t))
	}
}

/*
// @lcpr case=start
// "ab"\n"acb"\n
// @lcpr case=end

// @lcpr case=start
// "cab"\n"ad"\n
// @lcpr case=end

*/
