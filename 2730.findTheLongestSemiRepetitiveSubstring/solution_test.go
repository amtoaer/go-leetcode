/*
 * @lc app=leetcode.cn id=2730 lang=golang
 * @lcpr version=30201
 *
 * [2730] 找到最长的半重复子字符串
 */

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// @lc code=start
func longestSemiRepetitiveSubstring(s string) int {
	res := 1
	var left, cnt int
	for right := 1; right < len(s); right++ {
		if s[right] == s[right-1] {
			cnt++
		}
		if cnt > 1 {
			left++
			for s[left] != s[left+1] {
				left++
			}
			cnt--
		}
		res = max(res, right-left+1)
	}
	return res
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		s    string
		want int
	}{
		{"52233", 4},
		{"5494", 4},
		{"1111111", 2},
	}
	for _, tt := range tc {
		t.Run(tt.s, func(t *testing.T) {
			assert.Equal(t, tt.want, longestSemiRepetitiveSubstring(tt.s))
		})
	}
}

/*
// @lcpr case=start
// "52233"\n
// @lcpr case=end

// @lcpr case=start
// "5494"\n
// @lcpr case=end

// @lcpr case=start
// "1111111"\n
// @lcpr case=end

*/
