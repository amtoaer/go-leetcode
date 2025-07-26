/*
 * @lc app=leetcode.cn id=266 lang=golang
 * @lcpr version=30202
 *
 * [266] 回文排列
 */

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// @lc code=start
func canPermutePalindrome(s string) bool {
	m := make(map[byte]int)
	for i := range s {
		m[s[i]]++
	}
	var oddCnt int
	for _, cnt := range m {
		oddCnt += cnt % 2
	}
	return oddCnt <= 1
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		s    string
		want bool
	}{
		{"code", false},
		{"aab", true},
		{"carerac", true},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.want, canPermutePalindrome(tt.s))
	}
}

/*
// @lcpr case=start
// "code"\n
// @lcpr case=end

// @lcpr case=start
// "aab"\n
// @lcpr case=end

// @lcpr case=start
// "carerac"\n
// @lcpr case=end

*/
