/*
 * @lc app=leetcode.cn id=2697 lang=golang
 * @lcpr version=30111
 *
 * [2697] 字典序最小回文串
 */

// @lcpr-template-start
package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	assert.Equal(t, "abba", makeSmallestPalindrome("abcd"))
	assert.Equal(t, "efcfe", makeSmallestPalindrome("egcfe"))
	assert.Equal(t, "aiia", makeSmallestPalindrome("atie"))
}

// @lcpr-template-end
// @lc code=start
func min(a, b byte) byte {
	if a < b {
		return a
	}
	return b
}

func makeSmallestPalindrome(s string) string {
	left, right := 0, len(s)-1
	t := []byte(s)
	for left < right {
		if s[left] != s[right] {
			t[left] = min(s[left], s[right])
			t[right] = t[left]
		}
		left++
		right--
	}
	return string(t)
}

// @lc code=end

/*
// @lcpr case=start
// "egcfe"\n
// @lcpr case=end

// @lcpr case=start
// "abcd"\n
// @lcpr case=end

// @lcpr case=start
// "seven"\n
// @lcpr case=end

*/
